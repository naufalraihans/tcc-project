CREATE OR REPLACE FUNCTION public.handle_new_user()
RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO public.profiles (id, full_name, role)
  VALUES (NEW.id, NEW.raw_user_meta_data ->> 'full_name', 'user');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql SECURITY DEFINER;

CREATE OR REPLACE TRIGGER on_auth_user_created
  AFTER INSERT ON auth.users
  FOR EACH ROW EXECUTE FUNCTION public.handle_new_user();

CREATE EXTENSION IF NOT EXISTS moddatetime SCHEMA extensions;

CREATE TRIGGER set_updated_at_profiles   BEFORE UPDATE ON public.profiles
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_kelas      BEFORE UPDATE ON public.kelas
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_konsultasi BEFORE UPDATE ON public.konsultasi
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
CREATE TRIGGER set_updated_at_transaksi  BEFORE UPDATE ON public.transaksi
  FOR EACH ROW EXECUTE FUNCTION extensions.moddatetime(updated_at);
