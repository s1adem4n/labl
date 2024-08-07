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

	let image: HTMLImageElement | null = $state(null);

	let name = $state('');
	let tag = $state('');
</script>

<Dialog bind:open title="Bild hinzufügen">
	<div class="flex min-h-full flex-col items-center gap-2 overflow-y-auto p-2">
		<div class="aspect-video w-full max-w-md rounded-xl bg-gray-100">
			<img class="h-full w-full object-contain" src={url} alt="Vorschau" bind:this={image} />
		</div>
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
				if (!url && image?.naturalWidth !== 0) return;

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
