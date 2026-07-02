<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { supabase } from '$lib/supabase';
	import { apiAuth } from '$lib/authApi';

	let { children } = $props();
	let ready = $state(false);

	const nav = [
		['/admin', 'Dashboard'],
		['/admin/kelas', 'Kelas'],
		['/admin/topik', 'Topik'],
		['/admin/instruktur', 'Instruktur'],
		['/admin/pendaftaran', 'Pendaftaran'],
		['/admin/konsultasi', 'Konsultasi'],
		['/admin/transaksi', 'Transaksi']
	];

	onMount(async () => {
		const { data } = await supabase.auth.getSession();
		if (!data.session) {
			goto('/auth/login?next=' + $page.url.pathname);
			return;
		}
		try {
			const me = await apiAuth<{ role: string }>('/auth/me');
			if (me.role !== 'admin') {
				goto('/dashboard');
				return;
			}
			ready = true;
		} catch {
			goto('/auth/login');
		}
	});

	async function logout() {
		await supabase.auth.signOut();
		goto('/');
	}
</script>

{#if ready}
	<div class="shell">
		<aside class="sidebar">
			<div class="brand">TCC Admin</div>
			<nav>
				{#each nav as [href, label]}
					<a {href} class:active={$page.url.pathname === href}>{label}</a>
				{/each}
			</nav>
			<button class="logout" onclick={logout}>Keluar</button>
		</aside>
		<div class="content">
			{@render children()}
		</div>
	</div>
{:else}
	<div class="loading">Memuat…</div>
{/if}

<style>
	.shell {
		display: grid;
		grid-template-columns: 240px 1fr;
		min-height: 100vh;
	}
	.sidebar {
		background: var(--charcoal);
		padding: 28px 18px;
		display: flex;
		flex-direction: column;
		gap: 6px;
		position: sticky;
		top: 0;
		height: 100vh;
	}
	.brand {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 20px;
		color: #fff;
		margin-bottom: 24px;
	}
	nav {
		display: flex;
		flex-direction: column;
		gap: 4px;
		flex: 1;
	}
	nav a {
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		color: rgba(255, 255, 255, 0.7);
		font-weight: 500;
		font-size: 15px;
	}
	nav a:hover {
		background: rgba(255, 255, 255, 0.06);
		color: #fff;
	}
	nav a.active {
		background: var(--navy-teal);
		color: #fff;
	}
	.logout {
		text-align: left;
		padding: 10px 14px;
		border: none;
		background: none;
		color: rgba(255, 255, 255, 0.7);
		font-family: 'Inter', sans-serif;
		font-size: 15px;
		font-weight: 500;
		cursor: pointer;
		border-radius: var(--radius-sm);
	}
	.logout:hover {
		background: rgba(255, 255, 255, 0.06);
		color: #fff;
	}
	.content {
		padding: 36px;
		background: var(--off-white);
		min-width: 0;
	}
	.loading {
		min-height: 100vh;
		display: grid;
		place-items: center;
		color: var(--muted);
	}

	:global(.content h1) {
		font-size: 26px;
		margin-bottom: 6px;
	}
	:global(.content .lead) {
		color: var(--muted);
		margin-bottom: 24px;
	}
	:global(.content .panel) {
		background: #fff;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 22px;
	}
	:global(.content .toolbar) {
		display: flex;
		justify-content: space-between;
		align-items: center;
		gap: 12px;
		margin-bottom: 18px;
		flex-wrap: wrap;
	}
	:global(.content table) {
		width: 100%;
		border-collapse: collapse;
		background: #fff;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		overflow: hidden;
	}
	:global(.content th) {
		text-align: left;
		font-size: 13px;
		color: var(--muted);
		padding: 12px 16px;
		border-bottom: 1px solid var(--border);
		background: var(--off-white);
	}
	:global(.content td) {
		padding: 12px 16px;
		border-bottom: 1px solid var(--border);
		font-size: 14px;
		vertical-align: middle;
	}
	:global(.content tr:last-child td) {
		border-bottom: none;
	}
	:global(.content .btn-sm) {
		padding: 7px 12px;
		font-size: 13px;
		border-radius: var(--radius-sm);
		border: 1px solid var(--border);
		background: #fff;
		color: var(--cool-slate);
		cursor: pointer;
		font-family: 'Inter', sans-serif;
		font-weight: 500;
	}
	:global(.content .btn-sm:hover) {
		border-color: var(--sky-blue);
		color: var(--sky-blue);
	}
	:global(.content .btn-sm.danger:hover) {
		border-color: #dc2626;
		color: #dc2626;
	}
	:global(.content .adm-form) {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		gap: 16px;
	}
	:global(.content .adm-form label) {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}
	:global(.content .adm-form label.full) {
		grid-column: 1 / -1;
	}
	:global(.content .adm-form span) {
		font-size: 13px;
		font-weight: 500;
		color: var(--cool-slate);
	}
	:global(.content .adm-form input),
	:global(.content .adm-form select),
	:global(.content .adm-form textarea) {
		padding: 10px 12px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 14px;
		background: #fff;
	}
	:global(.content .adm-form input:focus),
	:global(.content .adm-form select:focus),
	:global(.content .adm-form textarea:focus) {
		outline: none;
		border-color: var(--sky-blue);
	}
	:global(.content .alert) {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
		border: 1px solid rgba(220, 38, 38, 0.2);
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
		margin-bottom: 16px;
	}
	@media (max-width: 820px) {
		.shell {
			grid-template-columns: 1fr;
		}
		.sidebar {
			position: static;
			height: auto;
		}
		:global(.content .adm-form) {
			grid-template-columns: 1fr;
		}
	}
</style>
