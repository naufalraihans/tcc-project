<script lang="ts">
	// taburan partikel hero (deterministik → aman untuk SSR/hydration)
	const dotColors = ['#0c4f6a', '#1a8db2', '#6fc4dd', '#2e4a5a'];
	let dotSeed = 7;
	const dotRand = () => {
		dotSeed = (dotSeed * 1664525 + 1013904223) % 4294967296;
		return dotSeed / 4294967296;
	};
	const dots = Array.from({ length: 72 }, () => ({
		top: +(dotRand() * 100).toFixed(2),
		left: +(dotRand() * 100).toFixed(2),
		len: +(6 + dotRand() * 12).toFixed(1),
		rot: Math.floor(dotRand() * 360),
		color: dotColors[Math.floor(dotRand() * dotColors.length)],
		op: +(0.2 + dotRand() * 0.45).toFixed(2),
		depth: +(8 + dotRand() * 30).toFixed(1) // parallax: makin besar makin jauh geraknya
	}));

	// posisi mouse relatif hero (-1..1), menggerakkan partikel (parallax)
	let mx = $state(0);
	let my = $state(0);
	function heroMove(e: MouseEvent) {
		const r = (e.currentTarget as HTMLElement).getBoundingClientRect();
		mx = ((e.clientX - r.left) / r.width - 0.5) * 2;
		my = ((e.clientY - r.top) / r.height - 0.5) * 2;
	}
	function heroLeave() {
		mx = 0;
		my = 0;
	}

	const stats = [
		{ value: '6.042', label: 'peserta dalam 3 tahun' },
		{ value: '11', label: 'bidang pelatihan' },
		{ value: '100+', label: 'program di katalog' },
		{ value: '6', label: 'program sertifikasi' }
	];

	const pilar = [
		{ kode: 'S/W', nama: 'Software', sub: 'Learning Program', desc: 'Konten dan kurikulum yang dirancang untuk pencapaian kompetensi yang terukur.' },
		{ kode: 'B/W', nama: 'Brainware', sub: 'Facilitator', desc: 'Instruktur, coach, dan mentor berpengalaman di bidangnya masing-masing.' },
		{ kode: 'H/W', nama: 'Hardware', sub: 'Learning Infrastructure', desc: 'Smart classroom, perangkat teknologi, dan akses digital yang menunjang.' }
	];

	const output = ['Competence', 'Performance', 'Certification'];

	const unggulan = [
		{ judul: 'Manajemen Aset Pembangkit ISO 55001:2014', bidang: 'Manajemen Aset', jp: '16 JP' },
		{ judul: 'Workshop K3 dan 5S', bidang: 'OHS / K3', jp: '24 JP' },
		{ judul: 'Pembangkit Listrik Tenaga Surya (PLTS)', bidang: 'EBT', jp: '8 JP' },
		{ judul: 'International Training AMI / Smart Grid', bidang: 'Digital Platform', jp: '40 JP' },
		{ judul: 'Waste to Energy (WTE)', bidang: 'EBT', jp: '24 JP' },
		{ judul: 'Basic Theory Power Plant', bidang: 'Engineering', jp: '8 JP' }
	];

	const topik = [
		'Engineering', 'Digital Platform', 'Manajemen Aset', 'Manajemen Risiko',
		'Manajemen Keuangan', 'Sumber Daya Manusia', 'EPC', 'Renewable Energy',
		'OHS / K3', 'Bimtek', 'Workshop'
	];

	const partners = [
		{ src: '/partners/danantara.svg', alt: 'Danantara' },
		{ src: '/partners/pln.png', alt: 'PLN' },
		{ src: '/partners/pln-icon-plus.png', alt: 'PLN Icon Plus' },
		{ src: '/partners/pln-ips.png', alt: 'PLN Indonesia Power Services' },
		{ src: '/partners/pln-npc.png', alt: 'PLN Nusantara Power Construction' },
		{ src: '/partners/pln-mctn.png', alt: 'PLN MCTN' }
	];
</script>

<svelte:head>
	<title>TCC ITPLN — Training &amp; Consulting Center</title>
	<meta name="description" content="Pusat pelatihan dan konsultasi Institut Teknologi PLN di bidang energi, teknologi berwawasan lingkungan, dan manajemen." />
</svelte:head>

<section class="hero" role="presentation" onmousemove={heroMove} onmouseleave={heroLeave}>
	<div class="particles" aria-hidden="true" style="--mx:{mx};--my:{my}">
		{#each dots as d}
			<span
				class="dot"
				style="top:{d.top}%;left:{d.left}%;width:{d.len}px;background:{d.color};opacity:{d.op};transform:translate(calc(var(--mx) * {d.depth}px), calc(var(--my) * {d.depth}px)) rotate({d.rot}deg)"
			></span>
		{/each}
	</div>

	<div class="container hero-inner">
		<h1 class="hero-title rise" style="animation-delay:.15s">
			Kembangkan kompetensi<br /><span class="hl">ketenagalistrikan</span>
		</h1>
		<p class="hero-desc rise" style="animation-delay:.28s">
			Pusat pelatihan dan konsultasi berkelas internasional di bidang energi, teknologi berwawasan
			lingkungan, dan manajemen — untuk individu maupun organisasi.
		</p>
		<div class="hero-cta rise" style="animation-delay:.4s">
			<a class="btn btn-primary" href="/kelas">Lihat Program</a>
			<a class="btn btn-ghost" href="/konsultasi">Ajukan Konsultasi</a>
		</div>
	</div>
</section>

<section class="stats-strip">
	<div class="marquee-wrap">
		<div class="marquee-track">
			{#each [0, 1] as g}
				<div class="marquee-group" aria-hidden={g === 1}>
					{#each stats as s}
						<div class="mstat">
							<span class="mvalue">{s.value}</span>
							<span class="mlabel">{s.label}</span>
						</div>
					{/each}
				</div>
			{/each}
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<span class="eyebrow">Kenapa TCC ITPLN</span>
		<h2 class="section-title">Tiga pilar pembelajaran<br />yang terintegrasi</h2>
		<p class="section-lead">
			Kerangka pembelajaran TCC memadukan konten, pengajar, dan infrastruktur untuk menghasilkan
			kompetensi yang nyata.
		</p>

		<div class="grid pilar-grid">
			{#each pilar as p}
				<div class="card">
					<span class="badge">{p.kode}</span>
					<h3 class="pilar-nama">{p.nama}</h3>
					<div class="pilar-sub">{p.sub}</div>
					<p class="pilar-desc">{p.desc}</p>
				</div>
			{/each}
		</div>

		<div class="output-row">
			{#each output as o, i}
				<div class="output-item"><span class="onum">0{i + 1}</span>{o}</div>
			{/each}
		</div>
	</div>
</section>

<section class="section soft">
	<div class="container">
		<div class="head-row">
			<div>
				<span class="eyebrow">Kelas Unggulan</span>
				<h2 class="section-title">Program pilihan dari katalog</h2>
			</div>
			<a class="btn btn-ghost" href="/kelas">Semua Kelas</a>
		</div>

		<div class="grid unggulan-grid">
			{#each unggulan as k}
				<a class="card kelas-card" href="/kelas">
					<span class="badge">{k.bidang}</span>
					<h3 class="kelas-judul">{k.judul}</h3>
					<div class="kelas-meta">{k.jp}</div>
				</a>
			{/each}
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<span class="eyebrow">Bidang Pelatihan</span>
		<h2 class="section-title">Sebelas bidang, satu ekosistem energi</h2>
		<div class="topik-wrap">
			{#each topik as t}
				<a class="chip" href="/topik">{t}</a>
			{/each}
		</div>
	</div>
</section>

<section class="trusted">
	<div class="container">
		<div class="trusted-label">
			<span class="tline"></span>
			Dipercaya oleh Instansi Engineering dan Administrasi
			<span class="tline"></span>
		</div>
	</div>
	<div class="logo-marquee">
		<div class="marquee-track">
			{#each [0, 1] as g}
				<div class="logo-group" aria-hidden={g === 1}>
					{#each partners as p}
						<div class="logo-item"><img src={p.src} alt={p.alt} /></div>
					{/each}
				</div>
			{/each}
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<div class="cta">
			<div class="cta-lines"></div>
			<div class="cta-content">
				<h2 class="cta-title">Butuh solusi pelatihan untuk organisasi Anda?</h2>
				<p class="cta-lead">
					Tim TCC ITPLN siap merancang program tailor-made sesuai kebutuhan kompetensi perusahaan.
				</p>
				<a class="btn cta-btn" href="/konsultasi">Ajukan Konsultasi</a>
			</div>
		</div>
	</div>
</section>

<style>
	.hero {
		position: relative;
		min-height: 88vh;
		display: flex;
		align-items: center;
		justify-content: center;
		text-align: center;
		overflow: hidden;
		background:
			radial-gradient(ellipse 70% 60% at 50% 38%, #ffffff 52%, transparent),
			linear-gradient(180deg, #eef3f8, #ffffff 55%);
	}
	.particles {
		position: absolute;
		inset: 0;
		pointer-events: none;
		z-index: 0;
		-webkit-mask-image: radial-gradient(ellipse 60% 56% at 50% 46%, transparent 30%, #000 76%);
		mask-image: radial-gradient(ellipse 60% 56% at 50% 46%, transparent 30%, #000 76%);
	}
	.dot {
		position: absolute;
		height: 3px;
		border-radius: 2px;
		transform-origin: center;
		transition: transform 0.28s cubic-bezier(0.22, 1, 0.36, 1);
		will-change: transform;
	}
	.hero-inner {
		position: relative;
		z-index: 1;
		padding: 90px 24px;
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.hero-title {
		font-size: clamp(2.6rem, 7.4vw, 5.6rem);
		font-weight: 600;
		line-height: 1.03;
		letter-spacing: -0.03em;
		color: var(--ink);
		max-width: 16ch;
	}
	.hl {
		background: linear-gradient(120deg, var(--navy-teal), var(--sky-blue));
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.hero-desc {
		font-size: 18px;
		color: var(--muted);
		max-width: 560px;
		margin-top: 26px;
	}
	.hero-cta {
		display: flex;
		gap: 14px;
		flex-wrap: wrap;
		justify-content: center;
		margin-top: 36px;
	}
	.hero-cta .btn {
		height: 52px;
		padding: 0 28px;
		font-size: 15px;
	}
	.stats-strip {
		background: var(--white);
	}

	.marquee-wrap {
		position: relative;
		z-index: 2;
		border-top: 1px solid var(--border);
		padding: 26px 0;
		overflow: hidden;
		mask-image: linear-gradient(90deg, transparent, black 8%, black 92%, transparent);
	}
	.marquee-group {
		display: flex;
		gap: 64px;
		padding-right: 64px;
	}
	.mstat {
		display: flex;
		align-items: baseline;
		gap: 12px;
		white-space: nowrap;
	}
	.mvalue {
		font-family: var(--font-display);
		font-weight: 800;
		font-size: 34px;
		color: var(--ink);
	}
	.mlabel {
		font-size: 14px;
		color: var(--muted);
	}

	.trusted {
		padding: 56px 0 52px;
		background: var(--white);
		border-top: 1px solid var(--border);
		border-bottom: 1px solid var(--border);
	}
	.trusted-label {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 18px;
		font-family: var(--font-mono);
		font-size: 13px;
		letter-spacing: 0.02em;
		color: var(--muted);
		text-align: center;
	}
	.tline {
		height: 1px;
		width: 44px;
		background: var(--border-strong);
		flex-shrink: 0;
	}
	.logo-marquee {
		margin-top: 34px;
		overflow: hidden;
		mask-image: linear-gradient(90deg, transparent, black 8%, black 92%, transparent);
	}
	.logo-group {
		display: flex;
		align-items: center;
		gap: 72px;
		padding-right: 72px;
	}
	.logo-item {
		display: flex;
		align-items: center;
	}
	.logo-item img {
		height: 40px;
		width: auto;
		object-fit: contain;
		opacity: 0.85;
	}

	.pilar-grid,
	.unggulan-grid {
		grid-template-columns: repeat(3, 1fr);
		margin-top: 44px;
	}
	.pilar-nama {
		font-size: 24px;
		margin-top: 20px;
	}
	.pilar-sub {
		color: var(--sky-blue);
		font-weight: 600;
		font-size: 14px;
		margin-top: 4px;
		font-family: var(--font-mono);
	}
	.pilar-desc {
		color: var(--muted);
		margin-top: 14px;
		font-size: 15px;
	}
	.output-row {
		display: flex;
		gap: 16px;
		margin-top: 22px;
		flex-wrap: wrap;
	}
	.output-item {
		flex: 1;
		min-width: 180px;
		display: flex;
		align-items: center;
		gap: 12px;
		padding: 22px 24px;
		border-radius: var(--radius-sm);
		background: var(--ink);
		color: #fff;
		font-family: var(--font-display);
		font-weight: 700;
		font-size: 17px;
	}
	.onum {
		font-family: var(--font-mono);
		font-size: 13px;
		font-weight: 500;
		color: var(--sky-blue);
	}

	.soft {
		background: var(--off-white);
		border-top: 1px solid var(--border);
		border-bottom: 1px solid var(--border);
	}
	.head-row {
		display: flex;
		justify-content: space-between;
		align-items: flex-end;
		gap: 20px;
		flex-wrap: wrap;
	}
	.kelas-card {
		display: flex;
		flex-direction: column;
	}
	.kelas-judul {
		font-size: 20px;
		margin-top: 18px;
		flex: 1;
	}
	.kelas-meta {
		margin-top: 18px;
		font-weight: 600;
		color: var(--cool-slate);
		font-family: var(--font-mono);
		font-size: 13px;
	}

	.topik-wrap {
		display: flex;
		flex-wrap: wrap;
		gap: 12px;
		margin-top: 36px;
	}
	.chip {
		padding: 11px 20px;
		border-radius: var(--radius-pill);
		border: 1px solid var(--border-strong);
		background: var(--white);
		font-weight: 500;
		color: var(--cool-slate);
		transition: all 0.2s ease;
	}
	.chip:hover {
		border-color: var(--navy-teal);
		color: var(--navy-teal);
		transform: translateY(-2px);
	}

	.cta {
		position: relative;
		overflow: hidden;
		border-radius: 28px;
		background: linear-gradient(135deg, var(--ink), var(--navy-teal) 70%, var(--sky-blue));
		padding: 72px 56px;
	}
	.cta-lines {
		position: absolute;
		inset: 0;
		opacity: 0.12;
		background-image:
			repeating-linear-gradient(to right, transparent 0, transparent 79px, #fff 79px, #fff 80px);
	}
	.cta-content {
		position: relative;
		max-width: 640px;
	}
	.cta-title {
		color: #fff;
		font-size: clamp(26px, 3.4vw, 38px);
	}
	.cta-lead {
		color: rgba(255, 255, 255, 0.82);
		margin-top: 14px;
		font-size: 17px;
	}
	.cta-btn {
		margin-top: 28px;
		background: #fff;
		color: var(--ink);
		height: 54px;
		padding: 0 30px;
	}
	.cta-btn:hover {
		transform: translateY(-2px);
		box-shadow: 0 12px 30px rgba(0, 0, 0, 0.25);
	}

	@media (max-width: 900px) {
		.pilar-grid,
		.unggulan-grid {
			grid-template-columns: 1fr;
		}
		.cta {
			padding: 44px 28px;
		}
	}
</style>
