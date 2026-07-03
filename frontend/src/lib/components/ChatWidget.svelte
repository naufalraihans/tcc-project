<script lang="ts">
	import { tick } from 'svelte';
	import { api } from '$lib/api';
	import { MessageCircle, X, ArrowUp } from 'lucide-svelte';

	type Msg = { role: 'user' | 'assistant'; content: string };

	let open = $state(false);
	let input = $state('');
	let loading = $state(false);
	let scroller = $state<HTMLDivElement | null>(null);
	let messages = $state<Msg[]>([
		{
			role: 'assistant',
			content:
				'Halo, saya asisten TCC ITPLN. Tanyakan apa saja seputar program pelatihan, sertifikasi, atau cara mendaftar.'
		}
	]);

	async function scrollDown() {
		await tick();
		if (scroller) scroller.scrollTop = scroller.scrollHeight;
	}

	async function send() {
		const text = input.trim();
		if (!text || loading) return;
		messages = [...messages, { role: 'user', content: text }];
		input = '';
		loading = true;
		scrollDown();
		try {
			const res = await api<{ reply: string }>('/chat', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ messages })
			});
			messages = [...messages, { role: 'assistant', content: res.reply }];
		} catch {
			messages = [
				...messages,
				{
					role: 'assistant',
					content:
						'Maaf, terjadi kendala sesaat. Coba lagi, atau hubungi trainingcenter@itpln.ac.id.'
				}
			];
		} finally {
			loading = false;
			scrollDown();
		}
	}

	function onKey(e: KeyboardEvent) {
		if (e.key === 'Enter' && !e.shiftKey) {
			e.preventDefault();
			send();
		}
	}

	function toggle() {
		open = !open;
		if (open) scrollDown();
	}
</script>

<div class="cw">
	{#if open}
		<div class="panel" role="dialog" aria-label="Asisten TCC ITPLN">
			<header class="head">
				<div class="who">
					<span class="dot"></span>
					<div>
						<div class="title">Asisten TCC ITPLN</div>
						<div class="sub">Biasanya membalas seketika</div>
					</div>
				</div>
				<button class="icon" aria-label="Tutup" onclick={toggle}><X size={18} /></button>
			</header>

			<div class="msgs" bind:this={scroller}>
				{#each messages as m}
					<div class="msg {m.role}">{m.content}</div>
				{/each}
				{#if loading}
					<div class="msg assistant typing"><span></span><span></span><span></span></div>
				{/if}
			</div>

			<div class="composer">
				<textarea
					bind:value={input}
					onkeydown={onKey}
					rows="1"
					placeholder="Tulis pertanyaan…"
					aria-label="Pesan"
				></textarea>
				<button class="send" aria-label="Kirim" onclick={send} disabled={loading || !input.trim()}>
					<ArrowUp size={18} strokeWidth={2.5} />
				</button>
			</div>
		</div>
	{/if}

	<button class="fab" class:hide={open} aria-label="Buka asisten" onclick={toggle}>
		<MessageCircle size={24} strokeWidth={2} />
	</button>
</div>

<style>
	.cw {
		position: fixed;
		right: 24px;
		bottom: 24px;
		z-index: 200;
		display: flex;
		flex-direction: column;
		align-items: flex-end;
		gap: 14px;
	}
	.fab {
		width: 56px;
		height: 56px;
		border-radius: 50%;
		border: none;
		cursor: pointer;
		display: grid;
		place-items: center;
		color: #fff;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		box-shadow: 0 12px 30px rgba(12, 79, 106, 0.4);
		transition:
			transform 0.2s ease,
			box-shadow 0.2s ease;
	}
	.fab:hover {
		transform: translateY(-3px) scale(1.04);
		box-shadow: 0 16px 38px rgba(12, 79, 106, 0.5);
	}
	.fab.hide {
		display: none;
	}
	.panel {
		width: min(380px, calc(100vw - 32px));
		height: min(560px, calc(100vh - 120px));
		display: flex;
		flex-direction: column;
		background: var(--white);
		border: 1px solid var(--border);
		border-radius: 20px;
		overflow: hidden;
		box-shadow: 0 30px 70px rgba(16, 21, 27, 0.28);
		animation: pop 0.24s cubic-bezier(0.22, 1, 0.36, 1);
	}
	@keyframes pop {
		from {
			opacity: 0;
			transform: translateY(16px) scale(0.98);
		}
		to {
			opacity: 1;
			transform: none;
		}
	}
	.head {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 14px 16px;
		color: #fff;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
	}
	.who {
		display: flex;
		align-items: center;
		gap: 11px;
	}
	.who .dot {
		width: 9px;
		height: 9px;
		border-radius: 50%;
		background: #4ade80;
		box-shadow: 0 0 0 3px rgba(74, 222, 128, 0.3);
	}
	.title {
		font-family: var(--font-display);
		font-weight: 700;
		font-size: 15px;
	}
	.sub {
		font-size: 12px;
		opacity: 0.85;
	}
	.icon {
		background: rgba(255, 255, 255, 0.15);
		border: none;
		color: #fff;
		width: 30px;
		height: 30px;
		border-radius: 8px;
		display: grid;
		place-items: center;
		cursor: pointer;
	}
	.icon:hover {
		background: rgba(255, 255, 255, 0.28);
	}
	.msgs {
		flex: 1;
		overflow-y: auto;
		padding: 16px;
		display: flex;
		flex-direction: column;
		gap: 10px;
		background: var(--off-white);
	}
	.msg {
		max-width: 82%;
		padding: 10px 14px;
		border-radius: 14px;
		font-size: 14px;
		line-height: 1.5;
		white-space: pre-line;
	}
	.msg.assistant {
		align-self: flex-start;
		background: var(--white);
		border: 1px solid var(--border);
		color: var(--ink);
		border-bottom-left-radius: 4px;
	}
	.msg.user {
		align-self: flex-end;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		color: #fff;
		border-bottom-right-radius: 4px;
	}
	.typing {
		display: flex;
		gap: 4px;
		align-items: center;
	}
	.typing span {
		width: 7px;
		height: 7px;
		border-radius: 50%;
		background: var(--muted);
		animation: blink 1.2s infinite ease-in-out;
	}
	.typing span:nth-child(2) {
		animation-delay: 0.2s;
	}
	.typing span:nth-child(3) {
		animation-delay: 0.4s;
	}
	@keyframes blink {
		0%,
		60%,
		100% {
			opacity: 0.25;
			transform: translateY(0);
		}
		30% {
			opacity: 1;
			transform: translateY(-3px);
		}
	}
	.composer {
		display: flex;
		align-items: flex-end;
		gap: 8px;
		padding: 12px;
		border-top: 1px solid var(--border);
		background: var(--white);
	}
	textarea {
		flex: 1;
		resize: none;
		max-height: 96px;
		border: 1px solid var(--border);
		border-radius: 12px;
		padding: 10px 12px;
		font-family: var(--font-sans);
		font-size: 14px;
		line-height: 1.4;
		color: var(--ink);
	}
	textarea:focus {
		outline: none;
		border-color: var(--sky-blue);
	}
	.send {
		width: 40px;
		height: 40px;
		flex-shrink: 0;
		border: none;
		border-radius: 11px;
		cursor: pointer;
		display: grid;
		place-items: center;
		color: #fff;
		background: linear-gradient(135deg, var(--navy-teal), var(--sky-blue));
		transition:
			opacity 0.15s ease,
			transform 0.15s ease;
	}
	.send:disabled {
		opacity: 0.45;
		cursor: not-allowed;
	}
	.send:not(:disabled):hover {
		transform: translateY(-1px);
	}
</style>
