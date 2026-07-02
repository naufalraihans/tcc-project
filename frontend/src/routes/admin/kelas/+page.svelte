<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { harga as fmtHarga } from '$lib/format';
	import type { KelasListItem, Kelas, Topik, Instruktur } from '$lib/types';

	let items = $state<KelasListItem[]>([]);
	let topikList = $state<Topik[]>([]);
	let instrukturList = $state<Instruktur[]>([]);
	let err = $state('');
	let showForm = $state(false);
	let editing = $state<string | null>(null);

	let judul = $state('');
	let slug = $state('');
	let deskripsi = $state('');
	let silabus = $state('');
	let topikId = $state('');
	let instrukturId = $state('');
	let format = $state('online');
	let tipeHarga = $state('gratis');
	let hargaVal = $state(0);
	let jadwalMulai = $state('');
	let jadwalSelesai = $state('');
	let durasiMenit = $state(0);
	let kuota = $state(0);
	let lokasi = $state('');
	let linkMeeting = $state('');

	const statusOpts = ['aktif', 'penuh', 'selesai'];

	async function load() {
		try {
			const [k, t, i] = await Promise.all([
				apiAuth<{ items: KelasListItem[] }>('/kelas?status=semua&limit=100'),
				apiAuth<Topik[]>('/topik'),
				apiAuth<Instruktur[]>('/instruktur')
			]);
			items = k.items;
			topikList = t;
			instrukturList = i;
		} catch (e) {
			err = (e as Error).message;
		}
	}
	onMount(load);

	function reset() {
		editing = null;
		showForm = false;
		judul = slug = deskripsi = silabus = topikId = instrukturId = '';
		format = 'online';
		tipeHarga = 'gratis';
		hargaVal = 0;
		jadwalMulai = jadwalSelesai = lokasi = linkMeeting = '';
		durasiMenit = 0;
		kuota = 0;
	}

	function openNew() {
		reset();
		showForm = true;
	}

	async function openEdit(item: KelasListItem) {
		err = '';
		try {
			const k = await apiAuth<Kelas>(`/kelas/${item.slug}`);
			editing = k.id;
			judul = k.judul;
			slug = k.slug;
			deskripsi = k.deskripsi;
			silabus = k.silabus;
			topikId = k.topik?.id ?? '';
			instrukturId = k.instruktur?.id ?? '';
			format = k.format;
			tipeHarga = k.tipe_harga;
			hargaVal = k.harga;
			jadwalMulai = k.jadwal_mulai ? k.jadwal_mulai.slice(0, 16) : '';
			jadwalSelesai = k.jadwal_selesai ? k.jadwal_selesai.slice(0, 16) : '';
			durasiMenit = k.durasi_menit;
			kuota = k.kuota;
			lokasi = k.lokasi;
			linkMeeting = k.link_meeting;
			showForm = true;
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function save(e: SubmitEvent) {
		e.preventDefault();
		err = '';
		const body = JSON.stringify({
			judul,
			slug,
			deskripsi,
			silabus,
			topik_id: topikId || null,
			instruktur_id: instrukturId || null,
			format,
			tipe_harga: tipeHarga,
			harga: tipeHarga === 'berbayar' ? Number(hargaVal) || 0 : 0,
			jadwal_mulai: jadwalMulai ? new Date(jadwalMulai).toISOString() : null,
			jadwal_selesai: jadwalSelesai ? new Date(jadwalSelesai).toISOString() : null,
			durasi_menit: Number(durasiMenit) || 0,
			kuota: Number(kuota) || 0,
			lokasi,
			link_meeting: linkMeeting
		});
		try {
			if (editing) await apiAuth(`/admin/kelas/${editing}`, { method: 'PUT', body });
			else await apiAuth('/admin/kelas', { method: 'POST', body });
			reset();
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function changeStatus(id: string, status: string) {
		try {
			await apiAuth(`/admin/kelas/${id}/status`, {
				method: 'PATCH',
				body: JSON.stringify({ status })
			});
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}

	async function remove(id: string) {
		if (!confirm('Hapus kelas ini?')) return;
		try {
			await apiAuth(`/admin/kelas/${id}`, { method: 'DELETE' });
			await load();
		} catch (e) {
			err = (e as Error).message;
		}
	}
</script>

<svelte:head><title>Kelola Kelas — Admin</title></svelte:head>

<div class="toolbar">
	<div>
		<h1>Kelola Kelas</h1>
		<p class="lead" style="margin:0">Buat dan kelola program pelatihan.</p>
	</div>
	{#if !showForm}
		<button class="btn btn-primary" onclick={openNew}>Tambah Kelas</button>
	{/if}
</div>

{#if err}<div class="alert">{err}</div>{/if}

{#if showForm}
	<div class="panel" style="margin-bottom:24px">
		<form class="adm-form" onsubmit={save}>
			<label class="full"><span>Judul</span><input type="text" bind:value={judul} required /></label>
			<label><span>Slug (opsional)</span><input type="text" bind:value={slug} /></label>
			<label>
				<span>Topik</span>
				<select bind:value={topikId}>
					<option value="">—</option>
					{#each topikList as t}<option value={t.id}>{t.nama}</option>{/each}
				</select>
			</label>
			<label>
				<span>Instruktur</span>
				<select bind:value={instrukturId}>
					<option value="">—</option>
					{#each instrukturList as i}<option value={i.id}>{i.nama}</option>{/each}
				</select>
			</label>
			<label>
				<span>Format</span>
				<select bind:value={format}>
					<option value="online">Online</option>
					<option value="offline">Offline</option>
					<option value="hybrid">Hybrid</option>
				</select>
			</label>
			<label>
				<span>Tipe Harga</span>
				<select bind:value={tipeHarga}>
					<option value="gratis">Gratis</option>
					<option value="berbayar">Berbayar</option>
				</select>
			</label>
			{#if tipeHarga === 'berbayar'}
				<label><span>Harga (Rp)</span><input type="number" bind:value={hargaVal} min="0" /></label>
			{/if}
			<label><span>Jadwal Mulai</span><input type="datetime-local" bind:value={jadwalMulai} /></label>
			<label><span>Jadwal Selesai</span><input type="datetime-local" bind:value={jadwalSelesai} /></label>
			<label><span>Durasi (menit)</span><input type="number" bind:value={durasiMenit} min="0" /></label>
			<label><span>Kuota (0 = tanpa batas)</span><input type="number" bind:value={kuota} min="0" /></label>
			<label class="full"><span>Lokasi</span><input type="text" bind:value={lokasi} /></label>
			<label class="full"><span>Link Meeting</span><input type="text" bind:value={linkMeeting} /></label>
			<label class="full"><span>Deskripsi</span><textarea bind:value={deskripsi} rows="3"></textarea></label>
			<label class="full"><span>Silabus</span><textarea bind:value={silabus} rows="3"></textarea></label>
			<div class="full actions">
				<button class="btn btn-primary" type="submit">{editing ? 'Simpan Perubahan' : 'Buat Kelas'}</button>
				<button class="btn btn-ghost" type="button" onclick={reset}>Batal</button>
			</div>
		</form>
	</div>
{/if}

<table>
	<thead><tr><th>Judul</th><th>Format</th><th>Harga</th><th>Status</th><th></th></tr></thead>
	<tbody>
		{#each items as k (k.id)}
			<tr>
				<td>{k.judul}</td>
				<td>{k.format}</td>
				<td>{fmtHarga(k.tipe_harga, k.harga)}</td>
				<td>
					<select value={k.status} onchange={(e) => changeStatus(k.id, e.currentTarget.value)}>
						{#each statusOpts as s}<option value={s}>{s}</option>{/each}
					</select>
				</td>
				<td class="right">
					<button class="btn-sm" onclick={() => openEdit(k)}>Ubah</button>
					<button class="btn-sm danger" onclick={() => remove(k.id)}>Hapus</button>
				</td>
			</tr>
		{/each}
		{#if items.length === 0}
			<tr><td colspan="5" class="empty">Belum ada kelas.</td></tr>
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
	td select {
		padding: 6px 10px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 13px;
		text-transform: capitalize;
	}
</style>
