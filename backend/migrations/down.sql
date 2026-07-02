DROP TRIGGER IF EXISTS on_auth_user_created ON auth.users;
DROP FUNCTION IF EXISTS public.handle_new_user();

DROP TABLE IF EXISTS public.sertifikat CASCADE;
DROP TABLE IF EXISTS public.transaksi CASCADE;
DROP TABLE IF EXISTS public.konsultasi CASCADE;
DROP TABLE IF EXISTS public.pendaftaran CASCADE;
DROP TABLE IF EXISTS public.materi_kelas CASCADE;
DROP TABLE IF EXISTS public.kelas CASCADE;
DROP TABLE IF EXISTS public.instruktur CASCADE;
DROP TABLE IF EXISTS public.topik CASCADE;
DROP TABLE IF EXISTS public.profiles CASCADE;
