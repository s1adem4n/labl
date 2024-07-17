<script lang="ts">
	import { type Template } from '$lib/pb';
	import { images } from '$lib/state.svelte';
	import { Images, Label, Input } from '$lib/ui';

	let {
		template,
		inputs = $bindable({})
	}: {
		template: Template;
		inputs: Record<string, any>;
	} = $props();

	const inputResources = $derived(
		Object.entries(template.data.resources).filter(([_, value]) => value.source.type === 'input')
	);

	const dialogs: Record<string, boolean> = $state({});
</script>

{#each inputResources as [name, resource]}
	<div class="flex flex-col">
		<Label for={name}>{resource.label !== '' ? resource.label : name}</Label>
		{#if resource.type === 'text'}
			<Input bind:value={inputs[name]} id={name} />
		{:else if resource.type === 'image'}
			<Images select={true} bind:open={dialogs[name]} bind:value={inputs[name]} />
			<button class="text-left text-blue-500 underline" onclick={() => (dialogs[name] = true)}>
				{!inputs[name]
					? 'Bild auswählen'
					: `Bild ändern (${images.find((image) => image.id === inputs[name])?.name})`}
			</button>
		{/if}
	</div>
{/each}
