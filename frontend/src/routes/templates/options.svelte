<script lang="ts">
	import type { RenderRequest } from '$lib/pb';
	import Checkbox from '$lib/ui/checkbox.svelte';
	import Input from '$lib/ui/input.svelte';
	import Label from '$lib/ui/label.svelte';

	let {
		request = $bindable()
	}: {
		request: RenderRequest;
	} = $props();
</script>

<div class="flex flex-col">
	<Label for="">Größe</Label>
	<div class="flex items-center gap-2">
		<Input
			type="number"
			bind:value={request.size.width}
			id="width"
			placeholder="Breite"
			disabled={request.size.height !== 0}
		/>
		<span class="text-xl -mt-1">x</span>
		<Input
			type="number"
			bind:value={request.size.height}
			id="height"
			placeholder="Höhe"
			disabled={request.size.width !== 0}
		/>
	</div>
</div>
<div class="flex flex-col">
	<Label for="quantity">Anzahl</Label>
	<Input bind:value={request.quantity} id="quantity" type="number" />
</div>
<div class="flex flex-col">
	<Label for="gap">Abstand</Label>
	<Input bind:value={request.gap} id="gap" type="number" />
</div>
<div class="flex flex-col">
	<Label for="margin">Abstand (Rand)</Label>
	<Input bind:value={request.margin} id="margin" type="number" />
</div>
<div class="flex gap-2 items-center">
	<Checkbox bind:value={request.outline} id="outline" />
	<Label for="outline">Umrandung</Label>
</div>
