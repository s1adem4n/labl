<script lang="ts">
	import X from '~icons/lucide/x';

	import { shouldCloseDialog } from '$lib/utils';
	import type { Snippet } from 'svelte';

	let {
		open = $bindable(),
		children,
		title,
		size = 'md'
	}: {
		open?: boolean;
		children?: Snippet;
		size?: 'sm' | 'md' | 'lg';
		title?: string;
	} = $props();

	let dialog: HTMLDialogElement;
	$effect(() => {
		if (open) {
			document.body.style.overflow = 'hidden';
			dialog.showModal();
		} else {
			document.body.style.overflow = '';
			dialog.close();
		}
	});

	const sizes = {
		sm: 'max-w-xl',
		md: 'max-w-2xl',
		lg: 'max-w-4xl'
	};
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<dialog
	bind:this={dialog}
	onclose={() => (open = false)}
	onclick={(e) => {
		if (shouldCloseDialog(e)) {
			open = false;
		}
	}}
	class="m-auto h-[90dvh] w-full bg-transparent backdrop:bg-white/50 backdrop:backdrop-blur-sm {sizes[
		size
	]} p-2"
>
	<div class="flex h-full flex-col rounded-xl border border-gray-200 bg-white">
		{#if title}
			<div class="relative flex w-full items-center justify-center border-b border-gray-200 py-2">
				<h1 class="text-xl font-bold">{title}</h1>
				<button class="absolute right-2" onclick={() => (open = false)}>
					<X class="h-6 w-6 text-gray-400" />
				</button>
			</div>
		{/if}
		<div class="h-full overflow-y-auto">
			{@render children?.()}
		</div>
	</div>
</dialog>
