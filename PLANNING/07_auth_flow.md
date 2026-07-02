# 🔑 Auth Flow — TCC ITPLN Web Platform

> **Dokumen ini menjelaskan alur autentikasi secara detail: dari register hingga role-based access.**
> Versi: v1.0 | Status: Draft

---

## 1. Gambaran Besar

```
[User/Browser]
     │
     │  1. Login/Register via Supabase Auth (di frontend)
     ▼
[SvelteKit Frontend]  ←──── Supabase JS Client (manage session otomatis)
     │
     │  2. Kirim request + JWT di header ke Go backend
     ▼
[Go Backend / API]
     │
     │  3. Verifikasi JWT pakai SUPABASE_JWT_SECRET
     │  4. Cek role dari DB (tabel profiles)
     ▼
[Supabase PostgreSQL]
```

**Prinsip utama:**
- Supabase Auth yang **pegang session** — bukan kita yang manage cookie/token manual
- Go backend **hanya verifikasi** JWT yang datang — tidak punya session sendiri
- Role (`user` / `admin`) disimpan di tabel `profiles`, **bukan** di JWT claims Supabase (karena claims tidak bisa di-custom di Supabase free tier)

---

## 2. Alur Register

```
User isi form register (nama, email, password)
     │
     ▼
SvelteKit → supabase.auth.signUp({ email, password, options: { data: { full_name } } })
     │
     ▼
Supabase Auth:
  - Buat user baru di auth.users
  - Kirim email konfirmasi ke user
  - Trigger otomatis → INSERT ke profiles (full_name, role='user')
     │
     ▼
Frontend terima response:
  - Jika email confirmation ON → tampilkan pesan "Cek email kamu"
  - Jika email confirmation OFF → langsung dapat session → redirect ke /dashboard
```

### Hal yang Perlu Diperhatikan
- `full_name` dikirim via `raw_user_meta_data` → ditangkap trigger di `SUPABASE.md` (010_create_triggers.sql)
- Role default `'user'` di-set oleh trigger, **bukan** oleh frontend — tidak bisa dimanipulasi user
- Jika email belum dikonfirmasi, Supabase tidak akan return session aktif

---

## 3. Alur Login

```
User isi email + password → klik Login
     │
     ▼
SvelteKit → supabase.auth.signInWithPassword({ email, password })
     │
     ├── Gagal (wrong credentials) → tampilkan error "Email atau password salah"
     │
     └── Berhasil → Supabase return:
           - access_token  (JWT, expire: 1 jam)
           - refresh_token (expire: lebih lama, dirotasi setiap pakai)
           - user object
           │
           ▼
     Supabase JS Client simpan session di localStorage/cookie otomatis
           │
           ▼
     Frontend ambil role user:
       → GET /api/v1/auth/me ke Go backend (kirim access_token)
       → Backend return { id, full_name, email, role, ... }
       → Simpan ke auth store (Svelte store)
           │
           ▼
     Redirect berdasarkan role:
       - role 'admin' → /admin/dashboard
       - role 'user'  → /dashboard
```

---

## 4. Session Management di SvelteKit

### 4.1 Inisialisasi Supabase Client

`src/lib/supabase.ts`
```typescript
import { createBrowserClient } from '@supabase/ssr'
import { PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_ANON_KEY } from '$env/static/public'

export const supabase = createBrowserClient(
  PUBLIC_SUPABASE_URL,
  PUBLIC_SUPABASE_ANON_KEY
)
```

### 4.2 Root Layout — Cek Session Awal

`src/routes/+layout.ts`
```typescript
// Load function ini jalan di setiap page load (SSR & client)
export const load = async ({ fetch, data }) => {
  const { data: { session } } = await supabase.auth.getSession()

  return {
    session,
    user: session?.user ?? null
  }
}
```

### 4.3 Auth Store (Global State)

`src/lib/stores/auth.store.ts`
```typescript
// Menyimpan: session, user object, role
// Diupdate saat login, logout, atau token refresh
// Diakses oleh komponen manapun untuk cek role
```

### 4.4 Auto Token Refresh

Supabase JS Client **otomatis** refresh `access_token` menggunakan `refresh_token` sebelum expired.

```typescript
// Di root +layout.svelte — listen perubahan auth state
supabase.auth.onAuthStateChange((event, session) => {
  if (event === 'TOKEN_REFRESHED') {
    // Update auth store dengan session baru
  }
  if (event === 'SIGNED_OUT') {
    // Clear auth store, redirect ke /auth/login
  }
})
```

---

## 5. Route Guard (Proteksi Halaman)

### 5.1 Guard Halaman User (`/dashboard/*`)

`src/routes/dashboard/+layout.ts`
```typescript
import { redirect } from '@sveltejs/kit'

export const load = async ({ parent }) => {
  const { session } = await parent()

  // Jika tidak ada session → redirect ke login
  if (!session) {
    throw redirect(302, '/auth/login')
  }

  return { session }
}
```

### 5.2 Guard Halaman Admin (`/admin/*`)

`src/routes/admin/+layout.ts`
```typescript
import { redirect } from '@sveltejs/kit'

export const load = async ({ parent, fetch }) => {
  const { session } = await parent()

  // Jika tidak ada session → redirect ke login
  if (!session) {
    throw redirect(302, '/auth/login')
  }

  // Ambil role dari Go backend
  const res = await fetch('/api/v1/auth/me', {
    headers: { Authorization: `Bearer ${session.access_token}` }
  })
  const profile = await res.json()

  // Jika bukan admin → redirect ke dashboard dengan pesan
  if (profile.data?.role !== 'admin') {
    throw redirect(302, '/dashboard?error=akses_ditolak')
  }

  return { session, profile: profile.data }
}
```

### 5.3 Guard Halaman Auth (sudah login tidak boleh buka /auth/login)

`src/routes/auth/+layout.ts`
```typescript
import { redirect } from '@sveltejs/kit'

export const load = async ({ parent }) => {
  const { session, user } = await parent()

  if (session) {
    // Sudah login → redirect sesuai role
    const role = user?.role ?? 'user'
    throw redirect(302, role === 'admin' ? '/admin/dashboard' : '/dashboard')
  }
}
```

---

## 6. Alur Request ke Go Backend

Setiap request ke endpoint protected Go backend:

```
Frontend ambil access_token dari Supabase session
     │
     ▼
Kirim HTTP request:
  GET /api/v1/dashboard/kelas
  Headers:
    Authorization: Bearer eyJhbGci...   ← access_token
     │
     ▼
Go Backend — Middleware Auth:
  1. Ekstrak token dari header Authorization
  2. Verifikasi signature JWT pakai SUPABASE_JWT_SECRET
  3. Cek expiry (exp claim)
  4. Ekstrak user ID dari JWT (sub claim)
  5. Query DB: SELECT role FROM profiles WHERE id = $1
  6. Set user context (id + role) untuk handler berikutnya
     │
     ├── Token invalid / expired → 401 UNAUTHORIZED
     ├── Role tidak cukup       → 403 FORBIDDEN
     └── OK → lanjut ke handler
```

### JWT Claims dari Supabase (isi token)
```json
{
  "sub": "uuid-user-id",
  "email": "user@example.com",
  "role": "authenticated",
  "aud": "authenticated",
  "exp": 1234567890,
  "iat": 1234567890
}
```

> **Penting:** `role` di JWT Supabase selalu `"authenticated"` — bukan `"admin"` atau `"user"`. Role custom kita disimpan di tabel `profiles`, bukan di JWT. Jadi Go backend **wajib query DB** untuk cek role.

---

## 7. Alur Logout

```
User klik tombol Logout
     │
     ▼
supabase.auth.signOut()
  - Supabase hapus session dari localStorage
  - Invalidate refresh token di server Supabase
     │
     ▼
Auth store di-clear (user = null, role = null)
     │
     ▼
Redirect ke / (beranda)
```

---

## 7.1 Alur Reset Password (Lupa Password)

Ditangani sepenuhnya oleh Supabase Auth — Go backend tidak terlibat.

```
Halaman /auth/lupa-password
  User isi email → klik "Kirim link reset"
     │
     ▼
supabase.auth.resetPasswordForEmail(email, {
  redirectTo: `${APP_URL}/auth/reset-password`
})
  → Supabase kirim email berisi magic link ke user
  → Selalu tampilkan pesan sukses generik ("Jika email terdaftar, link
    sudah dikirim") — JANGAN bocorkan apakah email ada di sistem
     │
     ▼
User klik link di email → mendarat di /auth/reset-password
  → Supabase JS otomatis menukar token di URL jadi session sementara
    (event onAuthStateChange: 'PASSWORD_RECOVERY')
     │
     ▼
Halaman /auth/reset-password (BARU — belum ada di prototipe)
  User isi password baru + konfirmasi
     │
     ▼
supabase.auth.updateUser({ password: newPassword })
  → Sukses → tampilkan pesan → redirect ke /auth/login
  → Token kedaluwarsa/invalid → "Link reset sudah tidak berlaku, minta lagi"
```

**Yang perlu disiapkan:**
- Halaman baru `/auth/reset-password` (tidak ada prototipe — pakai style Login).
- URL `${APP_URL}/auth/reset-password` masuk ke **Redirect URLs** di Supabase Auth settings (lihat `06` §6).
- Guard: halaman ini boleh diakses tanpa login penuh, tapi wajib ada session `PASSWORD_RECOVERY` — kalau dibuka langsung tanpa token, redirect ke `/auth/lupa-password`.

---

## 8. Edge Cases & Penanganannya

### 8.1 Token Expired di Tengah Sesi
```
User lagi buka halaman dashboard
Access token expired (1 jam habis)
     │
     ▼
Supabase JS Client deteksi → auto refresh pakai refresh_token
Dapat access_token baru → lanjut request seperti biasa
     │
Jika refresh_token juga expired (user lama tidak aktif):
     ▼
onAuthStateChange event: 'SIGNED_OUT'
→ Clear store → redirect ke /auth/login
→ Tampilkan pesan "Sesi kamu telah berakhir, silakan login kembali"
```

### 8.2 User Buka `/auth/login` saat Sudah Login
```
Guard di /auth/+layout.ts deteksi session aktif
→ Redirect otomatis ke /dashboard (user) atau /admin/dashboard (admin)
```

### 8.3 User Biasa Coba Akses `/admin/*`
```
Guard di /admin/+layout.ts cek role dari Go backend
Role = 'user' → redirect ke /dashboard?error=akses_ditolak
Halaman dashboard tampilkan toast/banner "Kamu tidak memiliki akses ke halaman tersebut"
```

### 8.4 Go Backend Terima Token Expired
```
Middleware cek exp claim → token expired
Return 401 { error: "UNAUTHORIZED", message: "Token telah kedaluwarsa" }
     │
Frontend harus handle 401:
→ Coba refresh session via Supabase
→ Jika berhasil → retry request dengan token baru
→ Jika gagal → redirect ke login
```

---

## 9. Ringkasan Tanggung Jawab

| Tanggung Jawab | Siapa yang Handle |
|----------------|-------------------|
| Simpan & refresh JWT | Supabase JS Client (otomatis) |
| Register & Login UI | SvelteKit Frontend |
| Proteksi route (redirect) | SvelteKit `+layout.ts` |
| Verifikasi JWT | Go Backend Middleware |
| Cek role `admin`/`user` | Go Backend (query tabel `profiles`) |
| State role di UI | Svelte Auth Store |
| Auto-create `profiles` saat register | Supabase Trigger (DB level) |

---

*Dokumen ini akan diperbarui seiring implementasi autentikasi.*
