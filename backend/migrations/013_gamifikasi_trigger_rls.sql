-- Extend register trigger: buat baris user_progress otomatis saat user baru dibuat
CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.profiles (id, full_name, role)
  VALUES (NEW.id, NEW.raw_user_meta_data ->> 'full_name', 'user');
  INSERT INTO public.user_progress (user_id) VALUES (NEW.id);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

-- Backfill user_progress untuk profil yang sudah ada
INSERT INTO public.user_progress (user_id)
SELECT id FROM public.profiles
ON CONFLICT (user_id) DO NOTHING;

-- RLS: baca hanya baris sendiri; tulis lewat backend (owner bypass RLS). Selaras 011.
ALTER TABLE public.user_progress    ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.aktivitas_harian ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.misi             ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.misi_user        ENABLE ROW LEVEL SECURITY;
ALTER TABLE public.pengumuman       ENABLE ROW LEVEL SECURITY;

CREATE POLICY "user_progress_select_own" ON public.user_progress
  FOR SELECT USING (auth.uid() = user_id);
CREATE POLICY "aktivitas_select_own" ON public.aktivitas_harian
  FOR SELECT USING (auth.uid() = user_id);
CREATE POLICY "misi_user_select_own" ON public.misi_user
  FOR SELECT USING (auth.uid() = user_id);
CREATE POLICY "misi_select_all" ON public.misi
  FOR SELECT USING (true);
CREATE POLICY "pengumuman_select_all" ON public.pengumuman
  FOR SELECT USING (true);
