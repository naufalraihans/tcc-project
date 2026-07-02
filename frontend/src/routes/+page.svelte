<script lang="ts">
	const hlWord = 'ketenagalistrikan'.split('');

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
</script>

<svelte:head>
	<title>TCC ITPLN — Training &amp; Consulting Center</title>
	<meta name="description" content="Pusat pelatihan dan konsultasi Institut Teknologi PLN di bidang energi, teknologi berwawasan lingkungan, dan manajemen." />
</svelte:head>

<section class="hero">
	<div class="blob"></div>
	<div class="lines"></div>

	<div class="container hero-inner">
		<span class="eyebrow rise" style="animation-delay:.05s">Institut Teknologi PLN</span>

		<h1 class="hero-title">
			<span class="line rise" style="animation-delay:.15s">Kembangkan kompetensi</span>
			<span class="line hl-line">
				<span class="hl">
					{#each hlWord as ch, i}
						<span class="char" style="animation-delay:{350 + i * 35}ms">{ch}</span>
					{/each}
				</span>
			</span>
		</h1>

		<div class="hero-bottom">
			<p class="hero-desc rise" style="animation-delay:.3s">
				Pusat pelatihan dan konsultasi berkelas internasional di bidang energi, teknologi berwawasan
				lingkungan, dan manajemen — untuk individu maupun organisasi.
			</p>
			<div class="hero-cta rise" style="animation-delay:.4s">
				<a class="btn btn-primary" href="/kelas">Lihat Program</a>
				<a class="btn btn-ghost" href="/konsultasi">Ajukan Konsultasi</a>
			</div>
		</div>
	</div>

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
		min-height: 92vh;
		display: flex;
		flex-direction: column;
		justify-content: center;
		overflow: hidden;
	}
	.blob {
		position: absolute;
		right: -180px;
		top: 44%;
		transform: translateY(-50%);
		width: 780px;
		height: 780px;
		border-radius: 50%;
		background: radial-gradient(
			circle at 32% 30%,
			var(--sky-blue),
			var(--navy-teal) 58%,
			transparent 72%
		);
		filter: blur(14px);
		opacity: 0.32;
		animation: floatblob 9s ease-in-out infinite;
		pointer-events: none;
	}
	@keyframes floatblob {
		0%,
		100% {
			transform: translateY(-50%) scale(1);
		}
		50% {
			transform: translateY(-54%) scale(1.06);
		}
	}
	.lines {
		position: absolute;
		inset: 0;
		pointer-events: none;
		opacity: 0.6;
		background-image:
			repeating-linear-gradient(
				to right,
				transparent 0,
				transparent calc(8.33% - 1px),
				var(--border) calc(8.33% - 1px),
				var(--border) 8.33%
			),
			repeating-linear-gradient(
				to bottom,
				transparent 0,
				transparent calc(25% - 1px),
				var(--border) calc(25% - 1px),
				var(--border) 25%
			);
		mask-image: radial-gradient(ellipse 80% 70% at 40% 45%, black 40%, transparent 100%);
	}
	.hero-inner {
		position: relative;
		z-index: 2;
		padding: 140px 24px 60px;
		width: 100%;
	}
	.hero-title {
		margin-top: 26px;
		font-size: clamp(2.8rem, 9vw, 7rem);
		font-weight: 800;
		line-height: 0.94;
	}
	.line {
		display: block;
	}
	.hl-line {
		margin-top: 0.08em;
	}
	.hl {
		position: relative;
		display: inline-block;
		background: linear-gradient(120deg, var(--navy-teal), var(--sky-blue));
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}
	.char {
		display: inline-block;
		white-space: pre;
		animation: char-in 0.6s cubic-bezier(0.22, 1, 0.36, 1) both;
	}
	.hero-bottom {
		display: grid;
		grid-template-columns: 1.1fr 0.9fr;
		gap: 40px;
		align-items: end;
		margin-top: 48px;
	}
	.hero-desc {
		font-size: 20px;
		color: var(--muted);
		max-width: 560px;
	}
	.hero-cta {
		display: flex;
		gap: 14px;
		flex-wrap: wrap;
	}
	.hero-cta .btn {
		height: 54px;
		padding: 0 30px;
		font-size: 16px;
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
		.hero-bottom {
			grid-template-columns: 1fr;
			gap: 28px;
		}
		.pilar-grid,
		.unggulan-grid {
			grid-template-columns: 1fr;
		}
		.cta {
			padding: 44px 28px;
		}
	}
</style>
