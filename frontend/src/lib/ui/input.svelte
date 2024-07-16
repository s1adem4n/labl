<script lang="ts">
	let {
		value = $bindable(),
		placeholder,
		type = 'text',
		id = '',
		onenter,
		onchange,
		disabled
	}: {
		value?: unknown;
		type?: string;
		placeholder?: string;
		id?: string;
		onenter?: () => void;
		onchange?: (
			e: Event & {
				currentTarget: EventTarget & HTMLInputElement;
			}
		) => void;
		disabled?: boolean;
	} = $props();
</script>

<input
	class="w-full rounded-xl border border-gray-200 px-4 py-2 shadow-sm focus-visible:border-blue-500 focus-visible:outline-none appearance-none disabled:cursor-not-allowed disabled:opacity-50"
	{placeholder}
	{id}
	{type}
	onkeydown={(e) => {
		if (e.key === 'Enter') {
			onenter?.();
		}
	}}
	onchange={(e) => {
		onchange?.(e);
	}}
	{disabled}
	bind:value
/>
