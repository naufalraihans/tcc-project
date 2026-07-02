import { supabase } from './supabase';
import { api } from './api';

export async function apiAuth<T>(path: string, opts: RequestInit = {}): Promise<T> {
	const { data } = await supabase.auth.getSession();
	const token = data.session?.access_token;
	const headers = new Headers(opts.headers);
	if (token) headers.set('Authorization', `Bearer ${token}`);
	if (opts.body && !headers.has('Content-Type')) headers.set('Content-Type', 'application/json');
	return api<T>(path, { ...opts, headers });
}
