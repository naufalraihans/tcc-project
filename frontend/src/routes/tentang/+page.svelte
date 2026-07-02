<script lang="ts">
	const struktur = [
		['Direktur Pusat Pelatihan & Konsultansi (TCC)', 'Ir. Suharto, M.T., IPU'],
		['MB Perencanaan & Pemasaran TCC', 'Ir. Dwi Listiawati, MT., IPM., ASEAN.Eng'],
		['MB Operasi TCC', 'Anna Agustina'],
		['Language Development Center', 'Dr. Widya Nita Suliyanti, ST., M.Comp.Sc']
	];

	const kontak = [
		['Kantor', 'Graha YPK PLN, Jl. Lebak Bulus Tengah No.5, Cilandak, Jakarta Selatan 12430'],
		['Kampus ITPLN', 'Jl. Lingkar Luar Barat, Duri Kosambi, Cengkareng, Jakarta Barat 11750'],
		['Email', 'trainingcenter@itpln.ac.id'],
		['Telepon', '+62 813-1084-5077 (Anna Agustina)']
	];
</script>

<svelte:head><title>Tentang — TCC ITPLN</title></svelte:head>

<section class="section">
	<div class="container narrow">
		<span class="section-tag">Tentang Kami</span>
		<h1 class="section-title">Training &amp; Consulting Center ITPLN</h1>
		<p class="section-lead">
			Unit bisnis di bawah Institut Teknologi PLN yang menyediakan layanan jasa pelatihan dan
			konsultasi di bidang energi, teknologi, dan manajemen berwawasan lingkungan. Berdiri sejak 7
			Februari 2023, dan bertransformasi menjadi Training &amp; Consulting Center pada Mei 2026.
		</p>

		<h2 class="sub">Visi</h2>
		<p class="body">
			Menyelenggarakan pelatihan berkelas internasional yang unggul dan profesional di bidang energi,
			teknologi berwawasan lingkungan, serta manajemen.
		</p>

		<h2 class="sub">Struktur Organisasi</h2>
		<div class="rows">
			{#each struktur as [jabatan, nama]}
				<div class="row">
					<div class="label">{jabatan}</div>
					<div class="value">{nama}</div>
				</div>
			{/each}
		</div>

		<h2 class="sub">Kontak &amp; Alamat</h2>
		<div class="rows">
			{#each kontak as [label, value]}
				<div class="row">
					<div class="label">{label}</div>
					<div class="value">{value}</div>
				</div>
			{/each}
		</div>
	</div>
</section>

<style>
	.narrow {
		max-width: 800px;
	}
	.sub {
		font-size: 22px;
		margin: 40px 0 14px;
	}
	.body {
		color: var(--cool-slate);
	}
	.rows {
		display: flex;
		flex-direction: column;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		overflow: hidden;
	}
	.row {
		display: grid;
		grid-template-columns: 1fr 1.4fr;
		gap: 16px;
		padding: 16px 20px;
		border-bottom: 1px solid var(--border);
	}
	.row:last-child {
		border-bottom: none;
	}
	.label {
		color: var(--muted);
		font-size: 14px;
	}
	.value {
		font-weight: 600;
	}
	@media (max-width: 640px) {
		.row {
			grid-template-columns: 1fr;
			gap: 4px;
		}
	}
</style>
