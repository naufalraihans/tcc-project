# 📡 API Contract — TCC ITPLN Web Platform

> **Dokumen ini mendefinisikan seluruh endpoint REST API Go Backend.**
> Versi: v1.0 | Status: Draft | Base URL: `https://api.tcc-itpln.id/api/v1`

---

## Konvensi Umum

### Auth Header
```
Authorization: Bearer <supabase_jwt_token>
```

### Format Response Sukses
```json
{
  "success": true,
  "data": { ... },
  "message": "OK"
}
```

### Format Response Error
```json
{
  "success": false,
  "error": "KODE_ERROR",
  "message": "Penjelasan error yang human-readable"
}
```

### Kode Error Umum
| Kode | HTTP | Keterangan |
|------|------|------------|
| `UNAUTHORIZED` | 401 | JWT tidak ada atau expired |
| `FORBIDDEN` | 403 | Role tidak cukup |
| `NOT_FOUND` | 404 | Resource tidak ditemukan |
| `VALIDATION_ERROR` | 422 | Input tidak valid |
| `INTERNAL_ERROR` | 500 | Server error |

### Level Akses
- 🌐 **Public** — Tanpa token
- 🔐 **User** — JWT valid (role: `user` atau `admin`)
- 🛡️ **Admin** — JWT valid + role `admin`

---

## 1. Auth

> Autentikasi dihandle langsung oleh **Supabase Auth** di frontend.
> Go backend hanya memverifikasi JWT yang diterima.

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/auth/me` | 🔐 User | Ambil profil user yang sedang login |
| `PUT` | `/auth/profile` | 🔐 User | Update profil (nama, telepon, avatar) |

### `GET /auth/me`
**Response:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "full_name": "Budi Santoso",
    "email": "budi@example.com",
    "phone": "08123456789",
    "avatar_url": "https://...",
    "role": "user",
    "created_at": "2026-07-01T10:00:00Z"
  }
}
```

### `PUT /auth/profile`
**Request Body:**
```json
{
  "full_name": "Budi Santoso",
  "phone": "08123456789",
  "avatar_url": "https://..."
}
```

---

## 2. Topik

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/topik` | 🌐 Public | Ambil semua topik |
| `GET` | `/topik/:slug` | 🌐 Public | Detail topik beserta kelas-kelasnya |
| `POST` | `/admin/topik` | 🛡️ Admin | Buat topik baru |
| `PUT` | `/admin/topik/:id` | 🛡️ Admin | Update topik |
| `DELETE` | `/admin/topik/:id` | 🛡️ Admin | Hapus topik |

### `GET /topik`
**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "nama": "Kelistrikan Industri",
      "slug": "kelistrikan-industri",
      "deskripsi": "...",
      "icon_url": "https://...",
      "jumlah_kelas": 5
    }
  ]
}
```

### `POST /admin/topik`
**Request Body:**
```json
{
  "nama": "Smart Grid",
  "slug": "smart-grid",
  "deskripsi": "Pelatihan terkait sistem smart grid modern",
  "icon_url": "https://..."
}
```

---

## 3. Instruktur

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/instruktur` | 🌐 Public | Ambil semua instruktur |
| `GET` | `/instruktur/:id` | 🌐 Public | Detail instruktur |
| `POST` | `/admin/instruktur` | 🛡️ Admin | Tambah instruktur |
| `PUT` | `/admin/instruktur/:id` | 🛡️ Admin | Update instruktur |
| `DELETE` | `/admin/instruktur/:id` | 🛡️ Admin | Hapus instruktur |

### `POST /admin/instruktur`
**Request Body:**
```json
{
  "nama": "Dr. Ahmad Fauzi",
  "jabatan": "Ahli K3 Kelistrikan",
  "foto_url": "https://...",
  "bio": "Berpengalaman 15 tahun di bidang..."
}
```

---

## 4. Kelas

### 4.1 Endpoint Publik

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/kelas` | 🌐 Public | List kelas dengan filter & pagination |
| `GET` | `/kelas/:slug` | 🌐 Public | Detail kelas |

### `GET /kelas`
**Query Params:**

| Param | Tipe | Contoh | Keterangan |
|-------|------|--------|------------|
| `topik` | string | `kelistrikan-industri` | Filter by topik slug |
| `format` | string | `online` / `offline` / `hybrid` | Filter format kelas |
| `harga` | string | `gratis` / `berbayar` | Filter tipe harga |
| `status` | string | `aktif` | Filter status (default: aktif) |
| `page` | int | `1` | Halaman (default: 1) |
| `limit` | int | `12` | Jumlah per halaman (default: 12) |

**Response:**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "uuid",
        "judul": "Workshop K3 Listrik Dasar",
        "slug": "workshop-k3-listrik-dasar",
        "topik": { "id": "uuid", "nama": "K3", "slug": "k3" },
        "instruktur": { "id": "uuid", "nama": "Dr. Ahmad", "foto_url": "..." },
        "format": "offline",
        "tipe_harga": "berbayar",
        "harga": 500000,
        "jadwal_mulai": "2026-08-01T08:00:00Z",
        "jadwal_selesai": "2026-08-01T17:00:00Z",
        "kuota": 30,
        "peserta_terdaftar": 12,
        "status": "aktif"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 12,
      "total": 45,
      "total_pages": 4
    }
  }
}
```

### `GET /kelas/:slug`
**Response:**
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "judul": "Workshop K3 Listrik Dasar",
    "slug": "workshop-k3-listrik-dasar",
    "deskripsi": "...",
    "silabus": "...",
    "topik": { ... },
    "instruktur": { ... },
    "format": "offline",
    "tipe_harga": "berbayar",
    "harga": 500000,
    "jadwal_mulai": "2026-08-01T08:00:00Z",
    "jadwal_selesai": "2026-08-01T17:00:00Z",
    "durasi_menit": 480,
    "kuota": 30,
    "peserta_terdaftar": 12,
    "status": "aktif",
    "link_meeting": null
  }
}
```

### 4.2 Endpoint Admin — Kelola Kelas

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `POST` | `/admin/kelas` | 🛡️ Admin | Buat kelas baru |
| `PUT` | `/admin/kelas/:id` | 🛡️ Admin | Update kelas |
| `DELETE` | `/admin/kelas/:id` | 🛡️ Admin | Hapus kelas |
| `PATCH` | `/admin/kelas/:id/status` | 🛡️ Admin | Ubah status kelas saja |

### `POST /admin/kelas`
**Request Body:**
```json
{
  "judul": "Workshop K3 Listrik Dasar",
  "slug": "workshop-k3-listrik-dasar",
  "deskripsi": "Deskripsi lengkap...",
  "silabus": "Outline materi...",
  "topik_id": "uuid",
  "instruktur_id": "uuid",
  "format": "offline",
  "tipe_harga": "berbayar",
  "harga": 500000,
  "jadwal_mulai": "2026-08-01T08:00:00Z",
  "jadwal_selesai": "2026-08-01T17:00:00Z",
  "durasi_menit": 480,
  "kuota": 30,
  "lokasi": "Graha YPK PLN, Jakarta Selatan",
  "link_meeting": null
}
```

### `PATCH /admin/kelas/:id/status`
**Request Body:**
```json
{ "status": "selesai" }
```

---

## 5. Materi Kelas

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/kelas/:id/materi` | 🔐 User | Ambil materi (user harus terdaftar) |
| `POST` | `/admin/kelas/:id/materi` | 🛡️ Admin | Tambah materi ke kelas |
| `PUT` | `/admin/materi/:id` | 🛡️ Admin | Update materi |
| `DELETE` | `/admin/materi/:id` | 🛡️ Admin | Hapus materi |

### `GET /kelas/:id/materi`
> Backend wajib verifikasi bahwa user memiliki `pendaftaran` dengan status `aktif` **atau** `selesai` pada kelas ini (materi tetap bisa diakses setelah kelas selesai).

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": "uuid",
      "judul": "Modul 1 - Pengenalan K3",
      "tipe": "file",
      "url": "https://storage.supabase.co/...",
      "urutan": 1
    }
  ]
}
```

---

## 6. Pendaftaran Kelas

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `POST` | `/pendaftaran` | 🔐 User | Daftar kelas (gratis atau mulai flow berbayar) |
| `GET` | `/pendaftaran/saya` | 🔐 User | Daftar kelas yang diikuti user |
| `GET` | `/pendaftaran/saya/:id` | 🔐 User | Detail pendaftaran spesifik |
| `GET` | `/admin/pendaftaran` | 🛡️ Admin | Semua pendaftaran (dengan filter) |
| `PATCH` | `/admin/pendaftaran/:id/status` | 🛡️ Admin | Update status pendaftaran |

### `POST /pendaftaran`
**Request Body:**
```json
{ "kelas_id": "uuid" }
```

**Response (kelas gratis):**
```json
{
  "success": true,
  "data": {
    "type": "gratis",
    "pendaftaran_id": "uuid",
    "message": "Berhasil mendaftar kelas"
  }
}
```

**Response (kelas berbayar):**
```json
{
  "success": true,
  "data": {
    "type": "berbayar",
    "transaksi_id": "uuid",
    "midtrans_snap_token": "abc123...",
    "redirect_url": "https://app.midtrans.com/snap/v2/vtweb/abc123"
  }
}
```

**Error — Konflik Jadwal:**
```json
{
  "success": false,
  "error": "SCHEDULE_CONFLICT",
  "message": "Jadwal bentrok dengan kelas 'Workshop Smart Grid' (01 Agt 2026, 08.00–17.00)"
}
```

### `GET /pendaftaran/saya`
**Query Params:** `status` (aktif / selesai / semua), `page`, `limit`

**Response:**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "pendaftaran_id": "uuid",
        "kelas": { "id": "uuid", "judul": "...", "slug": "...", "format": "online" },
        "status": "aktif",
        "tanggal_daftar": "2026-07-15T10:00:00Z"
      }
    ]
  }
}
```

---

## 7. Konsultasi

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `POST` | `/konsultasi` | 🔐 User | Ajukan konsultasi baru |
| `GET` | `/konsultasi/saya` | 🔐 User | Riwayat konsultasi milik user |
| `GET` | `/admin/konsultasi` | 🛡️ Admin | Semua pengajuan konsultasi |
| `GET` | `/admin/konsultasi/:id` | 🛡️ Admin | Detail konsultasi |
| `PATCH` | `/admin/konsultasi/:id` | 🛡️ Admin | Update status + isi balasan |

### `POST /konsultasi`
**Request Body:**
```json
{
  "nama_pengirim": "Budi Santoso",
  "topik_konsultasi": "Konsultasi K3 untuk Industri Manufaktur",
  "pesan": "Kami membutuhkan pelatihan K3 untuk 50 karyawan...",
  "kontak": "budi@company.com"
}
```

### `PATCH /admin/konsultasi/:id`
**Request Body:**
```json
{
  "status": "selesai",
  "balasan": "Terima kasih atas pengajuan Anda. Tim kami akan menghubungi Anda dalam 1x24 jam..."
}
```

---

## 8. Transaksi

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/transaksi/saya` | 🔐 User | Riwayat transaksi user |
| `POST` | `/webhook/midtrans` | 🌐 Public* | Callback dari Midtrans |
| `GET` | `/admin/transaksi` | 🛡️ Admin | Semua transaksi |
| `PATCH` | `/admin/transaksi/:id/status` | 🛡️ Admin | Update status transaksi manual |

> *`/webhook/midtrans` — Public tapi diverifikasi dengan **Midtrans Signature Key**, bukan JWT.

### `POST /webhook/midtrans`
**Payload dari Midtrans (contoh):**
```json
{
  "order_id": "TCC-TXN-uuid",
  "transaction_id": "midtrans-txn-id",
  "transaction_status": "settlement",
  "payment_type": "bank_transfer",
  "gross_amount": "500000.00",
  "signature_key": "hash_sha512..."
}
```

**Alur di Backend:**
1. Verifikasi `signature_key`
2. Cek idempotency (jika sudah `sukses`, return 200, stop)
3. Update `transaksi.status` → `sukses`
4. Enroll atomik: `UPDATE kelas +1` dgn guard kuota → `INSERT pendaftaran` (lihat `08` §6.3; jika kuota penuh → refund)
5. Update `transaksi.pendaftaran_id`

---

## 9. Sertifikat

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/sertifikat/saya` | 🔐 User | Daftar sertifikat milik user |
| `GET` | `/sertifikat/:nomor` | 🌐 Public | Verifikasi sertifikat by nomor seri |
| `POST` | `/admin/sertifikat` | 🛡️ Admin | Terbitkan sertifikat untuk pendaftaran |

### `GET /sertifikat/:nomor`
> Endpoint publik untuk verifikasi keaslian sertifikat.

**Response:**
```json
{
  "success": true,
  "data": {
    "nomor_sertifikat": "TCC-2026-0001",
    "nama_penerima": "Budi Santoso",
    "kelas": "Workshop K3 Listrik Dasar",
    "issued_at": "2026-08-05T10:00:00Z",
    "valid": true
  }
}
```

---

## 10. Dashboard & Statistik Admin

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/admin/dashboard/stats` | 🛡️ Admin | Statistik global platform |

### `GET /admin/dashboard/stats`
**Response:**
```json
{
  "success": true,
  "data": {
    "total_user": 150,
    "total_kelas_aktif": 8,
    "total_pendaftaran": 312,
    "total_konsultasi_menunggu": 5,
    "pendapatan_bulan_ini": 12500000,
    "kelas_terpopuler": [
      { "judul": "Workshop K3", "peserta": 28 }
    ]
  }
}
```

---

## 11. Kelola Peserta (Admin)

| Method | Endpoint | Akses | Keterangan |
|--------|----------|-------|------------|
| `GET` | `/admin/peserta` | 🛡️ Admin | List semua user terdaftar |
| `GET` | `/admin/peserta/:id` | 🛡️ Admin | Detail profil user |
| `PATCH` | `/admin/peserta/:id/role` | 🛡️ Admin | Ubah role user |
| `DELETE` | `/admin/peserta/:id` | 🛡️ Admin | Nonaktifkan / hapus akun user |

### `PATCH /admin/peserta/:id/role`
**Request Body:**
```json
{ "role": "admin" }
```

---

*Dokumen ini akan diperbarui seiring finalisasi development.*
