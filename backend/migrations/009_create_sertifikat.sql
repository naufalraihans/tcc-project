CREATE TABLE public.sertifikat (
  id                UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id           UUID NOT NULL REFERENCES public.profiles(id) ON DELETE CASCADE,
  kelas_id          UUID NOT NULL REFERENCES public.kelas(id) ON DELETE CASCADE,
  pendaftaran_id    UUID NOT NULL REFERENCES public.pendaftaran(id) ON DELETE CASCADE,
  nomor_sertifikat  TEXT NOT NULL UNIQUE,
  url_sertifikat    TEXT,
  issued_at         TIMESTAMPTZ DEFAULT NOW()
);
