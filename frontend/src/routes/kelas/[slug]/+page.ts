import { api } from '$lib/api';
import { error } from '@sveltejs/kit';
import type { Kelas } from '$lib/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
	try {
		const kelas = await api<Kelas>(`/kelas/${params.slug}`, { fetch });
		return { kelas };
	} catch (e) {
		error(404, (e as Error).message || 'Kelas tidak ditemukan');
	}
};
