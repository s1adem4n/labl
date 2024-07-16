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

<div class="flex flex-col h-full w-full max-w-xl mx-auto p-4 gap-4 overflow-y-auto">
	{@render children()}
</div>
