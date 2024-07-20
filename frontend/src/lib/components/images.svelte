<script lang="ts">
	import { Dialog, Image as ImageDisplay } from '$lib/ui';
	import Trash from '~icons/lucide/trash';
	import type { Image } from '$lib/pb';
	import pb from '$lib/pb';
	import { images } from '$lib/state.svelte';
	import ImagesSearch from './images-search.svelte';

	let {
		open = $bindable(),
		select = false,
		value = $bindable()
	}: {
		open?: boolean;
		select?: boolean;
		value?: string;
	} = $props();

	const tags = $derived.by(() => {
		const tags: string[] = [''];

		images.forEach((image) => {
			if (!tags.includes(image.tag)) {
				tags.push(image.tag);
			}
		});

		tags.sort((a, b) => a.localeCompare(b));

		return tags;
	});

	let searchFilter = $state('');
	let tagFilter = $state('');

	const filteredImages = $derived.by(() => {
		const filteredImages: Image[] = [];

		images.forEach((image) => {
			if (
				(tagFilter === '' || image.tag === tagFilter) &&
				(searchFilter === '' || image.name.toLowerCase().includes(searchFilter.toLowerCase()))
			) {
				filteredImages.push(image);
			}
		});

		filteredImages.sort((a, b) => a.name.localeCompare(b.name));

		return filteredImages;
	});

	let scrollContainer: HTMLDivElement;

	let searchOpen = $state(false);
</script>

<Dialog bind:open size="lg">
	<div class="flex h-full w-full">
		<div class="xs:min-w-48 flex flex-col gap-2 border-r border-gray-200 p-4">
			<span class="-mt-2 text-xl font-bold">Bilder</span>
			<div class="flex h-full flex-col gap-3 overflow-scroll">
				{#each tags as tag}
					<button
						class="text-left leading-snug"
						class:font-bold={tag === tagFilter}
						class:underline={tag === tagFilter}
						onclick={() => {
							tagFilter = tag;
							scrollContainer.scrollTo({ top: 0 });
						}}
					>
						{tag === '' ? 'Alle' : tag}
					</button>
				{/each}
				<button
					class="mt-auto rounded-md font-bold text-blue-500 hover:underline"
					onclick={() => {
						searchOpen = !searchOpen;
					}}
				>
					Bilder suchen
				</button>
				<ImagesSearch bind:open={searchOpen} />
			</div>
		</div>
		<div class="flex w-full flex-col">
			<input
				class="w-full rounded-none border-x-0 border-b border-t-0 border-gray-200 px-4 py-2 focus:outline-none"
				bind:value={searchFilter}
				placeholder="Suchen ..."
			/>
			<div
				class="xs:grid-cols-3 grid w-full grid-cols-2 gap-2 overflow-auto p-2"
				bind:this={scrollContainer}
			>
				{#if filteredImages.length === 0}
					<span class="col-span-3 text-center text-gray-400">Keine Bilder gefunden</span>
				{/if}

				{#each filteredImages as image}
					<div class="relative flex h-fit flex-col gap-1 rounded-xl border border-gray-200 p-2">
						<button
							class="absolute bottom-2 right-2 z-20"
							onclick={() => {
								pb.collection('images').delete(image.id);
							}}
						>
							<Trash class="h-4 w-4 text-red-500" />
						</button>
						<div class="aspect-square h-full w-full">
							<ImageDisplay
								alt={image.name}
								src={pb.files.getUrl(image, image.image, { thumb: '250x250f' })}
							/>
						</div>
						<span class="overflow-hidden text-ellipsis text-nowrap text-center">
							{image.name}
						</span>
						{#if tagFilter === ''}
							<span class="-my-1 text-center text-sm text-gray-400">
								{image.tag}
							</span>
						{/if}
						<button
							class="text-sm text-blue-500 underline"
							onclick={() => {
								if (select) {
									value = image.id;
									open = false;
								}
							}}
						>
							Auswählen
						</button>
					</div>
				{/each}
			</div>
		</div>
	</div>
</Dialog>
