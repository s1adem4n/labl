<script lang="ts">
	import { fade } from 'svelte/transition';

	let {
		src,
		alt = ''
	}: {
		src: string;
		alt?: string;
	} = $props();

	let loading = $state(true);

	let image: HTMLImageElement;

	$effect(() => {
		src;
		image.onload = () => {
			loading = false;
		};
	});
</script>

<div class="relative h-full w-full overflow-hidden rounded-xl">
	{#if loading}
		<div class="load-wrapper absolute inset-0">
			<div class="activity"></div>
		</div>
	{/if}

	<img
		bind:this={image}
		loading="lazy"
		class="h-full w-full object-contain transition-opacity duration-500"
		class:opacity-0={loading}
		class:opacity-100={!loading}
		{src}
		{alt}
	/>
</div>
