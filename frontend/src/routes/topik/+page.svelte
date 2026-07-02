<script lang="ts">
	let { data } = $props();
</script>

<svelte:head><title>Topik Pelatihan — TCC ITPLN</title></svelte:head>

<section class="section">
	<div class="container">
		<span class="section-tag">Bidang Pelatihan</span>
		<h1 class="section-title">Topik Pelatihan</h1>
		<p class="section-lead">Jelajahi kelas berdasarkan bidang kompetensi.</p>

		{#if data.error}
			<div class="empty">Gagal memuat topik: {data.error}</div>
		{:else if data.topik.length > 0}
			<div class="grid list">
				{#each data.topik as t (t.id)}
					<a class="card topik" href="/kelas?topik={t.slug}">
						<h3>{t.nama}</h3>
						{#if t.deskripsi}<p class="desc">{t.deskripsi}</p>{/if}
						<div class="count">{t.jumlah_kelas ?? 0} kelas</div>
					</a>
				{/each}
			</div>
		{:else}
			<div class="empty">Belum ada topik.</div>
		{/if}
	</div>
</section>

<style>
	.list {
		grid-template-columns: repeat(3, 1fr);
		margin-top: 32px;
	}
	.topik {
		display: flex;
		flex-direction: column;
	}
	.topik h3 {
		font-size: 20px;
	}
	.desc {
		color: var(--muted);
		font-size: 14px;
		margin-top: 10px;
		flex: 1;
	}
	.count {
		margin-top: 18px;
		font-weight: 600;
		font-size: 14px;
		color: var(--sky-blue);
	}
	.empty {
		padding: 64px 24px;
		text-align: center;
		color: var(--muted);
		background: var(--off-white);
		border-radius: var(--radius);
		margin-top: 32px;
	}
	@media (max-width: 900px) {
		.list {
			grid-template-columns: 1fr;
		}
	}
</style>
