<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { rupiah } from '$lib/format';
	import type { PagedKelas, PendaftaranItem, Konsultasi, Transaksi } from '$lib/types';

	let err = $state('');
	let totalKelas = $state(0);
	let totalPendaftaran = $state(0);
	let konsultasiMenunggu = $state(0);
	let pendapatan = $state(0);

	onMount(async () => {
		try {
			const [kelas, pend, kons, trx] = await Promise.all([
				apiAuth<PagedKelas>('/kelas?status=semua&limit=100'),
				apiAuth<PendaftaranItem[]>('/admin/pendaftaran'),
				apiAuth<Konsultasi[]>('/admin/konsultasi'),
				apiAuth<Transaksi[]>('/admin/transaksi')
			]);
			totalKelas = kelas.pagination.total;
			totalPendaftaran = pend.length;
			konsultasiMenunggu = kons.filter((k) => k.status === 'menunggu').length;
			pendapatan = trx
				.filter((t) => t.status === 'sukses')
				.reduce((sum, t) => sum + t.jumlah, 0);
		} catch (e) {
			err = (e as Error).message;
		}
	});

	const cards = $derived([
		{ num: String(totalKelas), cap: 'Total Kelas' },
		{ num: String(totalPendaftaran), cap: 'Total Pendaftaran' },
		{ num: String(konsultasiMenunggu), cap: 'Konsultasi Menunggu' },
		{ num: pendapatan > 0 ? rupiah(pendapatan) : 'Rp 0', cap: 'Pendapatan' }
	]);
</script>

<svelte:head><title>Dashboard Admin — TCC ITPLN</title></svelte:head>

<h1>Dashboard Admin</h1>
<p class="lead">Ringkasan platform TCC ITPLN.</p>

{#if err}
	<div class="panel">Gagal memuat statistik: {err}</div>
{:else}
	<div class="stats">
		{#each cards as c}
			<div class="panel stat">
				<div class="num">{c.num}</div>
				<div class="cap">{c.cap}</div>
			</div>
		{/each}
	</div>
{/if}

<style>
	.stats {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 18px;
	}
	.stat .num {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 28px;
		color: var(--navy-teal);
	}
	.stat .cap {
		color: var(--muted);
		font-size: 14px;
		margin-top: 4px;
	}
	@media (max-width: 900px) {
		.stats {
			grid-template-columns: repeat(2, 1fr);
		}
	}
</style>
