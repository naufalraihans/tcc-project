INSERT INTO public.topik (nama, slug, deskripsi) VALUES
	('Ketenagalistrikan', 'ketenagalistrikan', 'Dasar dan lanjutan sistem tenaga listrik.'),
	('OHS / K3', 'k3', 'Keselamatan dan kesehatan kerja di lingkungan ketenagalistrikan.'),
	('Energi Baru Terbarukan', 'ebt', 'PLTS, Waste to Energy, dan energi terbarukan.');

INSERT INTO public.instruktur (nama, jabatan, bio) VALUES
	('Ir. Suharto, M.T., IPU', 'Direktur TCC', 'Praktisi dan pengajar bidang ketenagalistrikan.'),
	('Anna Agustina', 'MB Operasi TCC', 'Spesialis operasi dan penyelenggaraan pelatihan.');

INSERT INTO public.kelas
	(judul, slug, deskripsi, topik_id, instruktur_id, format, tipe_harga, harga, kuota, status, jadwal_mulai, jadwal_selesai, durasi_menit, lokasi)
VALUES
	('Workshop K3 dan 5S', 'workshop-k3-dan-5s',
		'Membangun budaya keselamatan kerja dan penerapan 5S di area pembangkit.',
		(SELECT id FROM public.topik WHERE slug='k3'),
		(SELECT id FROM public.instruktur WHERE nama='Anna Agustina'),
		'offline', 'berbayar', 500000, 30, 'aktif',
		NOW() + INTERVAL '10 day', NOW() + INTERVAL '10 day' + INTERVAL '8 hour', 480,
		'Graha YPK PLN, Jakarta Selatan'),
	('Pengenalan PLTS', 'pengenalan-plts',
		'Dasar desain dan operasi pemeliharaan Pembangkit Listrik Tenaga Surya.',
		(SELECT id FROM public.topik WHERE slug='ebt'),
		(SELECT id FROM public.instruktur WHERE nama='Ir. Suharto, M.T., IPU'),
		'online', 'gratis', 0, 0, 'aktif',
		NOW() + INTERVAL '5 day', NOW() + INTERVAL '5 day' + INTERVAL '4 hour', 240, NULL);
