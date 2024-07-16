<script lang="ts">
	import Check from '~icons/lucide/check';

	let {
		value = $bindable(),
		id = '',
		onchange
	}: {
		value?: boolean;
		id?: string;
		onchange?: (
			event: Event & {
				currentTarget: HTMLInputElement;
			}
		) => void;
	} = $props();
</script>

<div class="relative flex items-center justify-center">
	<input
		type="checkbox"
		class="h-6 w-6 appearance-none rounded-full border border-gray-200 shadow-sm checked:border-blue-500 checked:bg-blue-500"
		{id}
		bind:checked={value}
		{onchange}
	/>
	<Check class="pointer-events-none absolute h-4 w-4 {value ? 'text-white' : 'text-gray-200'}" />
</div>
