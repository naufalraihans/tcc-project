<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { PendaftaranItem, Konsultasi, Profile } from '$lib/types';

	let profile = $state<Profile | null>(null);
	let kelas = $state<PendaftaranItem[]>([]);
	let konsultasi = $state<Konsultasi[]>([]);
	let err = $state('');

	const aktif = $derived(kelas.filter((k) => k.status === 'aktif').length);

	onMount(async () => {
		try {
			[profile, kelas, konsultasi] = await Promise.all([
				apiAuth<Profile>('/auth/me'),
				apiAuth<PendaftaranItem[]>('/pendaftaran/saya'),
				apiAuth<Konsultasi[]>('/konsultasi/saya')
			]);
		} catch (e) {
			err = (e as Error).message;
		}
	});
</script>

<svelte:head><title>Dashboard — TCC ITPLN</title></svelte:head>

<h1>Halo{profile ? `, ${profile.full_name}` : ''}</h1>
<p class="lead">Ringkasan aktivitas pelatihan Anda.</p>

{#if err}
	<div class="panel">Gagal memuat data: {err}</div>
{:else}
	<div class="stats">
		<div class="panel stat">
			<div class="num">{kelas.length}</div>
			<div class="cap">Total Kelas</div>
		</div>
		<div class="panel stat">
			<div class="num">{aktif}</div>
			<div class="cap">Kelas Aktif</div>
		</div>
		<div class="panel stat">
			<div class="num">{konsultasi.length}</div>
			<div class="cap">Konsultasi</div>
		</div>
	</div>

	<div class="links">
		<a class="btn btn-ghost" href="/dashboard/kelas">Lihat Kelas Saya</a>
		<a class="btn btn-ghost" href="/kelas">Jelajahi Kelas</a>
	</div>
{/if}

<style>
	.stats {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 20px;
	}
	.stat .num {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 36px;
		color: var(--navy-teal);
	}
	.stat .cap {
		color: var(--muted);
		font-size: 14px;
		margin-top: 4px;
	}
	.links {
		display: flex;
		gap: 12px;
		margin-top: 24px;
		flex-wrap: wrap;
	}
	@media (max-width: 640px) {
		.stats {
			grid-template-columns: 1fr;
		}
	}
</style>
