<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '$lib/api';
	import { harga, tanggal } from '$lib/format';
	import { ArrowRight, Check } from 'lucide-svelte';
	import type { PagedKelas, KelasListItem, Topik } from '$lib/types';

	let kelas = $state<KelasListItem[]>([]);
	let topik = $state<Topik[]>([]);
	let sel = $state(''); // slug topik terpilih; '' = semua
	let err = $state('');
	let loading = $state(true);

	type Group = { topik: Topik | null; items: KelasListItem[] };

	const groups = $derived.by<Group[]>(() => {
		const bySlug = new Map<string, Group>();
		for (const t of topik) bySlug.set(t.slug, { topik: t, items: [] });
		const lain: Group = { topik: null, items: [] };
		for (const k of kelas) {
			const g = (k.topik && bySlug.get(k.topik.slug)) || lain;
			g.items.push(k);
		}
		const out = [...bySlug.values()].filter((g) => g.items.length > 0);
		if (lain.items.length) out.push(lain);
		return sel ? out.filter((g) => g.topik?.slug === sel) : out;
	});

	// gradient thumbnail biru-teal (rotasi kecil, tetap dalam palet)
	const grads = [
		'linear-gradient(135deg, #0c4f6a, #1a8db2)',
		'linear-gradient(135deg, #14657f, #2e4a5a)',
		'linear-gradient(135deg, #1a8db2, #6fc4dd)',
		'linear-gradient(135deg, #0c4f6a, #2e4a5a)'
	];
	const gradOf = (s: string) => grads[[...s].reduce((a, c) => a + c.charCodeAt(0), 0) % grads.length];

	// ponytail: dummy thumbnail (deterministik per slug) sampai kolom thumbnail_url ada di DB
	const thumb = (k: KelasListItem) => `https://picsum.photos/seed/${k.slug}/640/360`;

	function bullets(k: KelasListItem): string[] {
		const b = [`Format ${k.format}`];
		if (k.jadwal_mulai) b.push(`Mulai ${tanggal(k.jadwal_mulai)}`);
		if (k.instruktur) b.push(`Instruktur ${k.instruktur.nama}`);
		if (k.kuota > 0) b.push(`Sisa ${Math.max(k.kuota - k.peserta_terdaftar, 0)} kuota`);
		return b.slice(0, 3);
	}

	onMount(async () => {
		try {
			const [k, t] = await Promise.all([
				api<PagedKelas>('/kelas?limit=100'),
				api<Topik[]>('/topik')
			]);
			kelas = k.items;
			topik = t;
		} catch (e) {
			err = (e as Error).message;
		} finally {
			loading = false;
		}
	});
</script>

<svelte:head><title>Jelajahi Kelas — TCC ITPLN</title></svelte:head>

<h1>Jelajahi Kelas</h1>
<p class="lead">Pilih program pelatihan sesuai bidang dan kebutuhan kompetensi Anda.</p>

{#if err}
	<div class="panel">Gagal memuat kelas: {err}</div>
{:else}
	<div class="chips">
		<button class="chip" class:active={sel === ''} onclick={() => (sel = '')}>Semua</button>
		{#each topik as t (t.id)}
			<button class="chip" class:active={sel === t.slug} onclick={() => (sel = t.slug)}>{t.nama}</button>
		{/each}
	</div>

	{#if loading}
		<div class="panel">Memuat…</div>
	{:else if groups.length === 0}
		<div class="panel empty">Belum ada kelas pada kategori ini.</div>
	{:else}
		{#each groups as g (g.topik?.slug ?? 'lain')}
			<section class="cat">
				<div class="cat-head">
					<div>
						<h2>{g.topik?.nama ?? 'Lainnya'}</h2>
						{#if g.topik?.deskripsi}<p>{g.topik.deskripsi}</p>{/if}
					</div>
					<span class="count">{g.items.length} program</span>
				</div>

				<div class="grid">
					{#each g.items as k (k.id)}
						<a class="card" href="/dashboard/jelajah/{k.slug}">
							<div class="thumb" style="background:{gradOf(k.slug)}">
								<img class="thumb-img" src={thumb(k)} alt="" loading="lazy" />
								<span class="thumb-badge">{harga(k.tipe_harga, k.harga)}</span>
							</div>
							<div class="body">
								<h3>{k.judul}</h3>
								<ul class="feat">
									{#each bullets(k) as f}
										<li><Check size={15} strokeWidth={2.5} /> {f}</li>
									{/each}
								</ul>
								<div class="price-row">
									<span class="price-lbl">Biaya</span>
									<span class="price">{harga(k.tipe_harga, k.harga)}</span>
								</div>
								<span class="cta">Lihat Detail <ArrowRight size={16} /></span>
							</div>
						</a>
					{/each}
				</div>
			</section>
		{/each}
	{/if}
{/if}

<style>
	.chips {
		display: flex;
		flex-wrap: wrap;
		gap: 10px;
		margin-bottom: 26px;
	}
	.chip {
		padding: 9px 17px;
		border-radius: 999px;
		border: 1px solid var(--border);
		background: var(--white);
		font-size: 13.5px;
		font-weight: 500;
		color: var(--cool-slate);
		cursor: pointer;
		font-family: var(--font-sans);
		transition:
			background 0.18s ease,
			color 0.18s ease,
			border-color 0.18s ease;
	}
	.chip:hover {
		border-color: var(--sky-blue);
		color: var(--sky-blue);
	}
	.chip.active {
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
		border-color: transparent;
	}

	.cat {
		margin-bottom: 34px;
	}
	.cat-head {
		display: flex;
		align-items: flex-start;
		justify-content: space-between;
		gap: 16px;
		margin-bottom: 16px;
	}
	.cat-head h2 {
		font-size: 19px;
		font-weight: 800;
		color: var(--ink);
	}
	.cat-head p {
		font-size: 13.5px;
		color: var(--muted);
		margin-top: 3px;
		max-width: 560px;
	}
	.count {
		flex-shrink: 0;
		font-size: 12px;
		font-weight: 600;
		color: var(--cool-slate);
		background: var(--off-white);
		border: 1px solid var(--border);
		padding: 5px 12px;
		border-radius: 999px;
		white-space: nowrap;
	}

	.grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
		gap: 20px;
	}
	.card {
		display: flex;
		flex-direction: column;
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: 18px;
		overflow: hidden;
		transition:
			transform 0.28s cubic-bezier(0.34, 1.4, 0.64, 1),
			box-shadow 0.28s ease,
			border-color 0.2s ease;
	}
	.card:hover {
		transform: translateY(-5px);
		box-shadow: 0 22px 44px rgba(12, 79, 106, 0.14);
		border-color: var(--border-strong, #cfd8e3);
	}
	.thumb {
		position: relative;
		height: 150px;
		display: grid;
		place-items: center;
		overflow: hidden;
	}
	.thumb::after {
		content: '';
		position: absolute;
		inset: 0;
		background: radial-gradient(circle at 75% 20%, rgba(255, 255, 255, 0.18), transparent 55%);
	}
	.thumb-img {
		position: absolute;
		inset: 0;
		width: 100%;
		height: 100%;
		object-fit: cover;
		z-index: 0;
	}
	.thumb-badge {
		position: absolute;
		top: 12px;
		left: 12px;
		z-index: 1;
		font-size: 12px;
		font-weight: 700;
		color: var(--navy-teal);
		background: rgba(255, 255, 255, 0.92);
		padding: 4px 11px;
		border-radius: 999px;
	}
	.body {
		display: flex;
		flex-direction: column;
		padding: 18px 20px 20px;
		flex: 1;
	}
	.body h3 {
		font-size: 17px;
		font-weight: 700;
		color: var(--ink);
		line-height: 1.3;
	}
	.feat {
		list-style: none;
		margin: 14px 0 0;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: 8px;
		flex: 1;
	}
	.feat li {
		display: flex;
		align-items: center;
		gap: 8px;
		font-size: 13px;
		color: var(--cool-slate);
	}
	.feat li :global(svg) {
		color: var(--sky-blue);
		flex-shrink: 0;
	}
	.price-row {
		display: flex;
		align-items: baseline;
		justify-content: space-between;
		margin: 18px 0 14px;
		padding-top: 14px;
		border-top: 1px solid var(--border);
	}
	.price-lbl {
		font-size: 12px;
		color: var(--muted);
	}
	.price {
		font-family: var(--font-display);
		font-weight: 800;
		font-size: 19px;
		color: var(--navy-teal);
	}
	.cta {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: 7px;
		padding: 12px;
		border-radius: 12px;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
		font-weight: 600;
		font-size: 14px;
	}
	.card:hover .cta {
		box-shadow: 0 10px 22px rgba(12, 79, 106, 0.28);
	}
	.empty {
		color: var(--muted);
	}
</style>
