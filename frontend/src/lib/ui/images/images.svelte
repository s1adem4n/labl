<script lang="ts">
	import type { Image } from '$lib/pb';
	import pb from '$lib/pb';
	import { images } from '$lib/state.svelte';
	import { shouldCloseDialog } from '$lib/utils';

	let {
		open = $bindable(),
		select = false,
		value = $bindable()
	}: {
		open?: boolean;
		select?: boolean;
		value?: string;
	} = $props();

	let dialog: HTMLDialogElement;
	$effect(() => {
		if (open) {
			dialog.showModal();
		} else {
			dialog.close();
		}
	});

	const tags = $derived.by(() => {
		const tags: string[] = [''];

		images.forEach((image) => {
			if (!tags.includes(image.tag)) {
				tags.push(image.tag);
			}
		});

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
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<dialog
	bind:this={dialog}
	onclose={() => {
		open = false;
	}}
	onclick={(e) => {
		if (shouldCloseDialog(e)) {
			open = false;
		}
	}}
	class="backdrop:bg-white/50 m-auto max-w-3xl bg-transparent p-4"
>
	<div class="flex h-[600px] rounded-xl border border-gray-200 bg-white">
		<div class="flex flex-col border-r border-gray-200 min-w-[30%] p-4 gap-2">
			<span class="text-xl font-bold -mt-2">Bilder</span>
			{#each tags as tag}
				<button
					class="text-left rounded-xl"
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
		</div>
		<div class="flex flex-col w-full">
			<input
				class="focus:outline-none border-t-0 border-x-0 border-b border-gray-200 py-2 px-4 w-full"
				bind:value={searchFilter}
				placeholder="Suchen ..."
			/>
			<div class="grid grid-cols-3 w-full p-2 gap-2 overflow-auto" bind:this={scrollContainer}>
				{#if filteredImages.length === 0}
					<span class="col-span-3 text-center text-gray-400">Keine Bilder gefunden</span>
				{/if}

				{#each filteredImages as image}
					<div class="flex flex-col rounded-xl gap-1 p-2 border border-gray-200 h-fit">
						<div class="w-full h-full aspect-square">
							<img
								alt={image.name}
								class="object-contain w-full h-full"
								src={pb.files.getUrl(image, image.image, { thumb: '300x300f' })}
							/>
						</div>
						<span class="text-ellipsis overflow-hidden">
							{image.name}
						</span>
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
</dialog>
