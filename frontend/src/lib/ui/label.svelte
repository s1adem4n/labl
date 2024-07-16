<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		children,
		for: forId
	}: {
		children: Snippet;
		for: string;
	} = $props();
</script>

<label class="text-sm font-bold" for={forId}>
	{@render children()}
</label>
