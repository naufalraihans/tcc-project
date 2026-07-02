CREATE TABLE public.pendaftaran (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  kelas_id          UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  status            TEXT NOT NULL DEFAULT 'aktif' CHECK (status IN ('aktif', 'selesai', 'dibatalkan')),
  tanggal_daftar    TIMESTAMPTZ DEFAULT NOW(),
  tanggal_selesai   TIMESTAMPTZ,
  CONSTRAINT unique_user_kelas UNIQUE (user_id, kelas_id)
);

CREATE INDEX idx_pendaftaran_user  ON public.pendaftaran(user_id);
CREATE INDEX idx_pendaftaran_kelas ON public.pendaftaran(kelas_id);
