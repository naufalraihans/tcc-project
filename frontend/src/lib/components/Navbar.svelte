<script lang="ts">
	import { onMount } from 'svelte';

	let scrolled = $state(false);
	let open = $state(false);

	const links = [
		{ href: '/kelas', label: 'Kelas' },
		{ href: '/topik', label: 'Topik' },
		{ href: '/konsultasi', label: 'Konsultasi' },
		{ href: '/tentang', label: 'Tentang' }
	];

	onMount(() => {
		const onScroll = () => (scrolled = window.scrollY > 20);
		onScroll();
		window.addEventListener('scroll', onScroll, { passive: true });
		return () => window.removeEventListener('scroll', onScroll);
	});
</script>

<header class:scrolled>
	<nav class:pill={scrolled}>
		<a class="logo" href="/">
			<img src="/logo-tcc.png" alt="TCC ITPLN" />
		</a>

		<div class="links">
			{#each links as l}
				<a href={l.href}>{l.label}</a>
			{/each}
		</div>

		<div class="actions">
			<a class="signin" href="/auth/login">Masuk</a>
			<a class="btn btn-primary" href="/auth/register">Daftar</a>
		</div>

		<button class="toggle" aria-expanded={open} onclick={() => (open = !open)}>
			{open ? 'Tutup' : 'Menu'}
		</button>
	</nav>

	{#if open}
		<div class="drawer">
			{#each links as l}
				<a href={l.href} onclick={() => (open = false)}>{l.label}</a>
			{/each}
			<div class="drawer-actions">
				<a class="btn btn-ghost" href="/auth/login">Masuk</a>
				<a class="btn btn-primary" href="/auth/register">Daftar</a>
			</div>
		</div>
	{/if}
</header>

<style>
	header {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		z-index: 50;
		display: flex;
		justify-content: center;
		padding: 0;
		transition: padding 0.4s ease;
	}
	header.scrolled {
		padding: 14px 16px;
	}
	nav {
		width: 100%;
		max-width: 1400px;
		display: flex;
		align-items: center;
		justify-content: space-between;
		height: 80px;
		padding: 0 32px;
		transition: all 0.4s ease;
		border: 1px solid transparent;
	}
	nav.pill {
		max-width: 1180px;
		height: 60px;
		background: rgba(255, 255, 255, 0.8);
		backdrop-filter: blur(16px);
		border-color: var(--border);
		border-radius: var(--radius);
		box-shadow: 0 12px 40px rgba(16, 21, 27, 0.08);
	}
	.logo {
		display: flex;
		align-items: center;
	}
	.logo img {
		height: 40px;
		width: auto;
		display: block;
		transition: height 0.4s ease;
	}
	nav.pill .logo img {
		height: 32px;
	}
	.links {
		display: flex;
		gap: 40px;
	}
	.links a {
		font-size: 15px;
		color: var(--cool-slate);
		position: relative;
	}
	.links a::after {
		content: '';
		position: absolute;
		left: 0;
		bottom: -6px;
		height: 1.5px;
		width: 0;
		background: var(--ink);
		transition: width 0.3s ease;
	}
	.links a:hover {
		color: var(--ink);
	}
	.links a:hover::after {
		width: 100%;
	}
	.actions {
		display: flex;
		align-items: center;
		gap: 20px;
	}
	.signin {
		font-size: 15px;
		color: var(--cool-slate);
	}
	.signin:hover {
		color: var(--ink);
	}
	.toggle {
		display: none;
		background: none;
		border: 1px solid var(--border-strong);
		border-radius: var(--radius-pill);
		padding: 8px 16px;
		font-weight: 600;
		font-family: var(--font-sans);
		color: var(--ink);
		cursor: pointer;
	}
	.drawer {
		position: absolute;
		top: 88px;
		left: 16px;
		right: 16px;
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: var(--radius);
		box-shadow: 0 20px 50px rgba(16, 21, 27, 0.12);
		padding: 20px;
		display: flex;
		flex-direction: column;
		gap: 6px;
	}
	.drawer a:not(.btn) {
		padding: 12px;
		font-weight: 500;
		color: var(--cool-slate);
		border-radius: var(--radius-sm);
	}
	.drawer a:not(.btn):hover {
		background: var(--off-white);
	}
	.drawer-actions {
		display: flex;
		gap: 10px;
		margin-top: 10px;
	}
	.drawer-actions .btn {
		flex: 1;
	}
	@media (max-width: 860px) {
		nav {
			height: 68px;
			padding: 0 20px;
		}
		.links,
		.actions {
			display: none;
		}
		.toggle {
			display: inline-flex;
		}
	}
</style>
