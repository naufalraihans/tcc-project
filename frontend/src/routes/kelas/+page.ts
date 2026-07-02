import { api } from '$lib/api';
import type { PagedKelas, Topik } from '$lib/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, url }) => {
	const qs = new URLSearchParams();
	for (const key of ['topik', 'format', 'harga', 'page']) {
		const v = url.searchParams.get(key);
		if (v) qs.set(key, v);
	}
	try {
		const [kelas, topik] = await Promise.all([
			api<PagedKelas>(`/kelas?${qs.toString()}`, { fetch }),
			api<Topik[]>('/topik', { fetch })
		]);
		return { kelas, topik, error: null as string | null };
	} catch (e) {
		return { kelas: null, topik: [] as Topik[], error: (e as Error).message };
	}
};
