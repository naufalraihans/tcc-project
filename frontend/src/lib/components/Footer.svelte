<footer>
	<div class="container grid-foot">
		<div>
			<div class="brand">TCC ITPLN</div>
			<p class="desc">
				Training &amp; Consulting Center — unit pelatihan dan konsultasi di bawah Institut Teknologi
				PLN.
			</p>
		</div>

		<div>
			<h4>Navigasi</h4>
			<a href="/kelas">Kelas</a>
			<a href="/topik">Topik</a>
			<a href="/konsultasi">Konsultasi</a>
			<a href="/tentang">Tentang</a>
		</div>

		<div>
			<h4>Kontak</h4>
			<p>Graha YPK PLN, Jl. Lebak Bulus Tengah No.5, Cilandak, Jakarta Selatan 12430</p>
			<p>trainingcenter@itpln.ac.id</p>
		</div>
	</div>
	<div class="container copy">© {new Date().getFullYear()} Training &amp; Consulting Center ITPLN.</div>
</footer>

<style>
	footer {
		background: var(--charcoal);
		color: rgba(255, 255, 255, 0.72);
		padding: 64px 0 28px;
		margin-top: 40px;
	}
	.grid-foot {
		display: grid;
		grid-template-columns: 1.4fr 1fr 1.4fr;
		gap: 40px;
	}
	.brand {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 20px;
		color: #fff;
	}
	.desc {
		margin-top: 12px;
		max-width: 340px;
		font-size: 15px;
	}
	h4 {
		color: #fff;
		font-size: 15px;
		margin-bottom: 14px;
	}
	footer a {
		display: block;
		padding: 5px 0;
		font-size: 15px;
		color: rgba(255, 255, 255, 0.72);
	}
	footer a:hover {
		color: var(--sky-blue);
	}
	.copy {
		margin-top: 40px;
		padding-top: 24px;
		border-top: 1px solid rgba(255, 255, 255, 0.12);
		font-size: 14px;
	}
	@media (max-width: 720px) {
		.grid-foot {
			grid-template-columns: 1fr;
			gap: 28px;
		}
	}
</style>
