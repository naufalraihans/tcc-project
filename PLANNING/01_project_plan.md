# 📋 Project Plan — TCC ITPLN Web Platform

> **Dokumen ini adalah perencanaan (v2.0) untuk pembangunan platform web TCC ITPLN.**
> Dibuat: 2026-07-02 | Status: **Draft**

---

## 1. Latar Belakang

**TCC (Training and Consultant Center)** adalah lembaga yang berdiri di bawah naungan **ITPLN (Institut Teknologi PLN)**. TCC berperan sebagai pusat pelatihan dan konsultasi yang menyediakan program pembelajaran berbasis kelistrikan dan hal-hal terapan yang berkaitan dengan PLN (bukan ITPLN).

Platform web ini akan menjadi **wajah digital TCC** — media informasi sekaligus sarana pembelajaran bagi masyarakat umum yang ingin mengembangkan kompetensi di bidang ketenagalistrikan.

---

## 2. Objektif Project

| # | Objektif |
|---|----------|
| 1 | Menyediakan media informasi publik mengenai TCC dan program-programnya |
| 2 | Menyediakan platform pendaftaran dan akses kelas pelatihan (online & offline) |
| 3 | Menyediakan layanan konsultasi antara pengguna dan tenaga ahli TCC |
| 4 | Memperluas jangkauan TCC ke masyarakat luas secara digital |

---

## 3. Role & Hak Akses

### 3.1 Definisi Role

Platform ini memiliki **3 tingkat akses** yang berlaku secara hierarkis:

```
Publik (Guest)  ⊂  User (Login)  ⊂  Admin
```

| Role | Siapa | Cara Masuk |
|------|-------|------------|
| **Publik / Guest** | Siapa saja yang membuka website tanpa login | Langsung akses, tidak perlu akun |
| **User (Peserta)** | Masyarakat yang sudah mendaftar dan login | Register → Login via Supabase Auth |
| **Admin** | Staff TCC yang ditunjuk | Akun dibuat manual oleh Super Admin; role di-set dari DB |

---

### 3.2 Matriks Hak Akses Per Halaman

> **Keterangan Simbol:**
> - ✅ = Boleh akses / bisa melakukan aksi
> - ❌ = Tidak bisa akses (redirect / tampil pesan)
> - 👁️ = Bisa lihat saja (read-only)
> - ✏️ = Bisa aksi (buat, ubah, hapus)

#### Halaman Publik

| Halaman | URL | Publik | User | Admin |
|---------|-----|--------|------|-------|
| Landing Page | `/` | ✅ | ✅ | ✅ |
| Daftar Kelas | `/kelas` | ✅ (lihat saja) | ✅ (+ bisa daftar) | ✅ |
| Detail Kelas | `/kelas/[slug]` | ✅ (lihat saja) | ✅ (+ bisa daftar) | ✅ |
| Topik | `/topik` | ✅ | ✅ | ✅ |
| Konsultasi | `/konsultasi` | 👁️ (form bisa diisi, tapi redirect login) | ✅ (submit form) | ✅ |
| Tentang | `/tentang` | ✅ | ✅ | ✅ |

#### Halaman Autentikasi

| Halaman | URL | Publik | User | Admin |
|---------|-----|--------|------|-------|
| Login | `/auth/login` | ✅ | ❌ (redirect dashboard) | ❌ (redirect admin) |
| Register | `/auth/register` | ✅ | ❌ (redirect dashboard) | ❌ (redirect admin) |
| Lupa Password | `/auth/lupa-password` | ✅ | ✅ | ✅ |

#### Halaman User (Requires Login)

| Halaman | URL | Publik | User | Admin |
|---------|-----|--------|------|-------|
| Dashboard User | `/dashboard` | ❌ (redirect login) | ✅ | ✅ |
| Kelas Saya | `/dashboard/kelas` | ❌ | ✅ | ✅ |
| Detail Kelas (Akses Materi) | `/dashboard/kelas/[id]` | ❌ | ✅ (jika terdaftar) | ✅ |
| Riwayat Konsultasi | `/dashboard/konsultasi` | ❌ | ✅ | ✅ |
| Profil | `/dashboard/profil` | ❌ | ✅ | ✅ |

#### Halaman Admin (Requires Admin Role)

| Halaman | URL | Publik | User | Admin |
|---------|-----|--------|------|-------|
| Dashboard Admin | `/admin/dashboard` | ❌ (redirect login) | ❌ (redirect dashboard) | ✅ |
| Kelola Kelas | `/admin/kelas` | ❌ | ❌ | ✅ (CRUD penuh) |
| Kelola Peserta | `/admin/peserta` | ❌ | ❌ | ✅ (lihat, kelola, verifikasi) |
| Kelola Konsultasi | `/admin/konsultasi` | ❌ | ❌ | ✅ (lihat & respon) |
| Kelola Topik | `/admin/topik` | ❌ | ❌ | ✅ (CRUD penuh) |
| Kelola Transaksi | `/admin/transaksi` | ❌ | ❌ | ✅ (lihat & verifikasi) |

---

### 3.3 Aturan Akses (Access Control Rules)

#### Rule 1 — Redirect Tidak Sah
| Kondisi | Aksi Sistem |
|---------|-------------|
| Publik mencoba akses `/dashboard/*` | Redirect ke `/auth/login` |
| Publik mencoba akses `/admin/*` | Redirect ke `/auth/login` |
| User (bukan Admin) mencoba akses `/admin/*` | Redirect ke `/dashboard` dengan pesan "Akses Ditolak" |
| User yang sudah login mencoba akses `/auth/login` atau `/auth/register` | Redirect ke `/dashboard` |
| Admin yang sudah login mencoba akses `/auth/login` | Redirect ke `/admin/dashboard` |

#### Rule 2 — Validasi Role di Backend
- Setiap request ke endpoint `/api/admin/*` wajib diverifikasi JWT **dan** dicek field `role = "admin"` dari database / JWT claims
- Setiap request ke endpoint `/api/user/*` wajib diverifikasi JWT minimal valid (role apapun yang sudah login)
- Endpoint publik `/api/public/*` tidak memerlukan JWT

#### Rule 3 — Role Assignment
- User baru yang register otomatis mendapat role **`user`**
- Role **`admin`** hanya bisa di-set secara manual oleh admin yang sudah ada (melalui panel admin atau langsung di database)
- Tidak ada self-upgrade role (user tidak bisa mengklaim dirinya admin)

---

## 4. Struktur Kelas / Pelatihan

### 4.1 Berdasarkan Harga
| Tipe | Keterangan |
|------|------------|
| **Gratis** | Dapat diakses setelah user terdaftar, tanpa biaya |
| **Berbayar** | User perlu melakukan pembayaran sebelum mengakses konten kelas |

### 4.2 Berdasarkan Format Penyelenggaraan
| Format | Keterangan |
|--------|------------|
| **Online** | Kelas dilaksanakan secara daring (via platform meeting / LMS) |
| **Offline (Workshop)** | Kelas dilaksanakan secara tatap muka di lokasi TCC / ITPLN |
| **Hybrid** | Kombinasi online dan offline — peserta bisa memilih atau bergantian |

### 4.3 Atribut Kelas
Setiap kelas memiliki atribut berikut:
- Judul & deskripsi
- Topik / kategori
- Instruktur
- Format: Online / Offline
- Harga: Gratis / Berbayar (+ nominal)
- Jadwal & durasi
- Kuota peserta
- Status: Aktif / Penuh / Selesai
- Lokasi (untuk offline/hybrid)

### 4.4 Aturan Pendaftaran Kelas
- **Multi-kelas diizinkan** — 1 user dapat mendaftar banyak kelas sekaligus tanpa batas
- **Validasi konflik jadwal** — sistem wajib mencegah pendaftaran jika jadwal kelas baru bentrok dengan kelas yang sudah diikuti user
- **User tidak bisa membatalkan** — pendaftaran hanya bisa dibatalkan oleh admin
- **Materi tetap bisa diakses** — setelah kelas selesai, user masih bisa mengakses materi kelas

### 4.5 Alur Pembayaran (Kelas Berbayar)
1. User klik "Daftar" pada kelas berbayar
2. Sistem membuat **order/transaksi pending** di database
3. User diarahkan ke halaman pembayaran **Midtrans** (Snap UI)
4. Setelah pembayaran sukses → Midtrans kirim **webhook callback** ke Go backend
5. Backend verifikasi callback → update status transaksi → user otomatis masuk kelas
6. User mendapat konfirmasi (notifikasi — *TBD*)

---

## 5. Fitur & Halaman Web

### 5.1 Halaman Publik
> Dapat diakses oleh **siapa saja** tanpa perlu login.

#### Landing Page
- **URL:** `/`
- Hero section dengan tagline & CTA (daftar / lihat program)
- Preview kelas unggulan (Online & Offline)
- Topik pelatihan yang tersedia
- Tentang TCC (ringkas)
- Statistik platform (peserta, kelas, instruktur)
- CTA konsultasi

#### Kelas
- **URL:** `/kelas`
- Daftar semua kelas dengan filter: Topik, Format (Online/Offline/Hybrid), Harga (Gratis/Berbayar)
- Kartu kelas: judul, format badge, instruktur, jadwal, harga, kuota sisa
- **Detail Kelas:** `/kelas/[slug]`
  - Deskripsi lengkap, silabus, instruktur, jadwal
  - Tombol daftar kelas → redirect ke login jika belum login

#### Konsultasi
- **URL:** `/konsultasi`
- Penjelasan layanan konsultasi TCC
- Form pengajuan konsultasi → redirect ke login jika belum login

#### Topik
- **URL:** `/topik`
- Daftar kategori/topik pelatihan (misal: Kelistrikan Industri, K3, Smart Grid)
- Setiap topik mengarah ke daftar kelas terkait

#### Tentang
- **URL:** `/tentang`
- Sejarah, visi misi TCC, struktur organisasi, instruktur, kontak & alamat

---

### 5.2 Halaman Autentikasi

| Halaman | URL | Keterangan |
|---------|-----|------------|
| Login | `/auth/login` | Login dengan Supabase Auth |
| Register | `/auth/register` | Pendaftaran akun baru (role default: `user`) |
| Lupa Password | `/auth/lupa-password` | Minta link reset password via email |
| Reset Password | `/auth/reset-password` | Set password baru (dibuka dari link email) |

---

### 5.3 Halaman User (Requires Login)
> Hanya dapat diakses oleh pengguna yang **sudah login** (role: `user` atau `admin`).

| Halaman | URL | Keterangan |
|---------|-----|------------|
| Dashboard | `/dashboard` | Ringkasan kelas yang diikuti, jadwal mendatang, notifikasi |
| Kelas Saya | `/dashboard/kelas` | Daftar kelas yang sedang & sudah diikuti |
| Detail Kelas | `/dashboard/kelas/[id]` | Akses materi, link Zoom, atau info kelas offline |
| Riwayat Konsultasi | `/dashboard/konsultasi` | Lihat status & riwayat pengajuan konsultasi |
| Profil | `/dashboard/profil` | Kelola data diri (nama, foto, email, password) |

**Aksi yang bisa dilakukan User:**
- Mendaftar kelas (gratis langsung masuk; berbayar via Midtrans)
- Mengajukan konsultasi melalui form
- Melihat dan mengelola profil pribadi
- Melihat riwayat kelas & konsultasi

---

### 5.4 Halaman Admin (Requires Admin Role)
> Hanya dapat diakses oleh pengguna dengan role **`admin`**. Semua fitur bersifat CRUD penuh.

| Halaman | URL | Keterangan |
|---------|-----|------------|
| Dashboard Admin | `/admin/dashboard` | Statistik global: total peserta, kelas aktif, pendapatan, pengajuan konsultasi baru |
| Kelola Kelas | `/admin/kelas` | CRUD kelas (buat, edit, hapus, ubah status kelas) |
| Kelola Peserta | `/admin/peserta` | Lihat, kelola, dan verifikasi akun peserta; kelola enrollment manual |
| Kelola Konsultasi | `/admin/konsultasi` | Lihat semua pengajuan; balas / ubah status konsultasi |
| Kelola Topik | `/admin/topik` | CRUD topik / kategori kelas |
| Kelola Transaksi | `/admin/transaksi` | Lihat semua transaksi; verifikasi pembayaran manual jika diperlukan |
| Kelola Instruktur | `/admin/instruktur` | CRUD data instruktur / pengajar |

**Aksi CRUD Admin secara lengkap:**

| Entitas | Create | Read | Update | Delete |
|---------|--------|------|--------|--------|
| Kelas | ✅ | ✅ | ✅ | ✅ |
| Topik | ✅ | ✅ | ✅ | ✅ |
| Instruktur | ✅ | ✅ | ✅ | ✅ |
| Peserta | ❌ (via register) | ✅ | ✅ (status, role) | ✅ |
| Konsultasi | ❌ (dari user) | ✅ | ✅ (status, balasan) | ✅ |
| Transaksi | ❌ (dari sistem) | ✅ | ✅ (status manual) | ❌ |

---

## 6. UI / UX

### 6.1 Arah Desain
**Minimalis + Bento Grid + Formal**

Desain mengutamakan kesan **profesional dan terpercaya**, dengan layout yang bersih (whitespace dominan), penggunaan grid asimetris bergaya bento untuk membagi informasi secara visual, serta tipografi yang tegas dan terstruktur.

### 6.2 Palet Warna

| Nama | HEX | Fungsi |
|------|-----|--------|
| **Navy Teal** | `#0C4F6A` | Warna utama — header, CTA primer, highlight |
| **Sky Blue** | `#1A8DB2` | Aksen sekunder — badge, link, ikon aktif |
| **Cool Slate** | `#2E4A5A` | Untuk hover state, elemen depth |
| **Off White** | `#F4F7FA` | Background section, card subtle |
| **Pure White** | `#FFFFFF` | Background utama |
| **Charcoal** | `#1A1A2E` | Teks heading utama |
| **Muted Gray** | `#6B7280` | Teks sekunder, caption |
| **Light Border** | `#E2E8F0` | Garis pembatas, border card |

> ✅ **Tidak ada warna oranye.** Palet sepenuhnya berbasis biru-teal yang dingin dan profesional.

### 6.3 Tipografi
| Elemen | Font | Style |
|--------|------|-------|
| Heading | `Plus Jakarta Sans` | Bold / Semibold |
| Body | `Inter` | Regular / Medium |
| Caption / Label | `Inter` | Regular, ukuran kecil |

### 6.4 Konsep Layout (Bento Grid)
- Landing Page menggunakan **bento grid**: blok-blok kartu berukuran variatif (1x1, 2x1, 1x2, dll) yang menyusun informasi secara visual menarik namun tetap terstruktur
- Card kelas menggunakan corner radius yang konsisten (rounded-xl)
- Whitespace dimanfaatkan maksimal untuk kesan bersih dan tidak sesak
- Badge format kelas (Online/Offline) dan harga (Gratis/Berbayar) tampil dengan desain chip/pill yang minimalis

### 6.5 Prinsip Desain
- **Mobile-first** — responsif di semua ukuran layar
- **Konsisten** — komponen (button, card, badge, input) menggunakan design system yang seragam
- **Formal** — hindari elemen yang terlalu playful; font besar, whitespace luas, warna dingin
- **Accessible** — contrast ratio memadai, label form jelas, keyboard navigable
- **No emoji** — dilarang menggunakan emoji di seluruh UI (heading, label, button, card, dll)
- **Ikon fungsional saja** — ikon hanya digunakan jika memiliki fungsi navigasi atau aksi yang jelas (contoh: ikon search, ikon close). Ikon dekoratif yang tidak menambah informasi dilarang

---

## 7. Tech Stack

### Frontend
| Teknologi | Keterangan |
|-----------|------------|
| **SvelteKit** | Framework utama (SSR + SPA hybrid) |
| **TypeScript** | Type safety |
| **Tailwind CSS** | Utility-first CSS (digunakan sesuai kebutuhan) |

### Backend & Auth
| Teknologi | Keterangan |
|-----------|------------|
| **Go (Golang)** | REST API server |
| **Supabase Auth** | Autentikasi (email/password) + manajemen session |
| **JWT** | Token dari Supabase Auth, diverifikasi di Go backend |

### Database & Infrastruktur
| Teknologi | Keterangan |
|-----------|------------|
| **PostgreSQL** | Database relasional utama (via Supabase) |
| **Supabase** | Managed Postgres + Auth + Storage |

---

## 8. Arsitektur Sistem

```
[Browser / Client]
       ↓ HTTPS
[SvelteKit Frontend]
       ↓ REST API (JSON) + JWT dari Supabase
[Go Backend / API Server]
       ↓
[Supabase: PostgreSQL + Auth + Storage]
```

### Alur Autentikasi & Role Check:
1. User login via Supabase Auth di frontend → dapat JWT
2. Frontend kirim JWT di header `Authorization: Bearer <token>` ke Go backend
3. Go backend verifikasi JWT menggunakan Supabase JWT secret
4. Go backend cek role user dari JWT claims **atau** query ke tabel `profiles` di DB
5. Berdasarkan role → izinkan atau tolak akses ke resource yang diminta

### Penyimpanan Role di Database:
```sql
-- Tabel profiles (extends auth.users dari Supabase)
CREATE TABLE profiles (
  id          UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
  full_name   TEXT,
  phone       TEXT,
  avatar_url  TEXT,
  role        TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin')),
  created_at  TIMESTAMPTZ DEFAULT NOW(),
  updated_at  TIMESTAMPTZ DEFAULT NOW()
);
```

---

## 9. Rencana Pengembangan (Roadmap)

| Fase | Fokus | Estimasi |
|------|-------|----------|
| **Fase 0** | Setup repo, boilerplate FE & BE, konfigurasi Supabase, skema DB awal | 1 minggu |
| **Fase 1** | Halaman publik (Landing Page, Topik, Tentang, Kelas — tampilan saja) | 2 minggu |
| **Fase 2** | Autentikasi (register, login, logout, reset password) + middleware role | 1 minggu |
| **Fase 3** | Modul Kelas lengkap (listing, detail, pendaftaran, format Online/Offline) | 2 minggu |
| **Fase 4** | Modul Konsultasi (form pengajuan user, manajemen admin) | 1 minggu |
| **Fase 5** | Dashboard User (kelas saya, riwayat konsultasi, profil) | 1 minggu |
| **Fase 6** | Panel Admin (CRUD kelas, peserta, topik, konsultasi, transaksi) | 2 minggu |
| **Fase 7** | Integrasi Midtrans (pembayaran kelas berbayar) | 1 minggu |
| **Fase 8** | Polish UI, testing, deployment | 1 minggu |

**Total Estimasi: ~12 minggu**

---

## 10. Open Questions / Hal yang Masih Perlu Diputuskan

| Status | Item | Keputusan |
|--------|------|-----------|
| ✅ | Payment gateway | **Midtrans** (SDK open source, alur via webhook callback) |
| ✅ | Materi kelas online | **Zoom** untuk sesi live; **Supabase Storage** untuk upload PPT/PDF/file |
| ✅ | Sertifikat | **Ya**, akan ada sertifikat setelah kelas selesai |
| ✅ | Multi-kelas | **Ya**, tidak ada batas — selama jadwal tidak bentrok |
| ✅ | Role system | **2 role**: `user` (default) dan `admin` (manual assignment) |
| ⏳ | Notifikasi email/WA | Menunggu konfirmasi email perusahaan dari leader |
| ⏳ | Deployment target | Menunggu keputusan dari leader |
| ⏳ | Super Admin | Apakah perlu role `super_admin` terpisah yang bisa manage admin lain? |

---

*Dokumen ini akan diperbarui seiring perkembangan project.*