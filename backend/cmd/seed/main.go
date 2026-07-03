package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"

	"tcc-itpln/backend/config"
)

const (
	adminID = "00000000-0000-0000-0000-000000000001"
	userID  = "00000000-0000-0000-0000-000000000002"

	adminPass = "admin123"
	userPass  = "user123"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer conn.Close(ctx)

	if err := seed(ctx, conn); err != nil {
		log.Fatalf("seed: %v", err)
	}

	fmt.Println("\n=========== SEED OK ===========")
	fmt.Printf("Admin  ->  admin@tcc.local  /  %s\n", adminPass)
	fmt.Printf("User   ->  user@tcc.local   /  %s\n", userPass)
	fmt.Println("Atau daftar akun sendiri di /auth/register")
	fmt.Println("===============================")
}

func seed(ctx context.Context, conn *pgx.Conn) error {
	adminHash := hash(adminPass)
	userHash := hash(userPass)

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	stmts := []string{
		`TRUNCATE public.sertifikat, public.transaksi, public.konsultasi, public.pendaftaran,
			public.materi_kelas, public.kelas, public.instruktur, public.topik, public.profiles
			RESTART IDENTITY CASCADE`,
		`DELETE FROM auth.users`,

		`INSERT INTO auth.users (id, email, encrypted_password, raw_user_meta_data) VALUES
			('` + adminID + `', 'admin@tcc.local', '` + adminHash + `', '{"full_name":"Admin TCC"}'),
			('` + userID + `', 'user@tcc.local', '` + userHash + `', '{"full_name":"Budi Santoso"}')`,

		`UPDATE public.profiles SET role='admin', phone='0811000001' WHERE id='` + adminID + `'`,
		`UPDATE public.profiles SET phone='0812000002' WHERE id='` + userID + `'`,

		`INSERT INTO public.topik (nama, slug, deskripsi) VALUES
			('Ketenagalistrikan','ketenagalistrikan','Dasar dan lanjutan sistem tenaga listrik.'),
			('OHS / K3','k3','Keselamatan dan kesehatan kerja di ketenagalistrikan.'),
			('Energi Baru Terbarukan','ebt','PLTS, Waste to Energy, dan energi terbarukan.'),
			('Engineering','engineering','Rekayasa dan pemeliharaan pembangkit.')`,

		`INSERT INTO public.instruktur (nama, jabatan, bio) VALUES
			('Ir. Suharto, M.T., IPU','Direktur TCC','Praktisi dan pengajar bidang ketenagalistrikan.'),
			('Anna Agustina','MB Operasi TCC','Spesialis operasi dan penyelenggaraan pelatihan.'),
			('Dr. Ahmad Fauzi','Ahli K3 Kelistrikan','Berpengalaman 15 tahun di bidang K3.')`,

		`INSERT INTO public.kelas
			(judul, slug, deskripsi, silabus, topik_id, instruktur_id, format, tipe_harga, harga, kuota, status, jadwal_mulai, jadwal_selesai, durasi_menit, lokasi) VALUES
			('Workshop K3 dan 5S','workshop-k3-dan-5s',
				'Membangun budaya keselamatan kerja dan penerapan 5S di area pembangkit.',
				'HIRARC, JSA, APD, Behavior Based Safety, penerapan 5S.',
				(SELECT id FROM public.topik WHERE slug='k3'),
				(SELECT id FROM public.instruktur WHERE nama='Anna Agustina'),
				'offline','berbayar',500000,30,'aktif',
				NOW()+INTERVAL '10 day', NOW()+INTERVAL '10 day'+INTERVAL '8 hour',480,'Graha YPK PLN, Jakarta Selatan'),
			('Pengenalan PLTS','pengenalan-plts',
				'Dasar desain dan operasi pemeliharaan Pembangkit Listrik Tenaga Surya.',
				'Komponen PLTS, perhitungan kapasitas, O&M dasar.',
				(SELECT id FROM public.topik WHERE slug='ebt'),
				(SELECT id FROM public.instruktur WHERE nama='Ir. Suharto, M.T., IPU'),
				'online','gratis',0,0,'aktif',
				NOW()+INTERVAL '5 day', NOW()+INTERVAL '5 day'+INTERVAL '4 hour',240,NULL),
			('Basic Theory Power Plant','basic-theory-power-plant',
				'Konsep dasar pembangkit: termodinamika, mekanika fluida, instrumentasi dan kontrol.',
				'Termodinamika, mekanika fluida, grid code, instrumen & kontrol.',
				(SELECT id FROM public.topik WHERE slug='engineering'),
				(SELECT id FROM public.instruktur WHERE nama='Ir. Suharto, M.T., IPU'),
				'hybrid','berbayar',350000,25,'aktif',
				NOW()+INTERVAL '20 day', NOW()+INTERVAL '20 day'+INTERVAL '6 hour',360,'Kampus ITPLN')`,

		`INSERT INTO public.pendaftaran (user_id, kelas_id, status) VALUES
			('` + userID + `', (SELECT id FROM public.kelas WHERE slug='pengenalan-plts'),'aktif')`,
		`UPDATE public.kelas SET peserta_terdaftar=1 WHERE slug='pengenalan-plts'`,

		`INSERT INTO public.konsultasi (user_id, nama_pengirim, topik_konsultasi, pesan, kontak, status) VALUES
			('` + userID + `','Budi Santoso','Pelatihan K3 untuk 50 karyawan',
				'Kami membutuhkan in-house training K3 untuk 50 karyawan pabrik kami di Bekasi.',
				'budi@example.com','menunggu')`,
	}

	for _, s := range stmts {
		if _, err := tx.Exec(ctx, s); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

func hash(pw string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("hash: %v", err)
	}
	return string(h)
}
