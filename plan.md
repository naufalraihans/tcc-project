# 📋 Project Plan — TCC ITPLN Web Platform

> **Dokumen ini adalah perencanaan awal (v1.2) untuk pembangunan platform web TCC ITPLN.**
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

## 3. Role & Target Pengguna

Platform ini memiliki **2 role utama:**

| Role | Deskripsi |
|------|-----------|
| **User (Peserta)** | Pengguna terdaftar yang dapat mendaftar kelas, mengikuti pelatihan, dan mengajukan konsultasi |
| **Admin** | Staff TCC yang mengelola konten, kelas, peserta, dan sesi konsultasi melalui panel admin |

> Guest (belum login) hanya dapat mengakses halaman publik dan tidak bisa mendaftar kelas.

---

## 4. Struktur Kelas / Pelatihan

Kelas di TCC memiliki dua dimensi utama:

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

### 4.4 Aturan Pendaftaran Kelas
- **Multi-kelas diizinkan** — 1 user dapat mendaftar banyak kelas sekaligus tanpa batas
- **Validasi konflik jadwal** — sistem wajib mencegah pendaftaran jika jadwal kelas baru bentrok dengan kelas yang sudah diikuti user

### 4.5 Alur Pembayaran (Kelas Berbayar)
1. User klik "Daftar" pada kelas berbayar
2. Sistem membuat **order/transaksi pending** di database
3. User diarahkan ke halaman pembayaran **Midtrans** (Snap UI)
4. Setelah pembayaran sukses → Midtrans kirim **webhook callback** ke Go backend
5. Backend verifikasi callback → update status transaksi → user otomatis masuk kelas
6. User mendapat konfirmasi (notifikasi — *TBD*)

---

## 5. Fitur & Halaman Web

### 5.1 Halaman Publik (Tanpa Login)

#### Beranda
- **URL:** `/`
- Hero section dengan tagline & CTA (daftar / lihat program)
- Preview kelas unggulan (Online & Offline)
- Topik pelatihan yang tersedia
- Tentang TCC (ringkas)
- Statistik platform (peserta, kelas, instruktur)
- CTA konsultasi

#### Kelas
- **URL:** `/kelas`
- Daftar semua kelas dengan filter: Topik, Format (Online/Offline), Harga (Gratis/Berbayar)
- Kartu kelas: judul, format badge, instruktur, jadwal, harga, kuota sisa
- **Detail Kelas:** `/kelas/[slug]`
  - Deskripsi lengkap, silabus, instruktur, jadwal
  - Tombol daftar kelas (redirect ke login jika belum login)

#### Konsultasi
- **URL:** `/konsultasi`
- Penjelasan layanan konsultasi TCC
- Form pengajuan konsultasi (nama, topik, pesan, kontak)

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
| Register | `/auth/register` | Pendaftaran akun baru (role default: User) |
| Lupa Password | `/auth/lupa-password` | Reset password via email |

---

### 5.3 Halaman User / Peserta (Requires Login)

| Halaman | URL | Keterangan |
|---------|-----|------------|
| Dashboard | `/dashboard` | Ringkasan kelas yang diikuti, jadwal, notifikasi |
| Kelas Saya | `/dashboard/kelas` | Daftar kelas yang sedang & sudah diikuti |
| Detail Kelas | `/dashboard/kelas/[id]` | Akses materi atau info kelas |
| Konsultasi | `/dashboard/konsultasi` | Riwayat pengajuan konsultasi |
| Profil | `/dashboard/profil` | Kelola data diri |

---

### 5.4 Halaman Admin (Requires Admin Role)

| Halaman | URL | Keterangan |
|---------|-----|------------|
| Dashboard Admin | `/admin/dashboard` | Statistik global platform |
| Kelola Kelas | `/admin/kelas` | CRUD kelas (Online & Offline, Gratis & Berbayar) |
| Kelola Peserta | `/admin/peserta` | Lihat, kelola, dan verifikasi akun peserta |
| Kelola Konsultasi | `/admin/konsultasi` | Lihat & respon pengajuan konsultasi |
| Kelola Topik | `/admin/topik` | CRUD topik / kategori kelas |

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
- Beranda menggunakan **bento grid**: blok-blok kartu berukuran variatif (1x1, 2x1, 1x2, dll) yang menyusun informasi secara visual menarik namun tetap terstruktur
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

**Alur Autentikasi:**
1. User login via Supabase Auth di frontend → dapat JWT
2. Frontend kirim JWT di header `Authorization: Bearer <token>` ke Go backend
3. Go backend verifikasi JWT menggunakan Supabase JWT secret
4. Go backend cek role user dari JWT claims atau DB

---

## 9. Rencana Pengembangan (Roadmap)

| Fase | Fokus | Estimasi |
|------|-------|----------|
| **Fase 0** | Setup repo, boilerplate FE & BE, konfigurasi Supabase | 1 minggu |
| **Fase 1** | Halaman publik (Beranda, Topik, Tentang, Kelas) | 2 minggu |
| **Fase 2** | Autentikasi (Supabase Auth: login, register, reset password) | 1 minggu |
| **Fase 3** | Modul Kelas lengkap (listing, detail, pendaftaran, format Online/Offline) | 2 minggu |
| **Fase 4** | Modul Konsultasi (form pengajuan, manajemen admin) | 1 minggu |
| **Fase 5** | Dashboard User & Panel Admin | 2 minggu |
| **Fase 6** | Polish UI, testing, deployment | 1 minggu |

**Total Estimasi: ~10 minggu**

---

## 10. Open Questions / Hal yang Masih Perlu Diputuskan

| Status | Item | Keputusan |
|--------|------|-----------|
| ✅ | Payment gateway | **Midtrans** (SDK open source, alur via webhook callback) |
| ✅ | Materi kelas online | **Zoom** untuk sesi live; **Supabase Storage** untuk upload PPT/PDF/file |
| ✅ | Sertifikat | **Ya**, akan ada sertifikat setelah kelas selesai |
| ✅ | Multi-kelas | **Ya**, tidak ada batas — selama jadwal tidak bentrok |
| ⏳ | Notifikasi email/WA | Menunggu konfirmasi email perusahaan dari leader |
| ⏳ | Deployment target | Menunggu keputusan dari leader |

---

*Dokumen ini akan diperbarui seiring perkembangan project.*