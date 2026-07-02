<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { tanggal } from '$lib/format';
	import type { PendaftaranItem } from '$lib/types';

	let items = $state<PendaftaranItem[]>([]);
	let err = $state('');
	const statusOpts = ['aktif', 'selesai', 'dibatalkan'];

	async function load() {
		try {
			items = await apiAuth<PendaftaranItem[]>('/admin/pendaftaran');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	async function changeStatus(id: string, status: string) {
		try {
			await apiAuth(`/admin/pendaftaran/${id}/status`, {
				method: 'PATCH',
				body: JSON.stringify({ status })
			});
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Pendaftaran — Admin</title></svelte:head>

<h1>Kelola Pendaftaran</h1>
<p class="lead">Peserta yang terdaftar pada setiap kelas.</p>

{#if err}<div class="alert">{err}</div>{/if}

<table>
	<thead><tr><th>Peserta</th><th>Kelas</th><th>Tanggal</th><th>Status</th></tr></thead>
	<tbody>
		{#each items as p (p.pendaftaran_id)}
			<tr>
				<td>{p.user?.full_name || '—'}</td>
				<td>{p.kelas.judul}</td>
				<td>{tanggal(p.tanggal_daftar)}</td>
				<td>
					<select
						value={p.status}
						onchange={(e) => changeStatus(p.pendaftaran_id, e.currentTarget.value)}
					>
						{#each statusOpts as s}<option value={s}>{s}</option>{/each}
					</select>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="4" class="empty">Belum ada pendaftaran.</td></tr>
		{/if}
	</tbody>
</table>

<style>
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
