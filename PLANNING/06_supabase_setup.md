# 🔐 Supabase Setup Guide — TCC ITPLN Web Platform

> **Dokumen ini mendefinisikan konfigurasi Supabase: Tabel, Trigger, RLS, dan Storage.**
> Versi: v1.0 | Status: Draft

---

## 1. Project Setup di Supabase

### Langkah Awal
1. Buat project baru di [supabase.com](https://supabase.com)
2. Pilih region terdekat (Singapore / `ap-southeast-1`)
3. Catat: **Project URL**, **Anon Key**, **Service Role Key**, **JWT Secret**
4. Simpan semua key ke file `.env` (lihat `ENV.md`)

---

## 2. Urutan Eksekusi Migration

> Eksekusi SQL di **Supabase Dashboard → SQL Editor** atau via CLI.
> Urutan wajib diikuti karena ada dependency antar tabel (FK).

```
001_create_profiles.sql
002_create_topik.sql
003_create_instruktur.sql
004_create_kelas.sql
005_create_materi_kelas.sql
006_create_pendaftaran.sql
007_create_konsultasi.sql
008_create_transaksi.sql
009_create_sertifikat.sql
010_create_triggers.sql
011_seed_data.sql           ← (opsional, untuk data awal)
```

---

## 3. SQL Migration Scripts

### 001 — Profiles
```sql
CREATE TABLE public.profiles (
  id          UUID PRIMARY KEY REFERENCES auth.users(id) ON DELETE CASCADE,
  full_name   TEXT,
  phone       TEXT,
  avatar_url  TEXT,
  role        TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin')),
  created_at  TIMESTAMPTZ DEFAULT NOW(),
  updated_at  TIMESTAMPTZ DEFAULT NOW()
);
```

### 002 — Topik
```sql
CREATE TABLE public.topik (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  nama        TEXT NOT NULL,
  slug        TEXT NOT NULL UNIQUE,
  deskripsi   TEXT,
  icon_url    TEXT,
  created_at  TIMESTAMPTZ DEFAULT NOW()
);
```

### 003 — Instruktur
```sql
CREATE TABLE public.instruktur (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  nama        TEXT NOT NULL,
  jabatan     TEXT,
  foto_url    TEXT,
  bio         TEXT,
  created_at  TIMESTAMPTZ DEFAULT NOW()
);
```

### 004 — Kelas
```sql
CREATE TABLE public.kelas (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  judul               TEXT NOT NULL,
  slug                TEXT NOT NULL UNIQUE,
  deskripsi           TEXT,
  silabus             TEXT,
  topik_id            UUID REFERENCES public.topik(id) ON DELETE SET NULL,
  instruktur_id       UUID REFERENCES public.instruktur(id) ON DELETE SET NULL,
  format              TEXT NOT NULL CHECK (format IN ('online', 'offline', 'hybrid')),
  tipe_harga          TEXT NOT NULL CHECK (tipe_harga IN ('gratis', 'berbayar')),
  harga               NUMERIC DEFAULT 0,
  jadwal_mulai        TIMESTAMPTZ,
  jadwal_selesai      TIMESTAMPTZ,
  durasi_menit        INT,
  kuota               INT NOT NULL DEFAULT 0,
  peserta_terdaftar   INT NOT NULL DEFAULT 0,
  status              TEXT NOT NULL DEFAULT 'aktif' CHECK (status IN ('aktif', 'penuh', 'selesai')),
  lokasi              TEXT,
  link_meeting        TEXT,
  created_at          TIMESTAMPTZ DEFAULT NOW(),
  updated_at          TIMESTAMPTZ DEFAULT NOW()
);

-- Index untuk query listing & filter
CREATE INDEX idx_kelas_topik     ON public.kelas(topik_id);
CREATE INDEX idx_kelas_format    ON public.kelas(format);
CREATE INDEX idx_kelas_status    ON public.kelas(status);
CREATE INDEX idx_kelas_harga     ON public.kelas(tipe_harga);
```

### 005 — Materi Kelas
```sql
CREATE TABLE public.materi_kelas (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  kelas_id    UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  judul       TEXT NOT NULL,
  tipe        TEXT NOT NULL CHECK (tipe IN ('file', 'link', 'video')),
  url         TEXT NOT NULL,
  urutan      INT DEFAULT 0,
  created_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_materi_kelas_id ON public.materi_kelas(kelas_id);
```

### 006 — Pendaftaran
```sql
CREATE TABLE public.pendaftaran (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  kelas_id          UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  status            TEXT NOT NULL DEFAULT 'aktif' CHECK (status IN ('aktif', 'selesai', 'dibatalkan')),
  tanggal_daftar    TIMESTAMPTZ DEFAULT NOW(),
  tanggal_selesai   TIMESTAMPTZ,

  CONSTRAINT unique_user_kelas UNIQUE (user_id, kelas_id)
);

CREATE INDEX idx_pendaftaran_user   ON public.pendaftaran(user_id);
CREATE INDEX idx_pendaftaran_kelas  ON public.pendaftaran(kelas_id);
```

### 007 — Konsultasi
```sql
CREATE TABLE public.konsultasi (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  nama_pengirim     TEXT NOT NULL,
  topik_konsultasi  TEXT NOT NULL,
  pesan             TEXT NOT NULL,
  kontak            TEXT NOT NULL,
  status            TEXT NOT NULL DEFAULT 'menunggu' CHECK (status IN ('menunggu', 'diproses', 'selesai', 'ditolak')),
  balasan           TEXT,
  admin_id          UUID REFERENCES public.profiles(id) ON DELETE SET NULL,
  created_at        TIMESTAMPTZ DEFAULT NOW(),
  updated_at        TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_konsultasi_user    ON public.konsultasi(user_id);
CREATE INDEX idx_konsultasi_status  ON public.konsultasi(status);
```

### 008 — Transaksi
```sql
CREATE TABLE public.transaksi (
  id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id             UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  kelas_id            UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  pendaftaran_id      UUID REFERENCES public.pendaftaran(id) ON DELETE SET NULL,
  midtrans_order_id   TEXT NOT NULL UNIQUE,
  midtrans_txn_id     TEXT,
  jumlah              NUMERIC NOT NULL,
  status              TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'sukses', 'gagal', 'refund')),
  metode_pembayaran   TEXT,
  created_at          TIMESTAMPTZ DEFAULT NOW(),
  updated_at          TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_transaksi_user    ON public.transaksi(user_id);
CREATE INDEX idx_transaksi_status  ON public.transaksi(status);
```

### 009 — Sertifikat
```sql
CREATE TABLE public.sertifikat (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  kelas_id          UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  pendaftaran_id    UUID NOT NULL REFERENCES public.pendaftaran(id) ON DELETE CASCADE,
  nomor_sertifikat  TEXT NOT NULL UNIQUE,
  url_sertifikat    TEXT,
  issued_at         TIMESTAMPTZ DEFAULT NOW()
);
```

### 010 — Triggers
```sql
-- Trigger 1: Auto-create profile saat user baru register
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.profiles (id, full_name, role)
  VALUES (
    NEW.id,
    NEW.raw_user_meta_data ->> 'full_name',
    'user'
  );
  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

CREATE OR REPLACE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW EXECUTE FUNCTION public.handle_new_user();

-- Trigger 2: Auto-update kolom updated_at di setiap UPDATE
-- (moddatetime = extension bawaan Supabase)
CREATE EXTENSION IF NOT EXISTS moddatetime SCHEMA extensions;

CREATE TRIGGER set_updated_at_profiles   BEFORE UPDATE ON public.profiles
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_kelas      BEFORE UPDATE ON public.kelas
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_konsultasi BEFORE UPDATE ON public.konsultasi
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_transaksi  BEFORE UPDATE ON public.transaksi
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
```

> **`peserta_terdaftar` TIDAK pakai trigger.** Sebelumnya ada 2 trigger (increment on insert + set 'penuh') — dibuang karena: (a) trigger increment + update manual di webhook Midtrans = **double-count**, dan (b) trigger tidak bisa menjamin kuota atomik (dua pendaftar barengan sama-sama lolos → overbooking).
>
> Counter & status `'penuh'` dikelola **Go backend dalam satu transaksi**, sebagai satu-satunya penulis. Pola enroll (dipakai kelas gratis & webhook sukses):
> ```sql
> -- Atomik: naikkan counter HANYA jika masih ada slot & kelas aktif.
> -- Row di-lock oleh UPDATE, jadi tidak ada race.
> UPDATE public.kelas
> SET peserta_terdaftar = peserta_terdaftar + 1,
>     status = CASE WHEN peserta_terdaftar + 1 >= kuota THEN 'penuh' ELSE status END
> WHERE id = $1 AND status = 'aktif' AND peserta_terdaftar < kuota
> RETURNING id;
> -- 0 baris  → kuota penuh / kelas tidak aktif → batalkan, kembalikan error
> -- 1 baris  → lanjut INSERT pendaftaran (UNIQUE(user_id,kelas_id) cegah duplikat)
> ```
> Saat cancel (admin) / refund: kebalikannya dalam satu transaksi —
> `UPDATE kelas SET peserta_terdaftar = peserta_terdaftar - 1, status = CASE WHEN status='penuh' THEN 'aktif' ELSE status END WHERE id=$1` lalu set `pendaftaran.status='dibatalkan'`.

---

## 4. Row Level Security (RLS)

> **Keputusan arsitektur:** frontend **tidak pernah** membaca/menulis tabel langsung — **semua data lewat Go backend** (pakai Service Role Key yang bypass RLS). Supabase client di frontend hanya untuk **Auth**. Karena itu RLS di sebagian besar tabel tidak diperlukan; RLS penuh di semua tabel cuma menambah kompleksitas tanpa proteksi tambahan (backend sudah pegang otorisasi via JWT + cek role).
>
> Kita tetap pasang RLS hanya di **`profiles`** sebagai jaring pengaman (kalau suatu saat ada akses client langsung, data profil user lain tetap terlindungi). Tabel lain: RLS off — sumber otorisasi tunggal adalah middleware Go.

```sql
-- Hanya profiles yang di-RLS. Sisanya diproteksi di layer Go (middleware auth + role).
ALTER TABLE public.profiles ENABLE ROW LEVEL SECURITY;

-- User hanya bisa lihat & edit profil sendiri (via akses client langsung, jika ada).
CREATE POLICY "profiles_select_own" ON public.profiles
  FOR SELECT USING (auth.uid() = id);

CREATE POLICY "profiles_update_own" ON public.profiles
  FOR UPDATE USING (auth.uid() = id);
```

> **Catatan:** Semua operasi read & write dilakukan Go backend dengan **Service Role Key** (bypass RLS). Otorisasi per-endpoint (siapa boleh lihat apa) ditegakkan di middleware Go — lihat `07_auth_flow.md`, bukan di RLS.

---

## 5. Storage Buckets

Konfigurasi di **Supabase Dashboard → Storage → New Bucket**

| Bucket | Akses | Isi |
|--------|-------|-----|
| `avatars` | Public | Foto profil user |
| `instruktur-foto` | Public | Foto instruktur |
| `materi-kelas` | Private | File PDF/PPT materi kelas (hanya user terdaftar) |
| `sertifikat` | Private | File PDF sertifikat kelulusan |

### Policy Storage `avatars` (Public Read)
```sql
CREATE POLICY "avatars_public_read"
  ON storage.objects FOR SELECT
  USING (bucket_id = 'avatars');

CREATE POLICY "avatars_user_upload"
  ON storage.objects FOR INSERT
  WITH CHECK (bucket_id = 'avatars' AND auth.uid() IS NOT NULL);
```

### Policy Storage `materi-kelas` (Private)
> Akses file materi dilakukan via **Go backend** yang generate signed URL — tidak perlu policy langsung dari frontend.

---

## 6. Auth Settings

Konfigurasi di **Supabase Dashboard → Authentication → Settings**

| Setting | Nilai | Keterangan |
|---------|-------|------------|
| Email Confirmations | ✅ Enable | User harus konfirmasi email setelah register |
| Secure Email Change | ✅ Enable | Konfirmasi email lama saat ganti email |
| Site URL | `https://tcc-itpln.id` | URL redirect setelah konfirmasi email |
| Redirect URLs | `http://localhost:5173/**` | Tambahkan untuk dev environment |
| JWT Expiry | `3600` (1 jam) | Durasi token aktif |
| Refresh Token Rotation | ✅ Enable | Auto-rotate refresh token |

---

## 7. Checklist Setup Supabase

```
[ ] Buat project Supabase (region: Singapore)
[ ] Catat semua key → isi ke .env
[ ] Jalankan migration 001 s/d 010 via SQL Editor
[ ] Aktifkan RLS pada tabel profiles (+ policy select/update own)
[ ] Buat Storage buckets (avatars, instruktur-foto, materi-kelas, sertifikat)
[ ] Set Storage policies untuk bucket public
[ ] Konfigurasi Auth settings (email confirm, site URL, redirect URLs)
[ ] Test trigger: register user baru → cek apakah profiles ter-insert otomatis
[ ] Test enroll atomik: 2 pendaftaran barengan di kelas sisa 1 slot → hanya 1 lolos
```

---

*Dokumen ini akan diperbarui seiring perkembangan konfigurasi project.*
