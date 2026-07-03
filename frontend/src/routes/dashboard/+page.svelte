<script lang="ts">
	import { onMount } from 'svelte';
	import { apiAuth } from '$lib/authApi';
	import { tanggal } from '$lib/format';
	import { ArrowRight, ArrowUpRight, Check, Flame, Zap } from 'lucide-svelte';
	import type { DashboardData, PendaftaranItem } from '$lib/types';

	let dash = $state<DashboardData | null>(null);
	let kelas = $state<PendaftaranItem[]>([]);
	let err = $state('');

	const today = new Date().toLocaleDateString('en-CA'); // YYYY-MM-DD

	// progress ring: hari aktif minggu ini / 7
	const R = 34;
	const C = 2 * Math.PI * R;
	const pct = $derived((dash?.progress.hari_aktif_minggu_ini ?? 0) / 7);
	const dash_off = $derived(C * (1 - pct));

	const banner = $derived(dash?.pengumuman.find((p) => p.tipe === 'banner'));
	const infoCards = $derived(dash?.pengumuman.filter((p) => p.tipe === 'info') ?? []);

	onMount(async () => {
		try {
			[dash, kelas] = await Promise.all([
				apiAuth<DashboardData>('/me/dashboard'),
				apiAuth<PendaftaranItem[]>('/pendaftaran/saya')
			]);
		} catch (e) {
			err = (e as Error).message;
		}
	});
</script>

<svelte:head><title>Beranda — TCC ITPLN</title></svelte:head>

{#if err}
	<div class="panel">Gagal memuat data: {err}</div>
{:else if dash}
	<div class="layout">
		<div class="col-main">
			<section class="banner">
				<div class="banner-text">
					<span class="banner-eyebrow">{banner ? 'Pengumuman' : 'Selamat Datang'}</span>
					<h1 class="banner-title">{banner ? banner.judul : `Halo, ${dash.profil.full_name}`}</h1>
					<p class="banner-desc">
						{banner ? banner.isi : 'Kelola kelas dan konsultasi pelatihan Anda di satu tempat.'}
					</p>
				</div>
				<a class="banner-cta" href={banner?.url_aksi || '/dashboard/jelajah'}>
					{banner?.label_aksi || 'Jelajahi Kelas'} <ArrowRight size={17} />
				</a>
			</section>

			<div class="panel misi-card">
				<div class="card-head">
					<h2>Misi Hari Ini</h2>
					<span class="misi-badge">{dash.misi.selesai}/{dash.misi.total} selesai</span>
				</div>
				{#if dash.misi.items.length === 0}
					<div class="empty">Belum ada misi hari ini.</div>
				{:else}
					<div class="misi-list">
						{#each dash.misi.items as m (m.id)}
							<div class="misi-row" class:done={m.selesai}>
								<div class="misi-ic">
									{#if m.selesai}<Check size={18} strokeWidth={2.5} />{:else}<Zap size={18} />{/if}
								</div>
								<div class="misi-body">
									<div class="misi-top">
										<span class="misi-judul">{m.judul}</span>
										<span class="misi-xp">+{m.xp_reward} XP</span>
									</div>
									<div class="misi-bar"><span style="width:{Math.min(100, (m.progres / m.target) * 100)}%"></span></div>
									<div class="misi-meta">{m.progres}/{m.target}</div>
								</div>
							</div>
						{/each}
					</div>
				{/if}
			</div>

			<div class="panel list-card">
				<div class="card-head">
					<h2>Kelas Terbaru</h2>
					<a class="more" href="/dashboard/kelas">Lihat semua <ArrowRight size={15} /></a>
				</div>
				{#if kelas.length === 0}
					<div class="empty">Belum ada kelas. <a href="/dashboard/jelajah">Jelajahi kelas</a>.</div>
				{:else}
					<div class="rows">
						{#each kelas.slice(0, 5) as it (it.pendaftaran_id)}
							<a class="row" href="/dashboard/jelajah/{it.kelas.slug}">
								<div>
									<div class="judul">{it.kelas.judul}</div>
									<div class="meta">{it.kelas.format} · {tanggal(it.tanggal_daftar)}</div>
								</div>
								<span class="status status-{it.status}">{it.status}</span>
							</a>
						{/each}
					</div>
				{/if}
			</div>
		</div>

		<aside class="col-rail">
			<div class="panel target-card">
				<div class="target-head">
					<div class="ring-wrap">
						<svg viewBox="0 0 80 80" class="ring">
							<circle class="ring-bg" cx="40" cy="40" r={R} />
							<circle class="ring-fg" cx="40" cy="40" r={R} stroke-dasharray={C} stroke-dashoffset={dash_off} />
						</svg>
						<div class="ring-center">
							<span class="ring-num">{dash.progress.hari_aktif_minggu_ini}</span>
							<span class="ring-den">/7</span>
						</div>
					</div>
					<div class="target-meta">
						<div class="target-label">Target Harian</div>
						<div class="target-sub">{dash.progress.hari_aktif_minggu_ini}/7 hari aktif minggu ini</div>
						<div class="target-lv">
							<Flame size={14} /> Lv {dash.progress.level} · {dash.progress.xp} XP
						</div>
					</div>
				</div>
				<div class="dots">
					{#each dash.progress.aktivitas_minggu as d (d.tanggal)}
						<div class="dot-col">
							<div class="dot" class:on={d.aktif} class:now={d.tanggal === today}></div>
							<span class="dot-lbl">{d.hari}</span>
						</div>
					{/each}
				</div>
				{#if dash.progress.streak_saat_ini > 0}
					<div class="streak">Streak {dash.progress.streak_saat_ini} hari · rekor {dash.progress.streak_terpanjang}</div>
				{/if}
			</div>

			{#each infoCards as info (info.id)}
				<div class="panel rail-block">
					<div class="rail-head">
						<h3>{info.judul}</h3>
						{#if info.url_aksi}
							<a href={info.url_aksi} aria-label={info.label_aksi}><ArrowUpRight size={16} /></a>
						{/if}
					</div>
					<p class="rail-text">{info.isi}</p>
				</div>
			{/each}

			<div class="info-card">
				<h3>Butuh solusi pelatihan?</h3>
				<p>Tim TCC ITPLN siap merancang program sesuai kebutuhan Anda.</p>
				<a class="info-cta" href="/konsultasi">Ajukan Konsultasi <ArrowRight size={15} /></a>
			</div>
		</aside>
	</div>
{/if}

<style>
	.layout {
		display: grid;
		grid-template-columns: 1fr 320px;
		gap: 22px;
		align-items: start;
	}
	.col-main {
		display: flex;
		flex-direction: column;
		gap: 20px;
		min-width: 0;
	}
	.col-rail {
		display: flex;
		flex-direction: column;
		gap: 18px;
		position: sticky;
		top: 30px;
	}

	/* banner */
	.banner {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 20px;
		padding: 26px 28px;
		border-radius: 18px;
		background: linear-gradient(120deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
		overflow: hidden;
		position: relative;
	}
	.banner::after {
		content: '';
		position: absolute;
		right: -60px;
		top: -60px;
		width: 240px;
		height: 240px;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.1);
		pointer-events: none;
	}
	.banner-text {
		position: relative;
	}
	.banner-eyebrow {
		font-family: var(--font-mono);
		font-size: 12px;
		letter-spacing: 0.08em;
		text-transform: uppercase;
		color: rgba(255, 255, 255, 0.8);
	}
	.banner-title {
		font-size: 24px;
		font-weight: 800;
		color: #fff;
		margin: 6px 0 6px;
		letter-spacing: -0.02em;
	}
	.banner-desc {
		font-size: 14.5px;
		color: rgba(255, 255, 255, 0.85);
		max-width: 460px;
	}
	.banner-cta {
		flex-shrink: 0;
		position: relative;
		display: inline-flex;
		align-items: center;
		gap: 7px;
		background: #fff;
		color: var(--navy-teal);
		font-weight: 600;
		font-size: 14.5px;
		padding: 12px 20px;
		border-radius: 999px;
		transition:
			transform 0.2s ease,
			box-shadow 0.2s ease;
	}
	.banner-cta:hover {
		transform: translateY(-2px);
		box-shadow: 0 12px 26px rgba(0, 0, 0, 0.18);
	}

	/* misi hari ini */
	.card-head {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 16px;
	}
	.card-head h2 {
		font-size: 16px;
		font-weight: 700;
		color: var(--ink);
	}
	.misi-badge {
		font-size: 12px;
		font-weight: 600;
		color: var(--sky-blue);
		background: rgba(26, 141, 178, 0.1);
		padding: 4px 11px;
		border-radius: 999px;
	}
	.misi-list {
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	.misi-row {
		display: flex;
		gap: 14px;
		align-items: flex-start;
		padding: 14px;
		border: 1px solid var(--border);
		border-radius: 13px;
		transition: border-color 0.2s ease;
	}
	.misi-row.done {
		background: rgba(30, 123, 69, 0.05);
		border-color: rgba(30, 123, 69, 0.2);
	}
	.misi-ic {
		width: 40px;
		height: 40px;
		border-radius: 11px;
		display: grid;
		place-items: center;
		flex-shrink: 0;
		background: rgba(26, 141, 178, 0.12);
		color: var(--sky-blue);
	}
	.misi-row.done .misi-ic {
		background: rgba(30, 123, 69, 0.15);
		color: var(--success);
	}
	.misi-body {
		flex: 1;
		min-width: 0;
	}
	.misi-top {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 10px;
	}
	.misi-judul {
		font-weight: 600;
		font-size: 14.5px;
		color: var(--ink);
	}
	.misi-xp {
		font-size: 12px;
		font-weight: 600;
		color: var(--muted);
		white-space: nowrap;
	}
	.misi-bar {
		height: 6px;
		border-radius: 999px;
		background: var(--border);
		margin: 9px 0 5px;
		overflow: hidden;
	}
	.misi-bar span {
		display: block;
		height: 100%;
		border-radius: 999px;
		background: linear-gradient(90deg, var(--navy-teal), var(--sky-blue));
		transition: width 0.5s cubic-bezier(0.22, 1, 0.36, 1);
	}
	.misi-row.done .misi-bar span {
		background: var(--success);
	}
	.misi-meta {
		font-size: 12px;
		color: var(--muted);
	}

	/* kelas terbaru */
	.more {
		display: inline-flex;
		align-items: center;
		gap: 5px;
		font-size: 13px;
		font-weight: 600;
		color: var(--sky-blue);
	}
	.more:hover {
		color: var(--navy-teal);
	}
	.rows {
		display: flex;
		flex-direction: column;
	}
	.row {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 16px;
		padding: 13px 12px;
		margin: 0 -12px;
		border-radius: 11px;
		border-top: 1px solid var(--border);
		transition: background 0.18s ease;
	}
	.row:first-child {
		border-top: none;
	}
	.row:hover {
		background: var(--off-white);
	}
	.judul {
		font-weight: 600;
		font-size: 15px;
		color: var(--ink);
	}
	.meta {
		color: var(--muted);
		font-size: 13px;
		margin-top: 3px;
	}
	.empty {
		color: var(--muted);
		padding: 6px 0;
	}
	.status {
		font-size: 12px;
		font-weight: 600;
		padding: 4px 11px;
		border-radius: 999px;
		text-transform: capitalize;
		background: rgba(46, 74, 90, 0.08);
		color: var(--cool-slate);
		white-space: nowrap;
	}
	.status-aktif {
		background: rgba(30, 123, 69, 0.12);
		color: var(--success);
	}
	.status-dibatalkan {
		background: rgba(220, 38, 38, 0.1);
		color: #dc2626;
	}

	/* target harian */
	.target-head {
		display: flex;
		align-items: center;
		gap: 16px;
	}
	.ring-wrap {
		position: relative;
		width: 80px;
		height: 80px;
		flex-shrink: 0;
	}
	.ring {
		width: 80px;
		height: 80px;
		transform: rotate(-90deg);
	}
	.ring-bg {
		fill: none;
		stroke: var(--border);
		stroke-width: 8;
	}
	.ring-fg {
		fill: none;
		stroke: var(--sky-blue);
		stroke-width: 8;
		stroke-linecap: round;
		transition: stroke-dashoffset 0.6s cubic-bezier(0.22, 1, 0.36, 1);
	}
	.ring-center {
		position: absolute;
		inset: 0;
		display: grid;
		place-items: center;
		font-family: var(--font-display);
	}
	.ring-num {
		font-size: 22px;
		font-weight: 800;
		color: var(--ink);
	}
	.ring-den {
		font-size: 13px;
		color: var(--muted);
	}
	.target-label {
		font-size: 15px;
		font-weight: 700;
		color: var(--ink);
	}
	.target-sub {
		font-size: 12.5px;
		color: var(--muted);
		margin-top: 2px;
	}
	.target-lv {
		display: inline-flex;
		align-items: center;
		gap: 5px;
		font-size: 12.5px;
		font-weight: 600;
		color: var(--sky-blue);
		margin-top: 8px;
	}
	.dots {
		display: flex;
		justify-content: space-between;
		margin-top: 20px;
	}
	.dot-col {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 6px;
	}
	.dot {
		width: 26px;
		height: 26px;
		border-radius: 50%;
		background: var(--off-white);
		border: 1px solid var(--border);
	}
	.dot.on {
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		border-color: transparent;
	}
	.dot.now {
		box-shadow: 0 0 0 2px rgba(26, 141, 178, 0.35);
	}
	.dot-lbl {
		font-size: 11px;
		color: var(--muted);
	}
	.streak {
		margin-top: 16px;
		padding-top: 14px;
		border-top: 1px solid var(--border);
		font-size: 12.5px;
		color: var(--cool-slate);
	}

	/* rail info blocks */
	.rail-head {
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 8px;
	}
	.rail-head h3 {
		font-size: 14.5px;
		font-weight: 700;
		color: var(--ink);
	}
	.rail-head a {
		color: var(--muted);
		display: inline-flex;
	}
	.rail-head a:hover {
		color: var(--sky-blue);
	}
	.rail-text {
		font-size: 13.5px;
		color: var(--muted);
	}

	/* info card */
	.info-card {
		border-radius: 16px;
		padding: 22px;
		background: var(--off-white);
		border: 1px solid var(--border);
	}
	.info-card h3 {
		font-size: 15.5px;
		font-weight: 700;
		color: var(--ink);
	}
	.info-card p {
		font-size: 13.5px;
		color: var(--muted);
		margin: 8px 0 16px;
	}
	.info-cta {
		display: inline-flex;
		align-items: center;
		gap: 6px;
		font-size: 13.5px;
		font-weight: 600;
		color: var(--sky-blue);
	}
	.info-cta:hover {
		color: var(--navy-teal);
	}

	@media (max-width: 1100px) {
		.layout {
			grid-template-columns: 1fr;
		}
		.col-rail {
			position: static;
		}
	}
	@media (max-width: 620px) {
		.banner {
			flex-direction: column;
			align-items: flex-start;
		}
	}
</style>
