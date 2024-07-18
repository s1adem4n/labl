<script lang="ts">
	import Search from '~icons/lucide/search';
	import X from '~icons/lucide/x';
	import { shouldCloseDialog } from '$lib/utils';
	import type { SearchResponse } from '$lib/search';
	import { searchImages } from '$lib/pb';
	import AddImage from './add-image.svelte';

	let {
		open = $bindable()
	}: {
		open?: boolean;
	} = $props();

	let dialog: HTMLDialogElement;
	$effect(() => {
		if (open) {
			document.body.style.overflow = 'hidden';
			dialog.showModal();
		} else {
			document.body.style.overflow = '';
			dialog.close();
		}
	});

	let search = $state('');
	let loading = $state(false);

	let searchResponse: SearchResponse | null = $state(null);

	const getSearch = async () => {
		if (loading) return;

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
	let url = $state('');
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
	class="m-auto h-[90dvh] w-full max-w-2xl max-w-3xl bg-transparent backdrop:bg-white/50"
>
	<div class="flex h-full flex-col rounded-xl border border-gray-200 bg-white">
		<div class="relative flex w-full items-center justify-center border-b border-gray-200 py-2">
			<h1 class="text-xl font-bold">Bilder suchen</h1>
			<button class="absolute right-2" onclick={() => (open = false)}>
				<X class="h-6 w-6 text-gray-400" />
			</button>
		</div>
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
					<div class="flex flex-col gap-2 rounded-xl border border-gray-200 p-2">
						<div class="relative flex aspect-square w-full overflow-hidden rounded-md">
							<img
								class="h-full w-full object-contain text-transparent"
								src={item.link}
								alt={item.title}
							/>
						</div>
						<span class="h-px w-full bg-gray-200"></span>
						<button
							class="text-blue-500"
							onclick={() => {
								url = item.link;
								dialogOpen = true;
							}}>Hinzufügen</button
						>
					</div>
				{/each}
				<button class="col-span-full py-2 text-blue-500" onclick={getMore}
					>{loading ? 'Lade ...' : 'Mehr laden'}</button
				>
			{:else}
				<span class="col-span-full text-center text-gray-400">Keine Bilder gefunden</span>
				<button
					class="col-span-full py-2 text-blue-500"
					onclick={() => {
						url = '';
						dialogOpen = true;
					}}
				>
					Manuell hinzufügen?
				</button>
			{/if}
		</div>
	</div>

	<AddImage bind:open={dialogOpen} bind:url />
</dialog>
