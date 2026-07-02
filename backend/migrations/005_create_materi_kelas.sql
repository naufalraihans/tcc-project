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
