<script lang="ts">
	import pb from '$lib/pb';
	import { Label, Button, Input, Dialog } from '$lib/ui';

	let {
		open = $bindable(),
		url = $bindable()
	}: {
		open?: boolean;
		url?: string;
	} = $props();
	let name = $state('');
	let tag = $state('');
</script>

<Dialog bind:open title="Bild hinzufügen">
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
</Dialog>
