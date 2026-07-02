import { api } from '$lib/api';
import type { Topik } from '$lib/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }) => {
	try {
		const topik = await api<Topik[]>('/topik', { fetch });
		return { topik, error: null as string | null };
	} catch (e) {
		return { topik: [] as Topik[], error: (e as Error).message };
	}
};
