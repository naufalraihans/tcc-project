<script lang="ts">
	import '../app.css';
	import { page } from '$app/stores';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';

	let { children } = $props();

	const bare = $derived(
		$page.url.pathname.startsWith('/auth') ||
			$page.url.pathname.startsWith('/dashboard') ||
			$page.url.pathname.startsWith('/admin')
	);
</script>

{#if bare}
	{@render children()}
{:else}
	<Navbar />
	<main class="noise">
		{@render children()}
	</main>
	<Footer />
{/if}
