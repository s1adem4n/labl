<script lang="ts">
	import Search from '~icons/lucide/search';
	import type { SearchResponse } from '$lib/search';
	import { searchImages } from '$lib/pb';
	import AddImage from './add-image.svelte';
	import { Dialog } from '$lib/ui';
	import AddImageFile from './add-image-file.svelte';
	import ImagesSearchImage from './images-search-image.svelte';

	let {
		open = $bindable()
	}: {
		open?: boolean;
	} = $props();

	let search = $state('');
	let loading = $state(false);

	let searchResponse: SearchResponse | null = $state(null);

	const getSearch = async () => {
		if (loading) return;

		if (search.trim() === '') {
			searchResponse = null;
			return;
		}

		loading = true;
		searchResponse = await searchImages(search, true);
		loading = false;
	};

	const getMore = async () => {
		if (!searchResponse || loading) return;

		loading = true;
		const newResponse = await searchImages(
			search,
			true,
			searchResponse.queries.nextPage[0].startIndex
		);
		const items = searchResponse.items.concat(newResponse.items);

		searchResponse = {
			...newResponse,
			items
		};
		loading = false;
	};

	let dialogOpen = $state(false);
	let fileDialogOpen = $state(false);
	let url = $state('');
</script>

<Dialog bind:open title="Bilder suchen">
	<div class="relative flex items-center">
		<input
			class="w-full rounded-none border-x-0 border-b border-t-0 border-gray-200 px-4 py-2 focus:outline-none"
			bind:value={search}
			placeholder="Suchen ..."
			onkeydown={(e) => {
				if (e.key === 'Enter') {
					getSearch();
				}
			}}
		/>
		<button class="absolute right-2" onclick={getSearch}>
			<Search class="h-6 w-6 text-gray-400" />
		</button>
	</div>
	<div class="xs:grid-cols-3 grid grid-cols-2 gap-2 overflow-y-auto p-2 sm:grid-cols-4">
		{#if searchResponse}
			{#each searchResponse.items as item}
				<ImagesSearchImage
					{item}
					onclick={() => {
						url = item.link;
						dialogOpen = true;
					}}
				/>
			{/each}
			<button class="col-span-full py-2 text-blue-500" onclick={getMore}
				>{loading ? 'Lade ...' : 'Mehr laden'}</button
			>
		{:else}
			<span class="col-span-full text-center text-gray-400">Keine Bilder gefunden</span>
			<button class="col-span-full text-blue-500" onclick={() => (dialogOpen = true)}
				>Datei hochladen?</button
			>
		{/if}
	</div>

	<AddImage bind:open={dialogOpen} {url} />
	<AddImageFile bind:open={fileDialogOpen} />
</Dialog>
