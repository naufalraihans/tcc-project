<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { MisiDef } from '$lib/types';

	let items = $state<MisiDef[]>([]);
	let err = $state('');
	let editing = $state<string | null>(null);

	let kode = $state('');
	let judul = $state('');
	let deskripsi = $state('');
	let tipe = $state('harian');
	let target = $state(1);
	let xp_reward = $state(10);
	let aktif = $state(true);

	async function load() {
		try {
			items = await apiAuth<MisiDef[]>('/admin/misi');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function reset() {
		editing = null;
		kode = '';
		judul = '';
		deskripsi = '';
		tipe = 'harian';
		target = 1;
		xp_reward = 10;
		aktif = true;
	}

	function edit(m: MisiDef) {
		editing = m.id;
		kode = m.kode;
		judul = m.judul;
		deskripsi = m.deskripsi;
		tipe = m.tipe;
		target = m.target;
		xp_reward = m.xp_reward;
		aktif = m.aktif;
	}

	async function save(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		const body = JSON.stringify({ kode, judul, deskripsi, tipe, target, xp_reward, aktif });
		try {
			if (editing) await apiAuth(`/admin/misi/${editing}`, { method: 'PUT', body });
			else await apiAuth('/admin/misi', { method: 'POST', body });
			reset();
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function remove(id: string) {
		if (!confirm('Hapus misi ini?')) return;
		try {
			await apiAuth(`/admin/misi/${id}`, { method: 'DELETE' });
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Misi — Admin</title></svelte:head>

<h1>Kelola Misi</h1>
<p class="lead">Definisi misi harian/mingguan yang tampil di dashboard peserta. Kode memicu hook XP di backend — ubah dengan hati-hati.</p>

{#if err}<div class="alert">{err}</div>{/if}

<div class="panel" style="margin-bottom:24px">
	<form class="adm-form" onsubmit={save}>
		<label><span>Kode</span><input type="text" bind:value={kode} placeholder="buka_materi" required /></label>
		<label><span>Judul</span><input type="text" bind:value={judul} required /></label>
		<label class="full"><span>Deskripsi</span><input type="text" bind:value={deskripsi} /></label>
		<label>
			<span>Tipe</span>
			<select bind:value={tipe}>
				<option value="harian">harian</option>
				<option value="mingguan">mingguan</option>
				<option value="sekali">sekali</option>
			</select>
		</label>
		<label><span>Target</span><input type="number" min="1" bind:value={target} required /></label>
		<label><span>XP Reward</span><input type="number" min="0" bind:value={xp_reward} /></label>
		<label class="chk"><input type="checkbox" bind:checked={aktif} /><span>Aktif</span></label>
		<div class="full actions">
			<button class="btn btn-primary" type="submit">{editing ? 'Simpan Perubahan' : 'Tambah Misi'}</button>
			{#if editing}<button class="btn btn-ghost" type="button" onclick={reset}>Batal</button>{/if}
		</div>
	</form>
</div>

<table>
	<thead><tr><th>Kode</th><th>Judul</th><th>Tipe</th><th>Target</th><th>XP</th><th>Aktif</th><th></th></tr></thead>
	<tbody>
		{#each items as m (m.id)}
			<tr>
				<td><code>{m.kode}</code></td>
				<td>{m.judul}</td>
				<td>{m.tipe}</td>
				<td>{m.target}</td>
				<td>{m.xp_reward}</td>
				<td>{m.aktif ? 'ya' : 'tidak'}</td>
				<td class="right">
					<button class="btn-sm" onclick={() => edit(m)}>Ubah</button>
					<button class="btn-sm danger" onclick={() => remove(m.id)}>Hapus</button>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="7" class="empty">Belum ada misi.</td></tr>
		{/if}
	</tbody>
</table>

<style>
	.actions {
		display: flex;
		gap: 10px;
	}
	.chk {
		flex-direction: row !important;
		align-items: center;
		gap: 8px !important;
	}
	.chk input {
		width: 16px;
		height: 16px;
	}
	.right {
		text-align: right;
		white-space: nowrap;
	}
	.right .btn-sm {
		margin-left: 6px;
	}
	code {
		font-family: var(--font-mono, monospace);
		font-size: 13px;
		color: var(--navy-teal);
	}
	.empty {
		text-align: center;
		color: var(--muted);
	}
</style>
