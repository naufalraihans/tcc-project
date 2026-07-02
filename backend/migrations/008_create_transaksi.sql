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

CREATE INDEX idx_transaksi_user   ON public.transaksi(user_id);
CREATE INDEX idx_transaksi_status ON public.transaksi(status);
