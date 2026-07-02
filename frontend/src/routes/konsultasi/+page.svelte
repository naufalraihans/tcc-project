<script lang="ts">
	import { goto } from '$app/navigation';

	function handleSubmit(e: SubmitEvent) {
		e.preventDefault();
		goto('/auth/login?next=/konsultasi');
	}
</script>

<svelte:head><title>Konsultasi — TCC ITPLN</title></svelte:head>

<section class="section">
	<div class="container split">
		<div class="intro">
			<span class="section-tag">Layanan Konsultasi</span>
			<h1 class="section-title">Ajukan Konsultasi</h1>
			<p class="section-lead">
				Tim TCC ITPLN siap membantu merancang program pelatihan dan pendampingan sesuai kebutuhan
				kompetensi organisasi Anda — mulai dari asesmen, kurikulum, hingga sertifikasi.
			</p>
			<ul class="points">
				<li>Program tailor-made sesuai kebutuhan klien</li>
				<li>Instruktur dan konsultan berpengalaman di bidang ketenagalistrikan</li>
				<li>Pendampingan menuju sertifikasi dan peningkatan kinerja</li>
			</ul>
		</div>

		<form class="card form" onsubmit={handleSubmit}>
			<label>
				<span>Nama Pengirim</span>
				<input type="text" name="nama_pengirim" required />
			</label>
			<label>
				<span>Topik Konsultasi</span>
				<input type="text" name="topik_konsultasi" required />
			</label>
			<label>
				<span>Pesan</span>
				<textarea name="pesan" rows="5" required></textarea>
			</label>
			<label>
				<span>Kontak (Email / WhatsApp)</span>
				<input type="text" name="kontak" required />
			</label>
			<button class="btn btn-primary block" type="submit">Kirim Pengajuan</button>
			<p class="note">Anda perlu masuk terlebih dahulu untuk mengirim pengajuan.</p>
		</form>
	</div>
</section>

<style>
	.split {
		display: grid;
		grid-template-columns: 1fr 0.9fr;
		gap: 48px;
		align-items: start;
	}
	.points {
		list-style: none;
		margin-top: 24px;
		display: flex;
		flex-direction: column;
		gap: 12px;
	}
	.points li {
		padding-left: 20px;
		position: relative;
		color: var(--cool-slate);
	}
	.points li::before {
		content: '';
		position: absolute;
		left: 0;
		top: 10px;
		width: 8px;
		height: 8px;
		border-radius: 50%;
		background: var(--sky-blue);
	}
	.form {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}
	label {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}
	label span {
		font-size: 14px;
		font-weight: 500;
		color: var(--cool-slate);
	}
	input,
	textarea {
		padding: 11px 14px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 15px;
		color: var(--charcoal);
		resize: vertical;
	}
	input:focus,
	textarea:focus {
		outline: none;
		border-color: var(--sky-blue);
	}
	.block {
		justify-content: center;
		width: 100%;
	}
	.note {
		font-size: 13px;
		color: var(--muted);
		text-align: center;
	}
	@media (max-width: 900px) {
		.split {
			grid-template-columns: 1fr;
		}
	}
</style>
