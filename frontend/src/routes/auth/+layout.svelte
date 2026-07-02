<script lang="ts">
	let { children } = $props();
</script>

<div class="auth-bg">
	<div class="auth-inner">
		<a class="brand" href="/">TCC ITPLN</a>
		{@render children()}
	</div>
</div>

<style>
	.auth-bg {
		min-height: 100vh;
		display: grid;
		place-items: center;
		padding: 24px;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
	}
	.auth-inner {
		width: 100%;
		max-width: 420px;
	}
	.brand {
		display: block;
		text-align: center;
		color: #fff;
		font-family: 'Plus Jakarta Sans', sans-serif;
		font-weight: 800;
		font-size: 24px;
		margin-bottom: 20px;
	}
	:global(.auth-bg .card) {
		padding: 32px;
	}
	:global(.auth-bg h1) {
		font-size: 24px;
	}
	:global(.auth-bg .sub) {
		color: var(--muted);
		margin: 6px 0 22px;
		font-size: 15px;
	}
	:global(.auth-bg label) {
		display: flex;
		flex-direction: column;
		gap: 6px;
		margin-bottom: 16px;
	}
	:global(.auth-bg label span) {
		font-size: 14px;
		font-weight: 500;
		color: var(--cool-slate);
	}
	:global(.auth-bg input) {
		padding: 11px 14px;
		border: 1px solid var(--border);
		border-radius: var(--radius-sm);
		font-family: 'Inter', sans-serif;
		font-size: 15px;
	}
	:global(.auth-bg input:focus) {
		outline: none;
		border-color: var(--sky-blue);
	}
	:global(.auth-bg .block) {
		display: flex;
		justify-content: center;
		width: 100%;
		margin-top: 4px;
	}
	:global(.auth-bg .block[disabled]) {
		opacity: 0.6;
		cursor: not-allowed;
	}
	:global(.auth-bg .foot) {
		text-align: center;
		margin-top: 18px;
		font-size: 14px;
		color: var(--muted);
	}
	:global(.auth-bg .foot a),
	:global(.auth-bg .row a) {
		color: var(--sky-blue);
		font-weight: 600;
	}
	:global(.auth-bg .row) {
		text-align: right;
		font-size: 13px;
		margin-bottom: 16px;
	}
	:global(.auth-bg .alert) {
		background: rgba(220, 38, 38, 0.08);
		color: #dc2626;
		border: 1px solid rgba(220, 38, 38, 0.2);
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
		margin-bottom: 16px;
	}
	:global(.auth-bg .ok) {
		background: rgba(30, 123, 69, 0.08);
		color: var(--success);
		border: 1px solid rgba(30, 123, 69, 0.2);
		padding: 10px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
		margin-bottom: 16px;
	}
</style>
