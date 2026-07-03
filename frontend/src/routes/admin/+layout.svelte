<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { apiAuth, currentToken, clearToken } from '$lib/authApi';
	import {
		LayoutDashboard,
		GraduationCap,
		Tags,
		Users,
		ClipboardList,
		MessagesSquare,
		Receipt,
		Target,
		Megaphone,
		LogOut,
		Menu
	} from 'lucide-svelte';
	import type { Profile } from '$lib/types';

	let { children } = $props();
	let ready = $state(false);
	let mobileOpen = $state(false);
	let profile = $state<Profile | null>(null);

	const name = $derived(profile?.full_name || 'Admin');
	const email = $derived(profile?.email ?? '');

	const nav = [
		{ href: '/admin', label: 'Dashboard', icon: LayoutDashboard },
		{ href: '/admin/kelas', label: 'Kelas', icon: GraduationCap },
		{ href: '/admin/topik', label: 'Topik', icon: Tags },
		{ href: '/admin/instruktur', label: 'Instruktur', icon: Users },
		{ href: '/admin/pendaftaran', label: 'Pendaftaran', icon: ClipboardList },
		{ href: '/admin/konsultasi', label: 'Konsultasi', icon: MessagesSquare },
		{ href: '/admin/transaksi', label: 'Transaksi', icon: Receipt },
		{ href: '/admin/misi', label: 'Misi', icon: Target },
		{ href: '/admin/pengumuman', label: 'Pengumuman', icon: Megaphone }
	];

	onMount(async () => {
		if (!currentToken()) {
			goto('/auth/login?next=' + $page.url.pathname);
			return;
		}
		try {
			profile = await apiAuth<Profile>('/auth/me');
			if (profile.role !== 'admin') {
				goto('/dashboard');
				return;
			}
			ready = true;
		} catch {
			goto('/auth/login');
		}
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
			<a class="brand" href="/">
				<img src="/logo-tcc.png" alt="TCC ITPLN" />
				<span class="brand-tag">Panel Admin</span>
			</a>

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
				<div class="avatar">{name[0]?.toUpperCase() ?? 'A'}</div>
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
				<span class="mobile-tag">Admin</span>
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
		height: 36px;
		width: auto;
		display: block;
	}
	.brand-tag {
		display: inline-block;
		margin-top: 8px;
		font-size: 11px;
		font-weight: 600;
		letter-spacing: 0.06em;
		text-transform: uppercase;
		color: var(--sky-blue);
		background: rgba(26, 141, 178, 0.1);
		padding: 3px 9px;
		border-radius: 999px;
	}
	nav {
		display: flex;
		flex-direction: column;
		gap: 4px;
		overflow-y: auto;
	}
	.nav-item {
		display: flex;
		align-items: center;
		gap: 13px;
		padding: 10px 14px;
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
	.mobile-tag {
		font-size: 11px;
		font-weight: 600;
		color: var(--sky-blue);
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
		padding: 32px;
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
		margin-bottom: 24px;
		font-size: 15px;
	}
	:global(.content .panel) {
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: 16px;
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
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: 16px;
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
		border-radius: 10px;
		border: 1px solid var(--border);
		background: var(--white);
		color: var(--cool-slate);
		cursor: pointer;
		font-family: var(--font-sans);
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
		border-radius: 10px;
		font-family: var(--font-sans);
		font-size: 14px;
		background: var(--white);
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
		border-radius: 10px;
		font-size: 14px;
		margin-bottom: 16px;
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
		:global(.content .adm-form) {
			grid-template-columns: 1fr;
		}
	}
</style>
