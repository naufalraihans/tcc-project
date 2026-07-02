export function rupiah(n: number): string {
	if (!n || n <= 0) return 'Gratis';
	return 'Rp ' + n.toLocaleString('id-ID');
}

export function harga(tipe: string, n: number): string {
	return tipe === 'gratis' ? 'Gratis' : rupiah(n);
}

export function tanggal(iso: string | null): string {
	if (!iso) return 'Jadwal menyusul';
	return new Date(iso).toLocaleDateString('id-ID', {
		day: 'numeric',
		month: 'long',
		year: 'numeric'
	});
}
