<script lang="ts">
	import { onMount } from 'svelte';
	import '../app.css';
	import * as state from '$lib/state.svelte';

	let { children } = $props();

	let subscribed = false;

	onMount(() => {
		if (subscribed) return;

		state.load();

		const unsubscribe = state.subscribe();
		subscribed = true;

		window.addEventListener('beforeunload', () => {
			unsubscribe.then((unsub) => {
				if (subscribed) {
					unsub();
					subscribed = false;
				}
			});
		});

		return () => {
			unsubscribe.then((unsub) => {
				if (subscribed) {
					unsub();
					subscribed = false;
				}
			});
		};
	});
</script>

<svelte:head>
	{#if import.meta.env.PROD}
		<script src="https://statify.bukohome.com/tracker.js"></script>
	{/if}
</svelte:head>

<div class="mx-auto flex h-full w-full max-w-2xl flex-col gap-4 overflow-y-auto p-4">
	{@render children()}
</div>
