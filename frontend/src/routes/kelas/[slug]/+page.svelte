<script lang="ts">
	import { harga, tanggal } from '$lib/format';

	let { data } = $props();
	const k = $derived(data.kelas);
	const sisa = $derived(k.kuota > 0 ? k.kuota - k.peserta_terdaftar : null);
</script>

<svelte:head><title>{k.judul} — TCC ITPLN</title></svelte:head>

<section class="section">
	<div class="container detail">
		<div class="main">
			<a class="back" href="/kelas">← Kembali ke daftar kelas</a>
			<div class="badges">
				<span class="badge">{k.format}</span>
				<span class="badge" class:free={k.tipe_harga === 'gratis'}>
					{harga(k.tipe_harga, k.harga)}
				</span>
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
			<div class="card info">
				<dl>
					<dt>Jadwal Mulai</dt>
					<dd>{tanggal(k.jadwal_mulai)}</dd>
					<dt>Format</dt>
					<dd>{k.format}</dd>
					{#if k.lokasi}
						<dt>Lokasi</dt>
						<dd>{k.lokasi}</dd>
					{/if}
					<dt>Kuota</dt>
					<dd>{sisa !== null ? `Sisa ${sisa} peserta` : 'Tidak dibatasi'}</dd>
					<dt>Biaya</dt>
					<dd class="price">{harga(k.tipe_harga, k.harga)}</dd>
				</dl>
				{#if k.status === 'aktif'}
					<a class="btn btn-primary block" href="/auth/login">Daftar Kelas</a>
				{:else}
					<button class="btn btn-ghost block" disabled>Pendaftaran ditutup</button>
				{/if}
			</div>
		</aside>
	</div>
</section>

<style>
	.detail {
		display: grid;
		grid-template-columns: 1.6fr 0.9fr;
		gap: 40px;
		align-items: start;
	}
	.back {
		display: inline-block;
		color: var(--muted);
		font-size: 14px;
		margin-bottom: 20px;
	}
	.back:hover {
		color: var(--sky-blue);
	}
	.badges {
		display: flex;
		gap: 8px;
		flex-wrap: wrap;
		margin-bottom: 16px;
	}
	.badge.free {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
		border-color: rgba(30, 123, 69, 0.2);
	}
	.detail h1 {
		font-size: clamp(26px, 4vw, 38px);
	}
	.sub {
		font-size: 18px;
		margin: 28px 0 10px;
	}
	.body {
		color: var(--cool-slate);
		white-space: pre-line;
	}
	.side {
		position: sticky;
		top: 92px;
	}
	dl {
		display: grid;
		grid-template-columns: 1fr;
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
		color: var(--charcoal);
	}
	dd.price {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-size: 20px;
		color: var(--navy-teal);
	}
	.block {
		display: flex;
		justify-content: center;
		width: 100%;
		margin-top: 24px;
	}
	button.block[disabled] {
		opacity: 0.6;
		cursor: not-allowed;
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
