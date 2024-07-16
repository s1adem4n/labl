<script lang="ts">
	import type { Snippet } from 'svelte';

	let {
		children,
		disabled = false,
		onclick,
		width = 'w-full',
		ghost = false
	}: {
		children: Snippet;
		disabled?: boolean;
		onclick?: () => void;
		width?: string;
		ghost?: boolean;
	} = $props();
</script>

<button
	class="{width} rounded-xl px-4 py-2 text-center shadow-sm transition-colors disabled:opacity-50 {ghost
		? 'border border-gray-200 hover:bg-gray-50'
		: 'bg-blue-500 text-white hover:bg-blue-600'}"
	{disabled}
	{onclick}
>
	{@render children()}
</button>
