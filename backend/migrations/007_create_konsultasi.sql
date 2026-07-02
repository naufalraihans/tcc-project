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

CREATE INDEX idx_konsultasi_user   ON public.konsultasi(user_id);
CREATE INDEX idx_konsultasi_status ON public.konsultasi(status);
