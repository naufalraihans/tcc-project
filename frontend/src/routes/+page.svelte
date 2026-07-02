<script lang="ts">
	const metrics = [
		{ angka: '6.042', label: 'Peserta dalam 3 tahun' },
		{ angka: '11', label: 'Bidang pelatihan' },
		{ angka: '100+', label: 'Program di katalog' },
		{ angka: '6', label: 'Program sertifikasi' }
	];

	const pilar = [
		{
			kode: 'S/W',
			nama: 'Software',
			sub: 'Learning Program',
			desc: 'Konten dan kurikulum yang dirancang untuk pencapaian kompetensi yang terukur.'
		},
		{
			kode: 'B/W',
			nama: 'Brainware',
			sub: 'Facilitator',
			desc: 'Instruktur, coach, dan mentor berpengalaman di bidangnya masing-masing.'
		},
		{
			kode: 'H/W',
			nama: 'Hardware',
			sub: 'Learning Infrastructure',
			desc: 'Smart classroom, perangkat teknologi, dan akses digital yang menunjang.'
		}
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
		'Engineering',
		'Digital Platform',
		'Manajemen Aset',
		'Manajemen Risiko',
		'Manajemen Keuangan',
		'Sumber Daya Manusia',
		'EPC',
		'Renewable Energy',
		'OHS / K3',
		'Bimtek',
		'Workshop'
	];
</script>

<svelte:head>
	<title>TCC ITPLN — Training &amp; Consulting Center</title>
	<meta
		name="description"
		content="Pusat pelatihan dan konsultasi Institut Teknologi PLN di bidang energi, teknologi berwawasan lingkungan, dan manajemen."
	/>
</svelte:head>

<section class="hero">
	<div class="container hero-grid">
		<div>
			<span class="section-tag">Institut Teknologi PLN</span>
			<h1>Kembangkan Kompetensi Ketenagalistrikan Bersama TCC ITPLN</h1>
			<p class="section-lead">
				Pusat pelatihan dan konsultasi berkelas internasional di bidang energi, teknologi berwawasan
				lingkungan, dan manajemen — dirancang untuk individu maupun organisasi.
			</p>
			<div class="hero-actions">
				<a class="btn btn-primary" href="/kelas">Lihat Program</a>
				<a class="btn btn-ghost" href="/konsultasi">Ajukan Konsultasi</a>
			</div>
			<div class="hero-stats">
				{#each metrics.slice(0, 3) as m}
					<div>
						<div class="angka">{m.angka}</div>
						<div class="stat-label">{m.label}</div>
					</div>
				{/each}
			</div>
		</div>
		<div class="hero-panel" aria-hidden="true">
			<div class="panel-inner">
				<div class="panel-kode">TCC</div>
				<div class="panel-sub">Training &amp; Consulting Center</div>
			</div>
		</div>
	</div>
</section>

<section class="section metrics-band">
	<div class="container grid metrics-grid">
		{#each metrics as m}
			<div class="metric">
				<div class="angka">{m.angka}</div>
				<div class="stat-label">{m.label}</div>
			</div>
		{/each}
	</div>
</section>

<section class="section">
	<div class="container">
		<span class="section-tag">Kenapa TCC ITPLN</span>
		<h2 class="section-title">Tiga pilar pembelajaran yang terintegrasi</h2>
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
			{#each output as o}
				<div class="output-item">{o}</div>
			{/each}
		</div>
	</div>
</section>

<section class="section soft">
	<div class="container">
		<span class="section-tag">Kelas Unggulan</span>
		<h2 class="section-title">Program pilihan dari katalog TCC</h2>

		<div class="grid unggulan-grid">
			{#each unggulan as k}
				<div class="card kelas-card">
					<span class="badge">{k.bidang}</span>
					<h3 class="kelas-judul">{k.judul}</h3>
					<div class="kelas-meta">{k.jp}</div>
				</div>
			{/each}
		</div>

		<div class="center-cta">
			<a class="btn btn-primary" href="/kelas">Lihat Semua Kelas</a>
		</div>
	</div>
</section>

<section class="section">
	<div class="container">
		<span class="section-tag">Bidang Pelatihan</span>
		<h2 class="section-title">Sebelas bidang, satu ekosistem energi</h2>
		<div class="topik-wrap">
			{#each topik as t}
				<a class="topik-chip" href="/topik">{t}</a>
			{/each}
		</div>
	</div>
</section>

<section class="section">
	<div class="container cta-box">
		<div>
			<h2 class="cta-title">Butuh solusi pelatihan untuk organisasi Anda?</h2>
			<p class="cta-lead">
				Tim TCC ITPLN siap merancang program tailor-made sesuai kebutuhan kompetensi perusahaan.
			</p>
		</div>
		<a class="btn btn-primary" href="/konsultasi">Ajukan Konsultasi</a>
	</div>
</section>

<style>
	.hero {
		padding: 80px 0 48px;
	}
	.hero-grid {
		display: grid;
		grid-template-columns: 1.15fr 0.85fr;
		gap: 48px;
		align-items: center;
	}
	.hero h1 {
		font-size: clamp(34px, 5vw, 54px);
		font-weight: 800;
		margin-top: 16px;
	}
	.hero-actions {
		display: flex;
		gap: 14px;
		margin-top: 28px;
		flex-wrap: wrap;
	}
	.hero-stats {
		display: flex;
		gap: 40px;
		margin-top: 44px;
	}
	.angka {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 30px;
		color: var(--navy-teal);
	}
	.stat-label {
		font-size: 14px;
		color: var(--muted);
		margin-top: 4px;
	}
	.hero-panel {
		aspect-ratio: 4 / 3;
		border-radius: 24px;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		display: grid;
		place-items: center;
		box-shadow: 0 24px 60px rgba(12, 79, 106, 0.25);
	}
	.panel-inner {
		text-align: center;
		color: #fff;
	}
	.panel-kode {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 64px;
		letter-spacing: -0.04em;
	}
	.panel-sub {
		font-size: 15px;
		opacity: 0.85;
		margin-top: 4px;
	}

	.metrics-band {
		padding: 40px 0;
		background: var(--off-white);
		border-top: 1px solid var(--border);
		border-bottom: 1px solid var(--border);
	}
	.metrics-grid {
		grid-template-columns: repeat(4, 1fr);
		text-align: center;
	}

	.soft {
		background: var(--off-white);
	}
	.pilar-grid,
	.unggulan-grid {
		grid-template-columns: repeat(3, 1fr);
		margin-top: 40px;
	}
	.pilar-nama {
		font-size: 22px;
		margin-top: 18px;
	}
	.pilar-sub {
		color: var(--sky-blue);
		font-weight: 600;
		font-size: 14px;
		margin-top: 4px;
	}
	.pilar-desc {
		color: var(--muted);
		margin-top: 12px;
		font-size: 15px;
	}
	.output-row {
		display: flex;
		gap: 16px;
		margin-top: 24px;
		flex-wrap: wrap;
	}
	.output-item {
		flex: 1;
		min-width: 160px;
		text-align: center;
		padding: 18px;
		border-radius: var(--radius-sm);
		background: var(--navy-teal);
		color: #fff;
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 700;
	}

	.kelas-card {
		display: flex;
		flex-direction: column;
	}
	.kelas-judul {
		font-size: 19px;
		margin-top: 16px;
		flex: 1;
	}
	.kelas-meta {
		margin-top: 16px;
		font-weight: 600;
		color: var(--cool-slate);
		font-size: 14px;
	}
	.center-cta {
		margin-top: 40px;
		text-align: center;
	}

	.topik-wrap {
		display: flex;
		flex-wrap: wrap;
		gap: 12px;
		margin-top: 32px;
	}
	.topik-chip {
		padding: 10px 18px;
		border-radius: var(--radius-full);
		border: 1px solid var(--border);
		background: var(--white);
		font-weight: 500;
		color: var(--cool-slate);
		transition: 0.18s ease;
	}
	.topik-chip:hover {
		border-color: var(--sky-blue);
		color: var(--sky-blue);
	}

	.cta-box {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 32px;
		padding: 48px;
		border-radius: 24px;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
	}
	.cta-title {
		color: #fff;
		font-size: clamp(24px, 3vw, 32px);
	}
	.cta-lead {
		margin-top: 10px;
		opacity: 0.9;
		max-width: 520px;
	}
	.cta-box .btn-primary {
		background: #fff;
		color: var(--navy-teal);
		white-space: nowrap;
	}

	@media (max-width: 900px) {
		.hero-grid {
			grid-template-columns: 1fr;
		}
		.hero-panel {
			display: none;
		}
		.pilar-grid,
		.unggulan-grid {
			grid-template-columns: 1fr;
		}
		.metrics-grid {
			grid-template-columns: repeat(2, 1fr);
			gap: 28px;
		}
		.cta-box {
			flex-direction: column;
			align-items: flex-start;
			padding: 32px;
		}
	}
</style>
