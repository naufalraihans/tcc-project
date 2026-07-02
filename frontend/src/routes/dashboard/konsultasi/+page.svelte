<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { tanggal } from '$lib/format';
	import type { Konsultasi } from '$lib/types';

	let items = $state<Konsultasi[]>([]);
	let loading = $state(true);
	let err = $state('');

	onMount(async () => {
		try {
			items = await apiAuth<Konsultasi[]>('/konsultasi/saya');
		} catch (e) {
			err = (e as Error).message;
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head><title>Konsultasi — TCC ITPLN</title></svelte:head>

<h1>Riwayat Konsultasi</h1>
<p class="lead">Status dan balasan dari pengajuan konsultasi Anda.</p>

{#if loading}
	<div class="panel">Memuat…</div>
{:else if err}
	<div class="panel">Gagal memuat: {err}</div>
{:else if items.length === 0}
	<div class="panel">
		Belum ada pengajuan konsultasi. <a href="/konsultasi">Ajukan konsultasi</a>.
	</div>
{:else}
	<div class="list">
		{#each items as k (k.id)}
			<div class="panel">
				<div class="top">
					<div class="judul">{k.topik_konsultasi}</div>
					<span class="status status-{k.status}">{k.status}</span>
				</div>
				<div class="tgl">{tanggal(k.created_at)}</div>
				<p class="pesan">{k.pesan}</p>
				{#if k.balasan}
					<div class="balasan">
						<div class="balasan-label">Balasan TCC</div>
						<p>{k.balasan}</p>
					</div>
				{/if}
			</div>
		{/each}
	</div>
{/if}

<style>
	.list {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}
	.top {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 12px;
	}
	.judul {
		font-weight: 600;
		font-size: 16px;
	}
	.tgl {
		color: var(--muted);
		font-size: 13px;
		margin-top: 2px;
	}
	.pesan {
		margin-top: 12px;
		color: var(--cool-slate);
		white-space: pre-line;
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
	.status-selesai {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
	}
	.status-ditolak {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
	}
	.balasan {
		margin-top: 16px;
		padding: 14px 16px;
		background: var(--off-white);
		border-radius: var(--radius-sm);
	}
	.balasan-label {
		font-size: 13px;
		font-weight: 600;
		color: var(--navy-teal);
		margin-bottom: 6px;
	}
</style>
