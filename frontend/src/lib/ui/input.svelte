<script lang="ts">
	let {
		value = $bindable(),
		placeholder,
		type = 'text',
		id = '',
		onenter,
		onchange,
		oninput,
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
		oninput?: () => void;
		disabled?: boolean;
	} = $props();
</script>

<input
	class="w-full appearance-none rounded-xl border border-gray-200 px-4 py-2 shadow-sm focus-visible:border-blue-500 focus-visible:outline-none disabled:cursor-not-allowed disabled:opacity-50"
	{placeholder}
	{id}
	{type}
	onkeydown={(e) => {
		if (e.key === 'Enter') {
			onenter?.();
		}
	}}
	{onchange}
	{oninput}
	{disabled}
	bind:value
/>
