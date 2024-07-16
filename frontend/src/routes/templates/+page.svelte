<script lang="ts">
	import { page } from '$app/stores';
	import { renderTemplate, type RenderRequest, type Template } from '$lib/pb';
	import { templates } from '$lib/state.svelte';
	import Button from '$lib/ui/button.svelte';
	import Inputs from './inputs.svelte';
	import Options from './options.svelte';

	let template: Template | null = $state(null);

	$effect(() => {
		const foundTemplate = templates.find((t) => t.id === $page.url.searchParams.get('id'));

		if (foundTemplate) {
			template = foundTemplate;
		}
	});

	let request: RenderRequest = $state({
		id: '',
		inputs: {},
		gap: 5,
		margin: 10,
		quantity: 15,
		size: {
			width: 100,
			height: 0
		},
		outline: false
	});

	const inputCount = $derived.by(() => {
		if (!template) return 0;

		return Object.entries(template.data.resources).filter(
			([_, value]) => value.source.type === 'input'
		).length;
	});
</script>

{#if template}
	<h1 class="text-2xl font-bold">Vorlage: {template.name}</h1>
	<div class="flex flex-col gap-2">
		<span class="text-xl"> Eingaben </span>
		<Inputs {template} bind:inputs={request.inputs} />
	</div>
	<div class="flex flex-col gap-2">
		<span class="text-xl"> Optionen </span>
		<Options bind:request />
	</div>

	<Button
		disabled={Object.values(request.inputs).some((v) => v === '') ||
			Object.values(request.inputs).length !== inputCount ||
			request.quantity === 0 ||
			(request.size.width === 0 && request.size.height === 0)}
		onclick={async () => {
			const res = await renderTemplate({
				...request,
				id: template!.id
			});

			const url = URL.createObjectURL(new Blob([res], { type: 'application/pdf' }));

			const a = document.createElement('a');
			a.href = url;
			a.target = '_blank';
			a.click();
		}}>Erstellen</Button
	>
{:else}
	<p>Template not found</p>
{/if}
