import { api } from './api';

export type AuthResult = { token: string; role: string; full_name: string };

const KEY = 'token';

export function currentToken(): string {
	if (typeof localStorage === 'undefined') return '';
	return localStorage.getItem(KEY) ?? '';
}

export function clearToken(): void {
	if (typeof localStorage !== 'undefined') localStorage.removeItem(KEY);
}

export async function login(email: string, password: string): Promise<AuthResult> {
	const res = await api<AuthResult>('/auth/login', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ email, password })
	});
	localStorage.setItem(KEY, res.token);
	return res;
}

export async function register(
	full_name: string,
	email: string,
	password: string
): Promise<AuthResult> {
	const res = await api<AuthResult>('/auth/register', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ full_name, email, password })
	});
	localStorage.setItem(KEY, res.token);
	return res;
}

export async function apiAuth<T>(path: string, opts: RequestInit = {}): Promise<T> {
	const headers = new Headers(opts.headers);
	const token = currentToken();
	if (token) headers.set('Authorization', `Bearer ${token}`);
	if (opts.body && !headers.has('Content-Type')) headers.set('Content-Type', 'application/json');
	return api<T>(path, { ...opts, headers });
}
