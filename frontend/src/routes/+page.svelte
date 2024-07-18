<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import pb from '$lib/pb';
	import { templates } from '$lib/state.svelte';
</script>

<h1 class="text-2xl font-bold">Vorlagen</h1>

<div class="grid w-full grid-cols-2 gap-2 sm:grid-cols-3">
	{#each templates as template}
		<button
			onclick={() => goto(`${base}/templates?id=${template.id}`)}
			class="group flex flex-col gap-2 rounded-xl border border-gray-200 p-2"
		>
			<div class="relative flex aspect-square w-full overflow-hidden rounded-md">
				<img
					class="h-full w-full object-contain text-transparent"
					src={pb.files.getUrl(template, template.thumbnail, { thumb: '250x250f' })}
					alt={template.name}
				/>
				{#if !template.thumbnail}
					<div class="absolute inset-0 flex items-center justify-center bg-gray-100 text-gray-400">
						Kein Vorschaubild
					</div>
				{/if}
			</div>
			<span class="h-px w-full bg-gray-200"></span>
			<span class="mx-auto text-center font-bold group-hover:underline">
				{template.name}
			</span>
		</button>
	{/each}
</div>
