<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { Topik } from '$lib/types';

	let items = $state<Topik[]>([]);
	let err = $state('');
	let editing = $state<string | null>(null);
	let nama = $state('');
	let slug = $state('');
	let deskripsi = $state('');

	async function load() {
		try {
			items = await apiAuth<Topik[]>('/topik');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function reset() {
		editing = null;
		nama = '';
		slug = '';
		deskripsi = '';
	}

	function edit(t: Topik) {
		editing = t.id;
		nama = t.nama;
		slug = t.slug;
		deskripsi = t.deskripsi;
	}

	async function save(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		const body = JSON.stringify({ nama, slug, deskripsi });
		try {
			if (editing) await apiAuth(`/admin/topik/${editing}`, { method: 'PUT', body });
			else await apiAuth('/admin/topik', { method: 'POST', body });
			reset();
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function remove(id: string) {
		if (!confirm('Hapus topik ini?')) return;
		try {
			await apiAuth(`/admin/topik/${id}`, { method: 'DELETE' });
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Topik — Admin</title></svelte:head>

<h1>Kelola Topik</h1>
<p class="lead">Tambah, ubah, atau hapus bidang pelatihan.</p>

{#if err}<div class="alert">{err}</div>{/if}

<div class="panel" style="margin-bottom:24px">
	<form class="adm-form" onsubmit={save}>
		<label><span>Nama</span><input type="text" bind:value={nama} required /></label>
		<label><span>Slug (opsional)</span><input type="text" bind:value={slug} /></label>
		<label class="full"><span>Deskripsi</span><textarea bind:value={deskripsi} rows="2"></textarea></label>
		<div class="full actions">
			<button class="btn btn-primary" type="submit">{editing ? 'Simpan Perubahan' : 'Tambah Topik'}</button>
			{#if editing}<button class="btn btn-ghost" type="button" onclick={reset}>Batal</button>{/if}
		</div>
	</form>
</div>

<table>
	<thead><tr><th>Nama</th><th>Slug</th><th>Kelas</th><th></th></tr></thead>
	<tbody>
		{#each items as t (t.id)}
			<tr>
				<td>{t.nama}</td>
				<td>{t.slug}</td>
				<td>{t.jumlah_kelas ?? 0}</td>
				<td class="right">
					<button class="btn-sm" onclick={() => edit(t)}>Ubah</button>
					<button class="btn-sm danger" onclick={() => remove(t.id)}>Hapus</button>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="4" class="empty">Belum ada topik.</td></tr>
		{/if}
	</tbody>
</table>

<style>
	.actions {
		display: flex;
		gap: 10px;
	}
	.right {
		text-align: right;
		white-space: nowrap;
	}
	.right .btn-sm {
		margin-left: 6px;
	}
	.empty {
		text-align: center;
		color: var(--muted);
	}
</style>
