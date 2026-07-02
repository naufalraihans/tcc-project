<script lang="ts">
	let open = $state(false);
	const links = [
		{ href: '/kelas', label: 'Kelas' },
		{ href: '/topik', label: 'Topik' },
		{ href: '/konsultasi', label: 'Konsultasi' },
		{ href: '/tentang', label: 'Tentang' }
	];
</script>

<header>
	<div class="container nav">
		<a class="logo" href="/">
			<span class="mark">TCC</span>
			<span class="divider"></span>
			<span class="sub">Training &amp; Consulting Center</span>
		</a>

		<nav class="links">
			{#each links as l}
				<a href={l.href}>{l.label}</a>
			{/each}
		</nav>

		<div class="actions">
			<a class="btn btn-ghost" href="/auth/login">Masuk</a>
			<a class="btn btn-primary" href="/auth/register">Daftar</a>
		</div>

		<button class="toggle" aria-expanded={open} onclick={() => (open = !open)}>Menu</button>
	</div>

	{#if open}
		<div class="drawer container">
			{#each links as l}
				<a href={l.href} onclick={() => (open = false)}>{l.label}</a>
			{/each}
			<a class="btn btn-ghost" href="/auth/login">Masuk</a>
			<a class="btn btn-primary" href="/auth/register">Daftar</a>
		</div>
	{/if}
</header>

<style>
	header {
		position: sticky;
		top: 0;
		z-index: 50;
		background: rgba(255, 255, 255, 0.85);
		backdrop-filter: blur(12px);
		border-bottom: 1px solid var(--border);
	}
	.nav {
		display: flex;
		align-items: center;
		justify-content: space-between;
		height: 68px;
	}
	.logo {
		display: flex;
		align-items: center;
		gap: 10px;
	}
	.mark {
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 20px;
		color: var(--navy-teal);
		letter-spacing: -0.03em;
	}
	.divider {
		width: 1px;
		height: 22px;
		background: var(--border);
	}
	.sub {
		font-size: 13px;
		font-weight: 600;
		color: var(--muted);
	}
	.links {
		display: flex;
		gap: 32px;
	}
	.links a {
		font-size: 15px;
		font-weight: 500;
		color: var(--cool-slate);
	}
	.links a:hover {
		color: var(--sky-blue);
	}
	.actions {
		display: flex;
		gap: 12px;
	}
	.toggle {
		display: none;
		background: none;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		padding: 8px 14px;
		font-weight: 600;
		font-family: 'Inter', sans-serif;
		color: var(--navy-teal);
		cursor: pointer;
	}
	.drawer {
		display: none;
	}
	@media (max-width: 820px) {
		.links,
		.actions {
			display: none;
		}
		.toggle {
			display: inline-flex;
		}
		.drawer {
			display: flex;
			flex-direction: column;
			gap: 12px;
			padding-bottom: 20px;
		}
		.drawer a:not(.btn) {
			padding: 8px 0;
			font-weight: 500;
			color: var(--cool-slate);
			border-bottom: 1px solid var(--border);
		}
	}
</style>
