<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { api } from '$lib/api';
	import { apiAuth } from '$lib/authApi';
	import { harga, tanggal } from '$lib/format';
	import { ArrowLeft } from 'lucide-svelte';
	import type { Kelas, PendaftaranItem, DaftarResult } from '$lib/types';

	const slug = $page.params.slug;

	let k = $state<Kelas | null>(null);
	let enrolled = $state(false);
	let err = $state('');
	let loading = $state(true);
	let submitting = $state(false);
	let ok = $state('');

	const sisa = $derived(k && k.kuota > 0 ? k.kuota - k.peserta_terdaftar : null);

	onMount(async () => {
		try {
			const [detail, saya] = await Promise.all([
				api<Kelas>(`/kelas/${slug}`),
				apiAuth<PendaftaranItem[]>('/pendaftaran/saya')
			]);
			k = detail;
			enrolled = saya.some((p) => p.kelas.slug === slug);
		} catch (e) {
			err = (e as Error).message;
		} finally {
			loading = false;
		}
	});

	async function daftar() {
		if (!k) return;
		submitting = true;
		err = '';
		ok = '';
		try {
			const res = await apiAuth<DaftarResult>('/pendaftaran', {
				method: 'POST',
				body: JSON.stringify({ kelas_id: k.id })
			});
			if (res.type === 'berbayar' && res.redirect_url) {
				window.location.href = res.redirect_url; // ke Midtrans Snap
				return;
			}
			enrolled = true;
			ok = res.message || 'Berhasil mendaftar kelas.';
		} catch (e) {
			err = (e as Error).message;
		} finally {
			submitting = false;
		}
	}
</script>

<svelte:head><title>{k ? k.judul : 'Kelas'} — TCC ITPLN</title></svelte:head>

<a class="back" href="/dashboard/jelajah"><ArrowLeft size={16} /> Kembali ke katalog</a>

{#if loading}
	<div class="panel">Memuat…</div>
{:else if err && !k}
	<div class="panel">Gagal memuat kelas: {err}</div>
{:else if k}
	<div class="hero">
		<img src={`https://picsum.photos/seed/${slug}/1280/440`} alt="" />
	</div>
	<div class="detail">
		<div class="main">
			<div class="badges">
				<span class="badge">{k.format}</span>
				<span class="badge" class:free={k.tipe_harga === 'gratis'}>{harga(k.tipe_harga, k.harga)}</span>
				{#if k.topik}<span class="badge">{k.topik.nama}</span>{/if}
			</div>
			<h1>{k.judul}</h1>

			{#if k.deskripsi}
				<h2 class="sub">Deskripsi</h2>
				<p class="body">{k.deskripsi}</p>
			{/if}
			{#if k.silabus}
				<h2 class="sub">Silabus</h2>
				<p class="body">{k.silabus}</p>
			{/if}
			{#if k.instruktur}
				<h2 class="sub">Instruktur</h2>
				<p class="body">{k.instruktur.nama}</p>
			{/if}
		</div>

		<aside class="side">
			<div class="panel info">
				<dl>
					<dt>Jadwal Mulai</dt>
					<dd>{tanggal(k.jadwal_mulai)}</dd>
					<dt>Format</dt>
					<dd class="cap">{k.format}</dd>
					{#if k.lokasi}<dt>Lokasi</dt><dd>{k.lokasi}</dd>{/if}
					<dt>Kuota</dt>
					<dd>{sisa !== null ? `Sisa ${sisa} peserta` : 'Tidak dibatasi'}</dd>
					<dt>Biaya</dt>
					<dd class="price">{harga(k.tipe_harga, k.harga)}</dd>
				</dl>

				{#if ok}<div class="okbox">{ok}</div>{/if}
				{#if err}<div class="alert">{err}</div>{/if}

				{#if enrolled}
					<a class="btn btn-ghost block" href="/dashboard/kelas">Sudah terdaftar — Kelas Saya</a>
				{:else if k.status === 'aktif'}
					<button class="btn btn-primary block" onclick={daftar} disabled={submitting}>
						{submitting ? 'Memproses…' : k.tipe_harga === 'gratis' ? 'Daftar Kelas' : 'Daftar & Bayar'}
					</button>
				{:else}
					<button class="btn btn-ghost block" disabled>Pendaftaran ditutup</button>
				{/if}
			</div>
		</aside>
	</div>
{/if}

<style>
	.back {
		display: inline-flex;
		align-items: center;
		gap: 6px;
		color: var(--muted);
		font-size: 14px;
		margin-bottom: 18px;
	}
	.back:hover {
		color: var(--sky-blue);
	}
	.hero {
		height: 260px;
		border-radius: 18px;
		overflow: hidden;
		margin-bottom: 26px;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
	}
	.hero img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		display: block;
	}
	.detail {
		display: grid;
		grid-template-columns: 1.6fr 0.9fr;
		gap: 32px;
		align-items: start;
	}
	@media (max-width: 640px) {
		.hero {
			height: 180px;
		}
	}
	.badges {
		display: flex;
		gap: 8px;
		flex-wrap: wrap;
		margin-bottom: 16px;
	}
	.badge {
		display: inline-flex;
		align-items: center;
		padding: 4px 11px;
		border-radius: 999px;
		font-size: 12px;
		font-weight: 600;
		background: var(--off-white);
		color: var(--cool-slate);
		border: 1px solid var(--border);
		text-transform: capitalize;
	}
	.badge.free {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
		border-color: rgba(30, 123, 69, 0.2);
	}
	.main h1 {
		font-size: clamp(24px, 3.4vw, 32px);
		font-weight: 800;
		color: var(--ink);
	}
	.sub {
		font-size: 17px;
		font-weight: 700;
		margin: 24px 0 8px;
		color: var(--ink);
	}
	.body {
		color: var(--cool-slate);
		white-space: pre-line;
	}
	.side {
		position: sticky;
		top: 30px;
	}
	dl {
		display: grid;
		gap: 2px;
	}
	dt {
		font-size: 13px;
		color: var(--muted);
		margin-top: 12px;
	}
	dt:first-child {
		margin-top: 0;
	}
	dd {
		font-weight: 600;
		color: var(--ink);
	}
	dd.cap {
		text-transform: capitalize;
	}
	dd.price {
		font-family: var(--font-display);
		font-size: 20px;
		color: var(--navy-teal);
	}
	.block {
		display: flex;
		justify-content: center;
		width: 100%;
		margin-top: 20px;
	}
	.btn.block[disabled] {
		opacity: 0.6;
		cursor: not-allowed;
	}
	.okbox {
		background: rgba(30, 123, 69, 0.08);
		color: var(--success);
		border: 1px solid rgba(30, 123, 69, 0.2);
		padding: 10px 14px;
		border-radius: 10px;
		font-size: 13.5px;
		margin-top: 16px;
	}
	.alert {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
		border: 1px solid rgba(220, 38, 38, 0.2);
		padding: 10px 14px;
		border-radius: 10px;
		font-size: 13.5px;
		margin-top: 16px;
	}
	@media (max-width: 900px) {
		.detail {
			grid-template-columns: 1fr;
		}
		.side {
			position: static;
		}
	}
</style>
