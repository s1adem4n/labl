<script lang="ts">
	import type { Item } from '$lib/search';
	import { onMount } from 'svelte';

	let { item, onclick }: { item: Item; onclick: () => void } = $props();

	let image: HTMLImageElement | null = $state(null);
	let failed = $state(false);

	onMount(() => {
		image!.onerror = () => {
			failed = true;
		};
	});
</script>

{#if !failed}
	<div class="flex flex-col gap-2 rounded-xl border border-gray-200 p-2">
		<div class="relative flex aspect-square w-full overflow-hidden rounded-md">
			<img
				class="h-full w-full object-contain text-transparent"
				src={item.link}
				alt={item.title}
				bind:this={image}
			/>
		</div>
		<span class="h-px w-full bg-gray-200"></span>
		<button class="text-blue-500" {onclick}>Hinzufügen</button>
	</div>
{/if}
