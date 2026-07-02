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

CREATE INDEX idx_kelas_topik  ON public.kelas(topik_id);
CREATE INDEX idx_kelas_format ON public.kelas(format);
CREATE INDEX idx_kelas_status ON public.kelas(status);
CREATE INDEX idx_kelas_harga  ON public.kelas(tipe_harga);
