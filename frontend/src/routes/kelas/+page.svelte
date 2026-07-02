<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import KelasCard from '$lib/components/KelasCard.svelte';

	let { data } = $props();

	const formats = [
		['online', 'Online'],
		['offline', 'Offline'],
		['hybrid', 'Hybrid']
	];
	const hargas = [
		['gratis', 'Gratis'],
		['berbayar', 'Berbayar']
	];

	function toggle(key: string, val: string): string {
		const p = new URLSearchParams($page.url.searchParams);
		if (p.get(key) === val) p.delete(key);
		else p.set(key, val);
		p.delete('page');
		return '?' + p.toString();
	}

	function setTopik(val: string): string {
		const p = new URLSearchParams($page.url.searchParams);
		if (val) p.set('topik', val);
		else p.delete('topik');
		p.delete('page');
		return '?' + p.toString();
	}

	function pageHref(n: number): string {
		const p = new URLSearchParams($page.url.searchParams);
		p.set('page', String(n));
		return '?' + p.toString();
	}

	function isActive(key: string, val: string): boolean {
		return $page.url.searchParams.get(key) === val;
	}
</script>

<svelte:head><title>Kelas — TCC ITPLN</title></svelte:head>

<section class="section">
	<div class="container">
		<span class="section-tag">Program Pelatihan</span>
		<h1 class="section-title">Daftar Kelas</h1>
		<p class="section-lead">
			Pilih program pelatihan sesuai bidang, format, dan kebutuhan kompetensi Anda.
		</p>

		<div class="filters">
			<div class="chip-row">
				{#each formats as [val, label]}
					<a class="chip" class:active={isActive('format', val)} href={toggle('format', val)}>
						{label}
					</a>
				{/each}
				{#each hargas as [val, label]}
					<a class="chip" class:active={isActive('harga', val)} href={toggle('harga', val)}>
						{label}
					</a>
				{/each}
			</div>
			<select
				aria-label="Filter topik"
				onchange={(e) => goto(setTopik(e.currentTarget.value))}
			>
				<option value="">Semua Topik</option>
				{#each data.topik as t}
					<option value={t.slug} selected={isActive('topik', t.slug)}>{t.nama}</option>
				{/each}
			</select>
		</div>

		{#if data.error}
			<div class="empty">Gagal memuat kelas: {data.error}</div>
		{:else if data.kelas && data.kelas.items.length > 0}
			<div class="grid list">
				{#each data.kelas.items as k (k.id)}
					<KelasCard kelas={k} />
				{/each}
			</div>

			{#if data.kelas.pagination.total_pages > 1}
				<div class="pager">
					{#if data.kelas.pagination.page > 1}
						<a class="btn btn-ghost" href={pageHref(data.kelas.pagination.page - 1)}>Sebelumnya</a>
					{/if}
					<span class="pageinfo">
						Halaman {data.kelas.pagination.page} dari {data.kelas.pagination.total_pages}
					</span>
					{#if data.kelas.pagination.page < data.kelas.pagination.total_pages}
						<a class="btn btn-ghost" href={pageHref(data.kelas.pagination.page + 1)}>Berikutnya</a>
					{/if}
				</div>
			{/if}
		{:else}
			<div class="empty">Belum ada kelas yang sesuai filter.</div>
		{/if}
	</div>
</section>

<style>
	.filters {
		display: flex;
		justify-content: space-between;
		align-items: center;
		gap: 16px;
		margin: 32px 0;
		flex-wrap: wrap;
	}
	.chip-row {
		display: flex;
		gap: 10px;
		flex-wrap: wrap;
	}
	.chip {
		padding: 8px 16px;
		border-radius: var(--radius-full);
		border: 1px solid var(--border);
		font-size: 14px;
		font-weight: 500;
		color: var(--cool-slate);
		cursor: pointer;
	}
	.chip.active {
		background: var(--navy-teal);
		color: #fff;
		border-color: var(--navy-teal);
	}
	select {
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		border: 1px solid var(--border);
		font-family: 'Inter', sans-serif;
		font-size: 14px;
		color: var(--cool-slate);
		background: #fff;
	}
	.list {
		grid-template-columns: repeat(3, 1fr);
	}
	.empty {
		padding: 64px 24px;
		text-align: center;
		color: var(--muted);
		background: var(--off-white);
		border-radius: var(--radius);
	}
	.pager {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 20px;
		margin-top: 40px;
	}
	.pageinfo {
		font-size: 14px;
		color: var(--muted);
	}
	@media (max-width: 900px) {
		.list {
			grid-template-columns: 1fr;
		}
	}
</style>
