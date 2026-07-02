<script lang="ts">
	import { page } from '$app/stores';
</script>

<section class="section container err">
	<div class="code">{$page.status}</div>
	<h1>Halaman tidak dapat ditampilkan</h1>
	<p>{$page.error?.message}</p>
	<a class="btn btn-primary" href="/">Kembali ke Beranda</a>
</section>

<style>
	.err {
		text-align: center;
		padding: 120px 24px;
	}
	.code {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 72px;
		color: var(--sky-blue);
	}
	.err h1 {
		margin: 8px 0 12px;
	}
	.err p {
		color: var(--muted);
		margin-bottom: 24px;
	}
</style>
