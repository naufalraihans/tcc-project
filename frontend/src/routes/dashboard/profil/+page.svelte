<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { Profile } from '$lib/types';

	let profile = $state<Profile | null>(null);
	let fullName = $state('');
	let phone = $state('');
	let loading = $state(true);
	let saving = $state(false);
	let err = $state('');
	let ok = $state('');

	onMount(async () => {
		try {
			profile = await apiAuth<Profile>('/auth/me');
			fullName = profile.full_name;
			phone = profile.phone;
		} catch (e) {
			err = (e as Error).message;
		} finally {
			loading = false;
		}
	});

	async function save(e: SubmitEvent) {
		e.preventDefault();
		saving = true;
		err = '';
		ok = '';
		try {
			profile = await apiAuth<Profile>('/auth/profile', {
				method: 'PUT',
				body: JSON.stringify({ full_name: fullName, phone, avatar_url: profile?.avatar_url ?? '' })
			});
			ok = 'Profil berhasil diperbarui.';
		} catch (e) {
			err = (e as Error).message;
		} finally {
			saving = false;
		}
	}
</script>

<svelte:head><title>Profil — TCC ITPLN</title></svelte:head>

<h1>Profil</h1>
<p class="lead">Kelola data diri Anda.</p>

{#if loading}
	<div class="panel">Memuat…</div>
{:else}
	<form class="panel form" onsubmit={save}>
		{#if err}<div class="alert">{err}</div>{/if}
		{#if ok}<div class="okbox">{ok}</div>{/if}
		<label>
			<span>Email</span>
			<input type="email" value={profile?.email ?? ''} disabled />
		</label>
		<label>
			<span>Nama Lengkap</span>
			<input type="text" bind:value={fullName} required />
		</label>
		<label>
			<span>Nomor Telepon</span>
			<input type="text" bind:value={phone} />
		</label>
		<button class="btn btn-primary" disabled={saving}>{saving ? 'Menyimpan…' : 'Simpan'}</button>
	</form>
{/if}

<style>
	.form {
		max-width: 480px;
		display: flex;
		flex-direction: column;
		gap: 16px;
	}
	label {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}
	label span {
		font-size: 14px;
		font-weight: 500;
		color: var(--cool-slate);
	}
	input {
		padding: 11px 14px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 15px;
	}
	input:disabled {
		background: var(--off-white);
		color: var(--muted);
	}
	input:focus {
		outline: none;
		border-color: var(--sky-blue);
	}
	.alert {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
		border: 1px solid rgba(220, 38, 38, 0.2);
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
	}
	.okbox {
		background: rgba(30, 123, 69, 0.08);
		color: var(--success);
		border: 1px solid rgba(30, 123, 69, 0.2);
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
	}
	button {
		align-self: flex-start;
	}
</style>
