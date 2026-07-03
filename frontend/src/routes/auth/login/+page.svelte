<script lang="ts">
	import { login } from '$lib/authApi';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	let email = $state('');
	let password = $state('');
	let loading = $state(false);
	let err = $state('');

	async function submit(e: SubmitEvent) {
		e.preventDefault();
		loading = true;
		err = '';
		try {
			const res = await login(email, password);
			const next = $page.url.searchParams.get('next');
			goto(res.role === 'admin' ? '/admin' : next || '/dashboard');
		} catch (e) {
			err = (e as Error).message;
			loading = false;
		}
	}
</script>

<svelte:head><title>Masuk — TCC ITPLN</title></svelte:head>

<form class="card" onsubmit={submit}>
	<h1>Selamat Datang</h1>
	<p class="sub">Masuk ke akun TCC ITPLN Anda.</p>
	{#if err}<div class="alert">{err}</div>{/if}
	<label><span>Email</span><input type="email" bind:value={email} required /></label>
	<label><span>Password</span><input type="password" bind:value={password} required /></label>
	<div class="row"><a href="/auth/lupa-password">Lupa password?</a></div>
	<button class="btn btn-primary block" disabled={loading}>{loading ? 'Memproses…' : 'Masuk'}</button>
	<p class="foot">Belum punya akun? <a href="/auth/register">Daftar</a></p>
</form>
