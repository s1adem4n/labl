<script lang="ts">
	import X from '~icons/lucide/x';
	import pb from '$lib/pb';
	import { shouldCloseDialog } from '$lib/utils';
	import Button from './button.svelte';
	import Input from './input.svelte';
	import Label from './label.svelte';

	let {
		open = $bindable(),
		url = $bindable()
	}: {
		open?: boolean;
		url?: string;
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

	let name = $state('');
	let tag = $state('');
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<dialog
	bind:this={dialog}
	onclose={() => {
		open = false;
	}}
	onclick={(e) => {
		if (shouldCloseDialog(e)) {
			open = false;
		}
	}}
	class="m-auto h-[90dvh] w-full max-w-3xl max-w-xl bg-transparent backdrop:bg-white/50"
>
	<div class="flex h-full w-full flex-col rounded-xl border border-gray-200 bg-white">
		<div class="relative flex w-full items-center justify-center border-b border-gray-200 py-2">
			<h1 class="text-xl font-bold">Bilder suchen</h1>
			<button class="absolute right-2" onclick={() => (open = false)}>
				<X class="h-6 w-6 text-gray-400" />
			</button>
		</div>
		<div class="flex h-full flex-col items-center gap-2 overflow-y-auto p-2">
			<img class="w-full max-w-xs" src={url} alt="Vorschau" />
			{#if !url}
				<div class="flex w-full flex-col">
					<Label for="url">URL (Link)</Label>
					<Input bind:value={url} id="url" />
				</div>
			{/if}
			<div class="flex w-full flex-col">
				<Label for="name">Name</Label>
				<Input bind:value={name} id="name" />
			</div>
			<div class="mb-auto flex w-full flex-col">
				<Label for="tag">Tag</Label>
				<Input bind:value={tag} id="tag" />
			</div>
			<Button
				onclick={async () => {
					if (!url) return;

					await pb.send('/images', {
						method: 'POST',
						body: JSON.stringify({
							name,
							tag,
							url
						})
					});
					open = false;
				}}>Hinzufügen</Button
			>
		</div>
	</div>
</dialog>
