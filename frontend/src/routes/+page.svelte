<script lang="ts">
	import { goto } from '$app/navigation';
	import { base } from '$app/paths';
	import pb from '$lib/pb';
	import { templates } from '$lib/state.svelte';
</script>

<h1 class="text-2xl font-bold">Vorlagen</h1>

<div class="grid grid-cols-2 w-full gap-2">
	{#each templates as template}
		<button
			onclick={() => goto(`${base}/templates?id=${template.id}`)}
			class="flex flex-col p-2 border border-gray-200 rounded-xl gap-2 group"
		>
			<div class="flex w-full aspect-square relative">
				<img
					class="object-contain w-full h-full text-transparent"
					src={pb.files.getUrl(template, template.thumbnail, { thumb: '250x250f' })}
					alt={template.name}
				/>
				{#if !template.thumbnail}
					<div
						class="absolute inset-0 flex items-center justify-center bg-gray-100 rounded-md text-gray-400"
					>
						Kein Vorschaubild
					</div>
				{/if}
			</div>
			<span class="w-full h-px bg-gray-200"></span>
			<span class="font-bold text-center group-hover:underline">
				{template.name}
			</span>
		</button>
	{/each}
</div>
