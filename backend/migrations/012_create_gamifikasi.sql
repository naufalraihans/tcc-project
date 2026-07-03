-- Gamifikasi dashboard (update1.md §7): user_progress, aktivitas_harian, misi, misi_user, pengumuman

CREATE TABLE public.user_progress (
  user_id                 UUID PRIMARY KEY REFERENCES public.profiles(id) ON DELETE CASCADE,
  xp                      INT NOT NULL DEFAULT 0 CHECK (xp >= 0),
  level                   INT NOT NULL DEFAULT 1 CHECK (level >= 1),
  streak_saat_ini         INT NOT NULL DEFAULT 0 CHECK (streak_saat_ini >= 0),
  streak_terpanjang       INT NOT NULL DEFAULT 0 CHECK (streak_terpanjang >= 0),
  tanggal_aktif_terakhir  DATE,
  updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE public.aktivitas_harian (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id     UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  tanggal     DATE NOT NULL,
  xp_didapat  INT NOT NULL DEFAULT 0 CHECK (xp_didapat >= 0),
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, tanggal)
);
CREATE INDEX idx_aktivitas_user_tanggal ON public.aktivitas_harian (user_id, tanggal DESC);

CREATE TABLE public.misi (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  kode        TEXT NOT NULL UNIQUE,
  judul       TEXT NOT NULL,
  deskripsi   TEXT,
  tipe        TEXT NOT NULL CHECK (tipe IN ('harian','mingguan','sekali')),
  target      INT NOT NULL DEFAULT 1 CHECK (target >= 1),
  xp_reward   INT NOT NULL DEFAULT 0 CHECK (xp_reward >= 0),
  aktif       BOOLEAN NOT NULL DEFAULT TRUE,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE public.misi_user (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id     UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  misi_id     UUID NOT NULL REFERENCES public.misi(id) ON DELETE CASCADE,
  tanggal     DATE NOT NULL,
  progres     INT NOT NULL DEFAULT 0 CHECK (progres >= 0),
  selesai     BOOLEAN NOT NULL DEFAULT FALSE,
  selesai_at  TIMESTAMPTZ,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, misi_id, tanggal)
);
CREATE INDEX idx_misi_user_harian ON public.misi_user (user_id, tanggal);

CREATE TABLE public.pengumuman (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  judul       TEXT NOT NULL,
  isi         TEXT,
  tipe        TEXT NOT NULL CHECK (tipe IN ('banner','info')),
  label_aksi  TEXT,
  url_aksi    TEXT,
  urutan      INT NOT NULL DEFAULT 0,
  aktif       BOOLEAN NOT NULL DEFAULT TRUE,
  mulai       TIMESTAMPTZ,
  selesai     TIMESTAMPTZ,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_pengumuman_aktif ON public.pengumuman (aktif, urutan);

-- Seed misi default (kode stabil — dipakai hook di backend, update1.md §14)
INSERT INTO public.misi (kode, judul, deskripsi, tipe, target, xp_reward) VALUES
  ('daftar_kelas',     'Daftar satu kelas',      'Daftar minimal satu kelas hari ini',     'harian', 1, 20),
  ('buka_materi',      'Buka satu materi kelas', 'Buka minimal satu materi kelas hari ini','harian', 1, 10),
  ('ajukan_konsultasi','Ajukan konsultasi',      'Ajukan satu konsultasi hari ini',        'harian', 1, 10),
  ('lengkapi_profil',  'Lengkapi profil',        'Isi nama lengkap dan nomor telepon',     'sekali', 1, 15)
ON CONFLICT (kode) DO NOTHING;

INSERT INTO public.pengumuman (judul, isi, tipe, label_aksi, url_aksi, urutan) VALUES
  ('Katalog pelatihan 2026 dibuka', 'Jelajahi program pelatihan terbaru dari TCC ITPLN.', 'banner', 'Lihat Katalog', '/kelas', 0);
