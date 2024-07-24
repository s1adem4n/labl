<script lang="ts">
	import type { RenderRequest, Template } from '$lib/pb';
	import { Checkbox, Input, Label } from '$lib/ui';

	let {
		request = $bindable(),
		template
	}: {
		request: RenderRequest;
		template: Template;
	} = $props();

	const round = (num: number) => parseFloat(num.toFixed(1));
</script>

<div class="flex flex-col">
	<Label for="size">Größe</Label>
	<div class="flex items-center gap-2">
		<Input
			type="number"
			min={1}
			bind:value={request.size.width}
			oninput={() => {
				request.size.height = round(
					(request.size.width || 0) * (template.data.aspectRatio[1] / template.data.aspectRatio[0])
				);
			}}
			id="width"
			placeholder="Breite"
		/>
		<span class="-mt-1 text-xl">x</span>
		<Input
			type="number"
			min={1}
			bind:value={request.size.height}
			oninput={() => {
				request.size.width = round(
					(request.size.height || 0) * (template.data.aspectRatio[0] / template.data.aspectRatio[1])
				);
			}}
			id="height"
			placeholder="Höhe"
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
<div class="flex items-center gap-2">
	<Checkbox bind:value={request.outline} id="outline" />
	<Label for="outline">Umrandung</Label>
</div>
