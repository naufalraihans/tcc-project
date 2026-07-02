export type MiniTopik = { id: string; nama: string; slug: string };
export type MiniInstruktur = { id: string; nama: string; foto_url: string };

export type Topik = {
	id: string;
	nama: string;
	slug: string;
	deskripsi: string;
	icon_url: string;
	jumlah_kelas?: number;
	created_at: string;
};

export type KelasListItem = {
	id: string;
	judul: string;
	slug: string;
	topik: MiniTopik | null;
	instruktur: MiniInstruktur | null;
	format: string;
	tipe_harga: string;
	harga: number;
	jadwal_mulai: string | null;
	jadwal_selesai: string | null;
	kuota: number;
	peserta_terdaftar: number;
	status: string;
};

export type Pagination = { page: number; limit: number; total: number; total_pages: number };
export type PagedKelas = { items: KelasListItem[]; pagination: Pagination };

export type Kelas = KelasListItem & {
	deskripsi: string;
	silabus: string;
	durasi_menit: number;
	lokasi: string;
	link_meeting: string;
	created_at: string;
	updated_at: string;
};

export type Instruktur = {
	id: string;
	nama: string;
	jabatan: string;
	foto_url: string;
	bio: string;
	created_at: string;
};

export type Profile = {
	id: string;
	full_name: string;
	phone: string;
	avatar_url: string;
	role: string;
	email?: string;
	created_at: string;
};

export type PendaftaranItem = {
	pendaftaran_id: string;
	kelas: { id: string; judul: string; slug: string; format: string };
	user?: { id: string; full_name: string };
	status: string;
	tanggal_daftar: string;
};

export type Konsultasi = {
	id: string;
	user_id: string;
	nama_pengirim: string;
	topik_konsultasi: string;
	pesan: string;
	kontak: string;
	status: string;
	balasan: string;
	admin_id: string | null;
	created_at: string;
	updated_at: string;
};

export type Transaksi = {
	id: string;
	user_id: string;
	kelas_id: string;
	pendaftaran_id: string | null;
	midtrans_order_id: string;
	midtrans_txn_id: string;
	jumlah: number;
	status: string;
	metode_pembayaran: string;
	created_at: string;
	updated_at: string;
};
