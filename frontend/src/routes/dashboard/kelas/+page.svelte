<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { tanggal } from '$lib/format';
	import type { PendaftaranItem } from '$lib/types';

	let items = $state<PendaftaranItem[]>([]);
	let loading = $state(true);
	let err = $state('');

	onMount(async () => {
		try {
			items = await apiAuth<PendaftaranItem[]>('/pendaftaran/saya');
		} catch (e) {
			err = (e as Error).message;
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head><title>Kelas Saya — TCC ITPLN</title></svelte:head>

<h1>Kelas Saya</h1>
<p class="lead">Kelas yang sedang dan telah Anda ikuti.</p>

{#if loading}
	<div class="panel">Memuat…</div>
{:else if err}
	<div class="panel">Gagal memuat: {err}</div>
{:else if items.length === 0}
	<div class="panel">
		Anda belum mendaftar kelas apa pun. <a href="/dashboard/jelajah">Jelajahi kelas</a>.
	</div>
{:else}
	<div class="list">
		{#each items as it (it.pendaftaran_id)}
			<a class="panel row" href="/dashboard/jelajah/{it.kelas.slug}">
				<div>
					<div class="judul">{it.kelas.judul}</div>
					<div class="meta">{it.kelas.format} · Terdaftar {tanggal(it.tanggal_daftar)}</div>
				</div>
				<span class="status status-{it.status}">{it.status}</span>
			</a>
		{/each}
	</div>
{/if}

<style>
	.list {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	.row {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 16px;
	}
	.judul {
		font-weight: 600;
		font-size: 16px;
	}
	.meta {
		color: var(--muted);
		font-size: 13px;
		margin-top: 4px;
	}
	.status {
		font-size: 13px;
		font-weight: 600;
		padding: 5px 12px;
		border-radius: var(--radius-full);
		text-transform: capitalize;
		background: var(--off-white);
		color: var(--cool-slate);
	}
	.status-aktif {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
	}
	.status-dibatalkan {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
	}
</style>
