<script lang="ts">
	import type { KelasListItem } from '$lib/types';
	import { harga, tanggal } from '$lib/format';

	let { kelas }: { kelas: KelasListItem } = $props();

	const sisa = $derived(kelas.kuota > 0 ? kelas.kuota - kelas.peserta_terdaftar : null);
</script>

<a class="card kelas-card" href="/kelas/{kelas.slug}">
	<div class="badges">
		<span class="badge">{kelas.format}</span>
		<span class="badge" class:free={kelas.tipe_harga === 'gratis'}>
			{harga(kelas.tipe_harga, kelas.harga)}
		</span>
	</div>
	<h3>{kelas.judul}</h3>
	{#if kelas.instruktur}
		<div class="inst">{kelas.instruktur.nama}</div>
	{/if}
	<div class="meta">
		<span>{tanggal(kelas.jadwal_mulai)}</span>
		{#if sisa !== null}
			<span>Sisa {sisa} kuota</span>
		{/if}
	</div>
</a>

<style>
	.kelas-card {
		display: flex;
		flex-direction: column;
	}
	.badges {
		display: flex;
		gap: 8px;
		margin-bottom: 16px;
	}
	.badge.free {
		background: rgba(30, 123, 69, 0.1);
		color: var(--success);
		border-color: rgba(30, 123, 69, 0.2);
	}
	.kelas-card h3 {
		font-size: 19px;
		flex: 1;
	}
	.inst {
		color: var(--cool-slate);
		font-size: 14px;
		margin-top: 8px;
	}
	.meta {
		display: flex;
		justify-content: space-between;
		gap: 12px;
		margin-top: 18px;
		padding-top: 16px;
		border-top: 1px solid var(--border);
		font-size: 13px;
		color: var(--muted);
	}
</style>
