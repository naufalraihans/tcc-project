# ⚙️ Environment Variables — TCC ITPLN Web Platform

> **Dokumen ini mendefinisikan semua environment variables yang dibutuhkan.**
> Versi: v1.0 | Status: Draft

> [!CAUTION]
> File `.env` TIDAK boleh di-commit ke Git. Selalu gunakan `.env.example` sebagai template.
> Pastikan `.env` sudah ada di `.gitignore`.

---

## 1. Frontend — SvelteKit (`.env`)

File lokasi: `frontend/.env`

```env
# ── SUPABASE ──────────────────────────────────────────────────
# Ambil dari: Supabase Dashboard → Project Settings → API
VITE_SUPABASE_URL=https://xxxxxxxxxxxx.supabase.co
VITE_SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# ── GO BACKEND ────────────────────────────────────────────────
# URL base API Go yang digunakan oleh frontend
VITE_API_BASE_URL=http://localhost:8080/api/v1

# ── MIDTRANS ──────────────────────────────────────────────────
# Client Key digunakan di frontend untuk inisialisasi Snap.js
VITE_MIDTRANS_CLIENT_KEY=SB-Mid-client-xxxxxxxxxxxx
# Env: 'sandbox' untuk development, 'production' untuk live
VITE_MIDTRANS_ENV=sandbox

# ── APP ───────────────────────────────────────────────────────
VITE_APP_NAME=TCC ITPLN
VITE_APP_URL=http://localhost:5173
```

### Penjelasan Variabel Frontend

| Variabel | Wajib | Keterangan |
|----------|-------|------------|
| `VITE_SUPABASE_URL` | ✅ | URL project Supabase — aman untuk expose ke client |
| `VITE_SUPABASE_ANON_KEY` | ✅ | Anon/public key Supabase — aman untuk expose ke client |
| `VITE_API_BASE_URL` | ✅ | Base URL Go backend |
| `VITE_MIDTRANS_CLIENT_KEY` | ✅ | Client key Midtrans untuk Snap.js di browser |
| `VITE_MIDTRANS_ENV` | ✅ | Mode Midtrans: `sandbox` atau `production` |
| `VITE_APP_NAME` | ❌ | Nama aplikasi untuk meta tags |
| `VITE_APP_URL` | ❌ | URL frontend untuk canonical URL / SEO |

> **Penting:** Semua variabel SvelteKit harus diawali `VITE_` agar bisa diakses di browser. Variabel tanpa prefix `VITE_` hanya bisa diakses di server-side (SSR).

---

## 2. Backend — Go (`.env`)

File lokasi: `backend/.env`

```env
# ── SERVER ────────────────────────────────────────────────────
APP_ENV=development
APP_PORT=8080

# ── DATABASE ──────────────────────────────────────────────────
# Connection string ke Supabase PostgreSQL
# Ambil dari: Supabase Dashboard → Project Settings → Database → Connection String (URI)
DATABASE_URL=postgresql://postgres:[PASSWORD]@db.xxxxxxxxxxxx.supabase.co:5432/postgres

# ── SUPABASE ──────────────────────────────────────────────────
# Ambil dari: Supabase Dashboard → Project Settings → API
SUPABASE_URL=https://xxxxxxxxxxxx.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
SUPABASE_SERVICE_ROLE_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# JWT Secret untuk verifikasi token dari Supabase Auth
# Ambil dari: Supabase Dashboard → Project Settings → API → JWT Settings
SUPABASE_JWT_SECRET=your-super-secret-jwt-secret-here

# ── MIDTRANS ──────────────────────────────────────────────────
# Server Key digunakan di backend untuk membuat transaksi & verifikasi webhook
MIDTRANS_SERVER_KEY=SB-Mid-server-xxxxxxxxxxxx
MIDTRANS_CLIENT_KEY=SB-Mid-client-xxxxxxxxxxxx
# Env: true = sandbox, false = production
MIDTRANS_IS_PRODUCTION=false

# ── CORS ──────────────────────────────────────────────────────
# Origins yang diizinkan (pisahkan dengan koma jika lebih dari 1)
CORS_ALLOWED_ORIGINS=http://localhost:5173,https://tcc-itpln.id

# ── STORAGE ───────────────────────────────────────────────────
SUPABASE_STORAGE_URL=https://xxxxxxxxxxxx.supabase.co/storage/v1
```

### Penjelasan Variabel Backend

| Variabel | Wajib | Keterangan |
|----------|-------|------------|
| `APP_ENV` | ✅ | `development` atau `production` |
| `APP_PORT` | ✅ | Port server Go (default: `8080`) |
| `DATABASE_URL` | ✅ | PostgreSQL connection string Supabase |
| `SUPABASE_URL` | ✅ | URL project Supabase |
| `SUPABASE_ANON_KEY` | ✅ | Public key (untuk operasi biasa) |
| `SUPABASE_SERVICE_ROLE_KEY` | ✅ | Admin key — **RAHASIA**, jangan expose ke client |
| `SUPABASE_JWT_SECRET` | ✅ | Secret untuk verifikasi JWT user — **RAHASIA** |
| `MIDTRANS_SERVER_KEY` | ✅ | Server key Midtrans — **RAHASIA**, jangan expose ke client |
| `MIDTRANS_CLIENT_KEY` | ✅ | Client key Midtrans |
| `MIDTRANS_IS_PRODUCTION` | ✅ | `false` = sandbox, `true` = production |
| `CORS_ALLOWED_ORIGINS` | ✅ | Whitelist origin untuk CORS |
| `SUPABASE_STORAGE_URL` | ❌ | Base URL Supabase Storage |

---

## 3. File `.env.example`

Template ini yang di-commit ke Git. Semua value dikosongkan atau diganti placeholder.

### `frontend/.env.example`
```env
VITE_SUPABASE_URL=https://your-project.supabase.co
VITE_SUPABASE_ANON_KEY=your-supabase-anon-key
VITE_API_BASE_URL=http://localhost:8080/api/v1
VITE_MIDTRANS_CLIENT_KEY=SB-Mid-client-your-client-key
VITE_MIDTRANS_ENV=sandbox
VITE_APP_NAME=TCC ITPLN
VITE_APP_URL=http://localhost:5173
```

### `backend/.env.example`
```env
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgresql://postgres:password@db.your-project.supabase.co:5432/postgres
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=your-supabase-anon-key
SUPABASE_SERVICE_ROLE_KEY=your-service-role-key
SUPABASE_JWT_SECRET=your-jwt-secret
MIDTRANS_SERVER_KEY=SB-Mid-server-your-server-key
MIDTRANS_CLIENT_KEY=SB-Mid-client-your-client-key
MIDTRANS_IS_PRODUCTION=false
CORS_ALLOWED_ORIGINS=http://localhost:5173
SUPABASE_STORAGE_URL=https://your-project.supabase.co/storage/v1
```

---

## 4. `.gitignore` yang Wajib Ada

```gitignore
# Environment
.env
.env.local
.env.*.local

# Frontend
frontend/node_modules/
frontend/.svelte-kit/
frontend/build/

# Backend
backend/bin/
backend/tmp/

# OS
.DS_Store
Thumbs.db
```

---

## 5. Cara Setup Environment (Onboarding Developer Baru)

```
1. Clone repo
2. Copy frontend/.env.example → frontend/.env
3. Copy backend/.env.example → backend/.env
4. Minta nilai rahasia (JWT Secret, Service Role Key, Midtrans Key) dari lead/project owner
5. Isi semua value di kedua .env
6. Jalankan: cd frontend && npm install && npm run dev
7. Jalankan: cd backend && go run ./cmd/server/main.go
```

---

*Dokumen ini akan diperbarui seiring penambahan dependency atau konfigurasi baru.*
