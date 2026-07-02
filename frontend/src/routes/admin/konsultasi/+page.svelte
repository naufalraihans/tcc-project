<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { tanggal } from '$lib/format';
	import type { Konsultasi } from '$lib/types';

	let items = $state<Konsultasi[]>([]);
	let err = $state('');
	let active = $state<Konsultasi | null>(null);
	let status = $state('menunggu');
	let balasan = $state('');

	const statusOpts = ['menunggu', 'diproses', 'selesai', 'ditolak'];

	async function load() {
		try {
			items = await apiAuth<Konsultasi[]>('/admin/konsultasi');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function open(k: Konsultasi) {
		active = k;
		status = k.status;
		balasan = k.balasan;
	}

	async function submit(e: SubmitEvent) {
		e.preventDefault();
		if (!active) return;
		err = '';
		try {
			await apiAuth(`/admin/konsultasi/${active.id}`, {
				method: 'PATCH',
				body: JSON.stringify({ status, balasan })
			});
			active = null;
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Konsultasi — Admin</title></svelte:head>

<h1>Kelola Konsultasi</h1>
<p class="lead">Tanggapi pengajuan konsultasi dari pengguna.</p>

{#if err}<div class="alert">{err}</div>{/if}

{#if active}
	<div class="panel" style="margin-bottom:24px">
		<div class="detail">
			<div><strong>{active.nama_pengirim}</strong> · {active.kontak}</div>
			<div class="topik">{active.topik_konsultasi}</div>
			<p class="pesan">{active.pesan}</p>
		</div>
		<form class="adm-form" onsubmit={submit}>
			<label>
				<span>Status</span>
				<select bind:value={status}>
					{#each statusOpts as s}<option value={s}>{s}</option>{/each}
				</select>
			</label>
			<label class="full"><span>Balasan</span><textarea bind:value={balasan} rows="3"></textarea></label>
			<div class="full actions">
				<button class="btn btn-primary" type="submit">Simpan Tanggapan</button>
				<button class="btn btn-ghost" type="button" onclick={() => (active = null)}>Batal</button>
			</div>
		</form>
	</div>
{/if}

<table>
	<thead><tr><th>Pengirim</th><th>Topik</th><th>Tanggal</th><th>Status</th><th></th></tr></thead>
	<tbody>
		{#each items as k (k.id)}
			<tr>
				<td>{k.nama_pengirim}</td>
				<td>{k.topik_konsultasi}</td>
				<td>{tanggal(k.created_at)}</td>
				<td><span class="pill pill-{k.status}">{k.status}</span></td>
				<td class="right"><button class="btn-sm" onclick={() => open(k)}>Tanggapi</button></td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="5" class="empty">Belum ada pengajuan.</td></tr>
		{/if}
	</tbody>
</table>

<style>
	.detail {
		margin-bottom: 18px;
		padding-bottom: 18px;
		border-bottom: 1px solid var(--border);
	}
	.topik {
		font-weight: 600;
		margin-top: 6px;
	}
	.pesan {
		margin-top: 8px;
		color: var(--cool-slate);
		white-space: pre-line;
	}
	.actions {
		display: flex;
		gap: 10px;
	}
	.right {
		text-align: right;
	}
	.empty {
		text-align: center;
		color: var(--muted);
	}
	.pill {
		font-size: 13px;
		font-weight: 600;
		padding: 4px 10px;
		border-radius: var(--radius-full);
		text-transform: capitalize;
		background: var(--off-white);
		color: var(--cool-slate);
	}
	.pill-selesai {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
	}
	.pill-ditolak {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
	}
</style>
