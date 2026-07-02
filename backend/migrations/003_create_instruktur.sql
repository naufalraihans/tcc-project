CREATE TABLE public.instruktur (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  nama        TEXT NOT NULL,
  jabatan     TEXT,
  foto_url    TEXT,
  bio         TEXT,
  created_at  TIMESTAMPTZ DEFAULT NOW()
);
