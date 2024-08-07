<script lang="ts">
	import { Dialog, Label, Input, Button } from '$lib/ui';
	import pb from '$lib/pb';

	let {
		open = $bindable()
	}: {
		open?: boolean;
	} = $props();

	let file: File | null = $state(null);
	let name = $state('');
	let tag = $state('');
</script>

<Dialog bind:open title="Bild hinzufügen">
	<div class="flex min-h-full flex-col items-center justify-center gap-2 p-2">
		<div class="aspect-video w-full max-w-md rounded-xl bg-gray-100">
			{#if file}
				<img class="h-full w-full object-contain" src={URL.createObjectURL(file)} alt="Vorschau" />
			{/if}
		</div>
		<div class="flex w-full flex-col">
			<Label for="file">Datei</Label>
			<Input
				type="file"
				id="file"
				oninput={(e) => {
					if (e.currentTarget.files) {
						file = e.currentTarget.files[0];
					}
				}}
			/>
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
			disabled={!file || !name || !tag}
			onclick={async () => {
				const formData = new FormData();
				formData.append('image', file as Blob);
				formData.append('name', name);
				formData.append('tag', tag);

				await pb.collection('images').create(formData);

				open = false;
			}}
		>
			Hochladen
		</Button>
	</div>
</Dialog>
