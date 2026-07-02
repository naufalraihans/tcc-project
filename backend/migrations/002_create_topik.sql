CREATE TABLE public.topik (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  nama        TEXT NOT NULL,
  slug        TEXT NOT NULL UNIQUE,
  deskripsi   TEXT,
  icon_url    TEXT,
  created_at  TIMESTAMPTZ DEFAULT NOW()
);
