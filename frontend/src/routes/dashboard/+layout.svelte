<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { supabase } from '$lib/supabase';

	let { children } = $props();
	let ready = $state(false);

	const nav = [
		['/dashboard', 'Dashboard'],
		['/dashboard/kelas', 'Kelas Saya'],
		['/dashboard/konsultasi', 'Konsultasi'],
		['/dashboard/profil', 'Profil']
	];

	onMount(async () => {
		const { data } = await supabase.auth.getSession();
		if (!data.session) {
			goto('/auth/login?next=' + $page.url.pathname);
			return;
		}
		ready = true;
	});

	async function logout() {
		await supabase.auth.signOut();
		goto('/');
	}
</script>

{#if ready}
	<div class="shell">
		<aside class="sidebar">
			<a class="brand" href="/">TCC ITPLN</a>
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
		grid-template-columns: 260px 1fr;
		min-height: 100vh;
	}
	.sidebar {
		background: var(--charcoal);
		padding: 28px 20px;
		display: flex;
		flex-direction: column;
		gap: 8px;
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
		padding: 11px 14px;
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
		padding: 11px 14px;
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
		padding: 40px;
		background: var(--off-white);
	}
	.loading {
		min-height: 100vh;
		display: grid;
		place-items: center;
		color: var(--muted);
	}
	:global(.content h1) {
		font-size: 28px;
		margin-bottom: 6px;
	}
	:global(.content .lead) {
		color: var(--muted);
		margin-bottom: 28px;
	}
	:global(.content .panel) {
		background: #fff;
		border: 1px solid var(--border);
		border-radius: var(--radius);
		padding: 24px;
	}
	@media (max-width: 820px) {
		.shell {
			grid-template-columns: 1fr;
		}
		.sidebar {
			position: static;
			height: auto;
			flex-direction: row;
			flex-wrap: wrap;
			align-items: center;
		}
		nav {
			flex-direction: row;
			flex-wrap: wrap;
			flex: 1 1 100%;
		}
		.content {
			padding: 24px;
		}
	}
</style>
