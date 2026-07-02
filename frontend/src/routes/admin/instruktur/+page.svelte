<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { Instruktur } from '$lib/types';

	let items = $state<Instruktur[]>([]);
	let err = $state('');
	let editing = $state<string | null>(null);
	let nama = $state('');
	let jabatan = $state('');
	let bio = $state('');

	async function load() {
		try {
			items = await apiAuth<Instruktur[]>('/instruktur');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function reset() {
		editing = null;
		nama = '';
		jabatan = '';
		bio = '';
	}

	function edit(i: Instruktur) {
		editing = i.id;
		nama = i.nama;
		jabatan = i.jabatan;
		bio = i.bio;
	}

	async function save(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		const body = JSON.stringify({ nama, jabatan, bio });
		try {
			if (editing) await apiAuth(`/admin/instruktur/${editing}`, { method: 'PUT', body });
			else await apiAuth('/admin/instruktur', { method: 'POST', body });
			reset();
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function remove(id: string) {
		if (!confirm('Hapus instruktur ini?')) return;
		try {
			await apiAuth(`/admin/instruktur/${id}`, { method: 'DELETE' });
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Instruktur — Admin</title></svelte:head>

<h1>Kelola Instruktur</h1>
<p class="lead">Data pengajar dan konsultan TCC.</p>

{#if err}<div class="alert">{err}</div>{/if}

<div class="panel" style="margin-bottom:24px">
	<form class="adm-form" onsubmit={save}>
		<label><span>Nama</span><input type="text" bind:value={nama} required /></label>
		<label><span>Jabatan</span><input type="text" bind:value={jabatan} /></label>
		<label class="full"><span>Bio</span><textarea bind:value={bio} rows="2"></textarea></label>
		<div class="full actions">
			<button class="btn btn-primary" type="submit">{editing ? 'Simpan Perubahan' : 'Tambah Instruktur'}</button>
			{#if editing}<button class="btn btn-ghost" type="button" onclick={reset}>Batal</button>{/if}
		</div>
	</form>
</div>

<table>
	<thead><tr><th>Nama</th><th>Jabatan</th><th></th></tr></thead>
	<tbody>
		{#each items as i (i.id)}
			<tr>
				<td>{i.nama}</td>
				<td>{i.jabatan}</td>
				<td class="right">
					<button class="btn-sm" onclick={() => edit(i)}>Ubah</button>
					<button class="btn-sm danger" onclick={() => remove(i.id)}>Hapus</button>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="3" class="empty">Belum ada instruktur.</td></tr>
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
