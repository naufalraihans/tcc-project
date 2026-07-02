<script lang="ts">
	import { supabase } from '$lib/supabase';
	import { goto } from '$app/navigation';

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
		const { error } = await supabase.auth.updateUser({ password });
		loading = false;
		if (error) {
			err = 'Tautan reset tidak berlaku atau sudah kedaluwarsa. Silakan minta ulang.';
			return;
		}
		ok = 'Password berhasil diperbarui. Mengalihkan ke halaman masuk…';
		setTimeout(() => goto('/auth/login'), 1500);
	}
</script>

<svelte:head><title>Reset Password — TCC ITPLN</title></svelte:head>

<form class="card" onsubmit={submit}>
	<h1>Atur Password Baru</h1>
	<p class="sub">Masukkan password baru untuk akun Anda.</p>
	{#if err}<div class="alert">{err}</div>{/if}
	{#if ok}<div class="ok">{ok}</div>{/if}
	<label><span>Password Baru</span><input type="password" bind:value={password} required /></label>
	<label><span>Konfirmasi Password</span><input type="password" bind:value={confirm} required /></label>
	<button class="btn btn-primary block" disabled={loading}>
		{loading ? 'Memproses…' : 'Simpan Password'}
	</button>
</form>
