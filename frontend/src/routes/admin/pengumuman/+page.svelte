<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import type { Pengumuman } from '$lib/types';

	let items = $state<Pengumuman[]>([]);
	let err = $state('');
	let editing = $state<string | null>(null);

	let judul = $state('');
	let isi = $state('');
	let tipe = $state('banner');
	let label_aksi = $state('');
	let url_aksi = $state('');
	let urutan = $state(0);
	let aktif = $state(true);
	let mulai = $state('');
	let selesai = $state('');

	// datetime-local <-> ISO (RFC3339). Kosong = tidak dijadwalkan.
	function toInput(iso?: string | null): string {
		if (!iso) return '';
		const d = new Date(iso);
		const p = (n: number) => String(n).padStart(2, '0');
		return `${d.getFullYear()}-${p(d.getMonth() + 1)}-${p(d.getDate())}T${p(d.getHours())}:${p(d.getMinutes())}`;
	}
	function toISO(v: string): string | null {
		return v ? new Date(v).toISOString() : null;
	}

	async function load() {
		try {
			items = await apiAuth<Pengumuman[]>('/admin/pengumuman');
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function reset() {
		editing = null;
		judul = '';
		isi = '';
		tipe = 'banner';
		label_aksi = '';
		url_aksi = '';
		urutan = 0;
		aktif = true;
		mulai = '';
		selesai = '';
	}

	function edit(p: Pengumuman) {
		editing = p.id;
		judul = p.judul;
		isi = p.isi;
		tipe = p.tipe;
		label_aksi = p.label_aksi;
		url_aksi = p.url_aksi;
		urutan = p.urutan;
		aktif = p.aktif ?? true;
		mulai = toInput(p.mulai);
		selesai = toInput(p.selesai);
	}

	async function save(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		const body = JSON.stringify({
			judul,
			isi,
			tipe,
			label_aksi,
			url_aksi,
			urutan,
			aktif,
			mulai: toISO(mulai),
			selesai: toISO(selesai)
		});
		try {
			if (editing) await apiAuth(`/admin/pengumuman/${editing}`, { method: 'PUT', body });
			else await apiAuth('/admin/pengumuman', { method: 'POST', body });
			reset();
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function remove(id: string) {
		if (!confirm('Hapus pengumuman ini?')) return;
		try {
			await apiAuth(`/admin/pengumuman/${id}`, { method: 'DELETE' });
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Pengumuman — Admin</title></svelte:head>

<h1>Kelola Pengumuman</h1>
<p class="lead">Banner atas & kartu info di dashboard peserta. Tipe "banner" tampil lebar di atas, "info" sebagai kartu di rail kanan.</p>

{#if err}<div class="alert">{err}</div>{/if}

<div class="panel" style="margin-bottom:24px">
	<form class="adm-form" onsubmit={save}>
		<label><span>Judul</span><input type="text" bind:value={judul} required /></label>
		<label>
			<span>Tipe</span>
			<select bind:value={tipe}>
				<option value="banner">banner</option>
				<option value="info">info</option>
			</select>
		</label>
		<label class="full"><span>Isi</span><textarea bind:value={isi} rows="2"></textarea></label>
		<label><span>Label Tombol (opsional)</span><input type="text" bind:value={label_aksi} placeholder="Lihat Katalog" /></label>
		<label><span>URL Tombol (opsional)</span><input type="text" bind:value={url_aksi} placeholder="/kelas" /></label>
		<label><span>Urutan</span><input type="number" bind:value={urutan} /></label>
		<label class="chk"><input type="checkbox" bind:checked={aktif} /><span>Aktif</span></label>
		<label><span>Mulai tayang (opsional)</span><input type="datetime-local" bind:value={mulai} /></label>
		<label><span>Selesai tayang (opsional)</span><input type="datetime-local" bind:value={selesai} /></label>
		<div class="full actions">
			<button class="btn btn-primary" type="submit">{editing ? 'Simpan Perubahan' : 'Tambah Pengumuman'}</button>
			{#if editing}<button class="btn btn-ghost" type="button" onclick={reset}>Batal</button>{/if}
		</div>
	</form>
</div>

<table>
	<thead><tr><th>Judul</th><th>Tipe</th><th>Urutan</th><th>Jadwal</th><th>Aktif</th><th></th></tr></thead>
	<tbody>
		{#each items as p (p.id)}
			<tr>
				<td>{p.judul}</td>
				<td>{p.tipe}</td>
				<td>{p.urutan}</td>
				<td>{p.mulai || p.selesai ? 'terjadwal' : '—'}</td>
				<td>{p.aktif ? 'ya' : 'tidak'}</td>
				<td class="right">
					<button class="btn-sm" onclick={() => edit(p)}>Ubah</button>
					<button class="btn-sm danger" onclick={() => remove(p.id)}>Hapus</button>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="6" class="empty">Belum ada pengumuman.</td></tr>
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
	.empty {
		text-align: center;
		color: var(--muted);
	}
</style>
