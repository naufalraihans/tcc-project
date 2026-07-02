import { env } from '$env/dynamic/public';

const BASE = env.PUBLIC_API_BASE_URL || 'http://localhost:8080/api/v1';

type Envelope<T> = { success: boolean; data: T; error?: string; message?: string };

export async function api<T>(
	path: string,
	opts: RequestInit & { fetch?: typeof fetch } = {}
): Promise<T> {
	const { fetch: f = fetch, ...init } = opts;
	const res = await f(`${BASE}${path}`, init);
	const json = (await res.json()) as Envelope<T>;
	if (!res.ok || !json.success) {
		throw new Error(json.message || 'Gagal memuat data');
	}
	return json.data;
}
