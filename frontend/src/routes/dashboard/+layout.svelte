<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { currentToken, clearToken, apiAuth } from '$lib/authApi';
	import {
		LayoutDashboard,
		GraduationCap,
		MessagesSquare,
		User,
		LogOut,
		Menu,
		Compass
	} from 'lucide-svelte';
	import type { Profile } from '$lib/types';

	let { children } = $props();
	let ready = $state(false);
	let mobileOpen = $state(false);
	let profile = $state<Profile | null>(null);

	const name = $derived(profile?.full_name || 'Pengguna');
	const email = $derived(profile?.email ?? '');

	const nav = [
		{ href: '/dashboard', label: 'Beranda', icon: LayoutDashboard },
		{ href: '/dashboard/kelas', label: 'Kelas Saya', icon: GraduationCap },
		{ href: '/dashboard/konsultasi', label: 'Konsultasi', icon: MessagesSquare },
		{ href: '/dashboard/jelajah', label: 'Jelajahi Kelas', icon: Compass },
		{ href: '/dashboard/profil', label: 'Profil', icon: User }
	];

	onMount(async () => {
		if (!currentToken()) {
			goto('/auth/login?next=' + $page.url.pathname);
			return;
		}
		try {
			profile = await apiAuth<Profile>('/auth/me');
		} catch {
			/* keep defaults */
		}
		ready = true;
	});

	function logout() {
		clearToken();
		goto('/');
	}
</script>

{#if ready}
	<div class="shell">
		{#if mobileOpen}
			<button class="scrim" aria-label="Tutup menu" onclick={() => (mobileOpen = false)}></button>
		{/if}

		<aside class="sidebar" class:open={mobileOpen}>
			<a class="brand" href="/"><img src="/logo-tcc.png" alt="TCC ITPLN" /></a>

			<nav>
				{#each nav as n (n.href)}
					{@const Icon = n.icon}
					<a
						href={n.href}
						class="nav-item"
						class:active={$page.url.pathname === n.href}
						onclick={() => (mobileOpen = false)}
					>
						<Icon size={19} strokeWidth={2} />
						<span>{n.label}</span>
					</a>
				{/each}
			</nav>

			<div class="user-card">
				<div class="avatar">{name[0]?.toUpperCase() ?? 'U'}</div>
				<div class="user-meta">
					<div class="user-name">{name}</div>
					<div class="user-mail">{email}</div>
				</div>
				<button class="logout" aria-label="Keluar" onclick={logout}>
					<LogOut size={17} strokeWidth={2} />
				</button>
			</div>
		</aside>

		<div class="main">
			<header class="mobilebar">
				<button class="menu-btn" aria-label="Menu" onclick={() => (mobileOpen = !mobileOpen)}>
					<Menu size={22} />
				</button>
				<img class="mobile-logo" src="/logo-tcc.png" alt="TCC ITPLN" />
			</header>
			<main class="content">
				{@render children()}
			</main>
		</div>
	</div>
{:else}
	<div class="loading">Memuat…</div>
{/if}

<style>
	.shell {
		display: grid;
		grid-template-columns: 256px 1fr;
		min-height: 100vh;
		background: var(--off-white);
	}
	.sidebar {
		background: var(--white);
		border-right: 1px solid var(--border);
		padding: 22px 16px 16px;
		display: flex;
		flex-direction: column;
		gap: 4px;
		position: sticky;
		top: 0;
		height: 100vh;
	}
	.brand {
		display: block;
		padding: 4px 8px 24px;
	}
	.brand img {
		height: 38px;
		width: auto;
		display: block;
	}
	nav {
		display: flex;
		flex-direction: column;
		gap: 4px;
	}
	.nav-item {
		display: flex;
		align-items: center;
		gap: 13px;
		padding: 11px 14px;
		border-radius: 12px;
		color: var(--cool-slate);
		font-size: 14.5px;
		font-weight: 500;
		font-family: var(--font-sans);
		transition:
			background 0.18s ease,
			color 0.18s ease;
	}
	.nav-item:hover {
		background: var(--off-white);
		color: var(--ink);
	}
	.nav-item.active {
		background: rgba(26, 141, 178, 0.12);
		color: var(--navy-teal);
		font-weight: 600;
	}
	.user-card {
		margin-top: auto;
		display: flex;
		align-items: center;
		gap: 11px;
		padding: 10px;
		border-radius: 14px;
		border: 1px solid var(--border);
		background: var(--off-white);
	}
	.avatar {
		width: 38px;
		height: 38px;
		border-radius: 50%;
		flex-shrink: 0;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
		display: grid;
		place-items: center;
		font-weight: 700;
		font-size: 15px;
	}
	.user-meta {
		min-width: 0;
		flex: 1;
	}
	.user-name {
		font-size: 13.5px;
		font-weight: 600;
		color: var(--ink);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}
	.user-mail {
		font-size: 12px;
		color: var(--muted);
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}
	.logout {
		flex-shrink: 0;
		border: none;
		background: none;
		cursor: pointer;
		color: var(--muted);
		padding: 7px;
		border-radius: 9px;
		display: grid;
		place-items: center;
		transition:
			background 0.18s ease,
			color 0.18s ease;
	}
	.logout:hover {
		background: rgba(220, 38, 38, 0.1);
		color: #dc2626;
	}
	.main {
		display: flex;
		flex-direction: column;
		min-width: 0;
	}
	.mobilebar {
		display: none;
		align-items: center;
		gap: 14px;
		height: 60px;
		padding: 0 18px;
		border-bottom: 1px solid var(--border);
		background: var(--white);
		position: sticky;
		top: 0;
		z-index: 20;
	}
	.mobile-logo {
		height: 30px;
	}
	.menu-btn {
		background: none;
		border: none;
		cursor: pointer;
		color: var(--cool-slate);
		padding: 4px;
		display: inline-flex;
	}
	.content {
		padding: 30px;
		flex: 1;
		min-width: 0;
	}
	.scrim {
		display: none;
	}
	.loading {
		min-height: 100vh;
		display: grid;
		place-items: center;
		color: var(--muted);
	}

	:global(.content h1) {
		font-size: 26px;
		font-weight: 800;
		color: var(--ink);
		margin-bottom: 6px;
		letter-spacing: -0.02em;
	}
	:global(.content .lead) {
		color: var(--muted);
		margin-bottom: 26px;
		font-size: 15px;
	}
	:global(.content .panel) {
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: 16px;
		padding: 22px;
	}

	@media (max-width: 900px) {
		.shell {
			grid-template-columns: 1fr;
		}
		.sidebar {
			position: fixed;
			z-index: 60;
			width: 264px;
			transform: translateX(-100%);
			transition: transform 0.28s cubic-bezier(0.22, 1, 0.36, 1);
		}
		.sidebar.open {
			transform: translateX(0);
			box-shadow: 0 24px 60px rgba(12, 79, 106, 0.22);
		}
		.scrim {
			display: block;
			position: fixed;
			inset: 0;
			z-index: 55;
			border: none;
			cursor: pointer;
			background: rgba(16, 21, 27, 0.4);
		}
		.mobilebar {
			display: flex;
		}
		.content {
			padding: 20px;
		}
	}
</style>
