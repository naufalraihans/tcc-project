<script lang="ts">
	import { supabase } from '$lib/supabase';

	let email = $state('');
	let loading = $state(false);
	let ok = $state('');
	let err = $state('');

	async function submit(e: SubmitEvent) {
		e.preventDefault();
		loading = true;
		err = '';
		const { error } = await supabase.auth.resetPasswordForEmail(email, {
			redirectTo: `${location.origin}/auth/reset-password`
		});
		loading = false;
		if (error) {
			err = error.message;
			return;
		}
		ok = 'Jika email terdaftar, tautan reset telah dikirim. Silakan cek kotak masuk Anda.';
	}
</script>

<svelte:head><title>Lupa Password — TCC ITPLN</title></svelte:head>

<form class="card" onsubmit={submit}>
	<h1>Lupa Password</h1>
	<p class="sub">Masukkan email Anda untuk menerima tautan reset password.</p>
	{#if err}<div class="alert">{err}</div>{/if}
	{#if ok}<div class="ok">{ok}</div>{/if}
	<label><span>Email</span><input type="email" bind:value={email} required /></label>
	<button class="btn btn-primary block" disabled={loading}>
		{loading ? 'Memproses…' : 'Kirim Tautan Reset'}
	</button>
	<p class="foot"><a href="/auth/login">Kembali ke Masuk</a></p>
</form>
