<script lang="ts">
	let { children } = $props();
</script>

<div class="auth-bg">
	<div class="orbs" aria-hidden="true">
		<span class="orb o1"></span>
		<span class="orb o2"></span>
		<span class="orb o3"></span>
	</div>
	<div class="auth-inner">
		<a class="brand" href="/">TCC ITPLN</a>
		{@render children()}
	</div>
</div>

<style>
	.auth-bg {
		position: relative;
		min-height: 100vh;
		display: grid;
		place-items: center;
		padding: 24px;
		overflow: hidden;
		background:
			radial-gradient(1200px 600px at 20% 10%, var(--sky-blue), transparent 60%),
			linear-gradient(135deg, var(--ink), var(--navy-teal) 70%, var(--sky-blue));
	}
	.orbs {
		position: absolute;
		inset: 0;
		pointer-events: none;
	}
	.orb {
		position: absolute;
		border-radius: 50%;
		background: rgba(255, 255, 255, 0.14);
		backdrop-filter: blur(20px) saturate(180%);
		border: 2px solid rgba(255, 255, 255, 0.3);
		box-shadow: 0 8px 32px rgba(255, 255, 255, 0.18);
		animation: pulse 5s ease-in-out infinite;
	}
	.o1 {
		top: 20%;
		left: 18%;
		width: 130px;
		height: 130px;
	}
	.o2 {
		bottom: 18%;
		right: 20%;
		width: 100px;
		height: 100px;
		animation-delay: 1.2s;
	}
	.o3 {
		top: 52%;
		right: 30%;
		width: 68px;
		height: 68px;
		animation-delay: 0.6s;
	}
	@keyframes pulse {
		0%,
		100% {
			opacity: 0.45;
			transform: translateY(0) scale(1);
		}
		50% {
			opacity: 0.65;
			transform: translateY(-12px) scale(1.05);
		}
	}
	.auth-inner {
		position: relative;
		z-index: 2;
		width: 100%;
		max-width: 430px;
	}
	.brand {
		display: block;
		text-align: center;
		color: #fff;
		font-family: var(--font-display);
		font-weight: 800;
		font-size: 24px;
		letter-spacing: -0.03em;
		margin-bottom: 22px;
	}

	:global(.auth-bg .card) {
		padding: 36px;
		background: rgba(255, 255, 255, 0.16);
		backdrop-filter: blur(40px) saturate(200%);
		border: 1px solid rgba(255, 255, 255, 0.35);
		border-radius: var(--radius);
		box-shadow:
			0 32px 80px rgba(0, 0, 0, 0.3),
			inset 0 2px 0 rgba(255, 255, 255, 0.4);
	}
	:global(.auth-bg h1) {
		font-size: 26px;
		color: #fff;
	}
	:global(.auth-bg .sub) {
		color: rgba(255, 255, 255, 0.75);
		margin: 8px 0 24px;
		font-size: 15px;
	}
	:global(.auth-bg label) {
		display: flex;
		flex-direction: column;
		gap: 7px;
		margin-bottom: 16px;
	}
	:global(.auth-bg label span) {
		font-size: 14px;
		font-weight: 500;
		color: rgba(255, 255, 255, 0.85);
	}
	:global(.auth-bg input) {
		padding: 12px 15px;
		border: 1px solid rgba(255, 255, 255, 0.35);
		border-radius: var(--radius-sm);
		font-family: var(--font-sans);
		font-size: 15px;
		background: rgba(255, 255, 255, 0.12);
		color: #fff;
	}
	:global(.auth-bg input::placeholder) {
		color: rgba(255, 255, 255, 0.5);
	}
	:global(.auth-bg input:focus) {
		outline: none;
		border-color: #fff;
		background: rgba(255, 255, 255, 0.18);
	}
	:global(.auth-bg .block) {
		display: flex;
		justify-content: center;
		width: 100%;
		margin-top: 6px;
		background: #fff;
		color: var(--navy-teal);
		height: 50px;
		font-weight: 700;
	}
	:global(.auth-bg .block:hover) {
		background: rgba(255, 255, 255, 0.9);
		transform: translateY(-2px);
		box-shadow: 0 12px 30px rgba(0, 0, 0, 0.25);
	}
	:global(.auth-bg .block[disabled]) {
		opacity: 0.7;
		cursor: not-allowed;
	}
	:global(.auth-bg .foot) {
		text-align: center;
		margin-top: 20px;
		font-size: 14px;
		color: rgba(255, 255, 255, 0.75);
	}
	:global(.auth-bg .foot a),
	:global(.auth-bg .row a) {
		color: #fff;
		font-weight: 600;
	}
	:global(.auth-bg .row) {
		text-align: right;
		font-size: 13px;
		margin-bottom: 16px;
		margin-top: -4px;
	}
	:global(.auth-bg .alert) {
		background: rgba(220, 38, 38, 0.2);
		color: #fff;
		border: 1px solid rgba(255, 255, 255, 0.35);
		padding: 11px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
		margin-bottom: 16px;
	}
	:global(.auth-bg .ok) {
		background: rgba(30, 123, 69, 0.25);
		color: #fff;
		border: 1px solid rgba(255, 255, 255, 0.35);
		padding: 11px 14px;
		border-radius: var(--radius-sm);
		font-size: 14px;
		margin-bottom: 16px;
	}
</style>
