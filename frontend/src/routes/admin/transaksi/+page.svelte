<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { rupiah, tanggal } from '$lib/format';
	import type { Transaksi } from '$lib/types';

	let items = $state<Transaksi[]>([]);
	let err = $state('');
	const statusOpts = ['pending', 'sukses', 'gagal', 'refund'];

	async function load() {
		try {
			items = await apiAuth<Transaksi[]>('/admin/transaksi');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	async function changeStatus(id: string, status: string) {
		try {
			await apiAuth(`/admin/transaksi/${id}/status`, {
				method: 'PATCH',
				body: JSON.stringify({ status })
			});
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Transaksi — Admin</title></svelte:head>

<h1>Kelola Transaksi</h1>
<p class="lead">Riwayat dan status pembayaran.</p>

{#if err}<div class="alert">{err}</div>{/if}

<table>
	<thead><tr><th>Order ID</th><th>Jumlah</th><th>Metode</th><th>Tanggal</th><th>Status</th></tr></thead>
	<tbody>
		{#each items as t (t.id)}
			<tr>
				<td class="mono">{t.midtrans_order_id}</td>
				<td>{rupiah(t.jumlah)}</td>
				<td>{t.metode_pembayaran || '—'}</td>
				<td>{tanggal(t.created_at)}</td>
				<td>
					<select value={t.status} onchange={(e) => changeStatus(t.id, e.currentTarget.value)}>
						{#each statusOpts as s}<option value={s}>{s}</option>{/each}
					</select>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="5" class="empty">Belum ada transaksi.</td></tr>
		{/if}
	</tbody>
</table>

<style>
	.mono {
		font-family: monospace;
		font-size: 13px;
	}
	.empty {
		text-align: center;
		color: var(--muted);
	}
	td select {
		padding: 6px 10px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 13px;
		text-transform: capitalize;
	}
</style>
