<script lang="ts">
	import { supabase } from '$lib/supabase';
	import { goto } from '$app/navigation';

	let fullName = $state('');
	let email = $state('');
	let password = $state('');
	let confirm = $state('');
	let loading = $state(false);
	let err = $state('');
	let ok = $state('');

	async function submit(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		if (password.length < 6) {
			err = 'Password minimal 6 karakter.';
			return;
		}
		if (password !== confirm) {
			err = 'Konfirmasi password tidak cocok.';
			return;
		}
		loading = true;
		const { data, error } = await supabase.auth.signUp({
			email,
			password,
			options: {
				data: { full_name: fullName },
				emailRedirectTo: `${location.origin}/auth/login`
			}
		});
		loading = false;
		if (error) {
			err = error.message;
			return;
		}
		if (data.session) {
			goto('/dashboard');
		} else {
			ok = 'Akun dibuat. Silakan cek email untuk konfirmasi sebelum masuk.';
		}
	}
</script>

<svelte:head><title>Daftar — TCC ITPLN</title></svelte:head>

<form class="card" onsubmit={submit}>
	<h1>Buat Akun Baru</h1>
	<p class="sub">Daftar untuk mengikuti pelatihan TCC ITPLN.</p>
	{#if err}<div class="alert">{err}</div>{/if}
	{#if ok}<div class="ok">{ok}</div>{/if}
	<label><span>Nama Lengkap</span><input type="text" bind:value={fullName} required /></label>
	<label><span>Email</span><input type="email" bind:value={email} required /></label>
	<label><span>Password</span><input type="password" bind:value={password} required /></label>
	<label><span>Konfirmasi Password</span><input type="password" bind:value={confirm} required /></label>
	<button class="btn btn-primary block" disabled={loading}>{loading ? 'Memproses…' : 'Daftar'}</button>
	<p class="foot">Sudah punya akun? <a href="/auth/login">Masuk</a></p>
</form>
