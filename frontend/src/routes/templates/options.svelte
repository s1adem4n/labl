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

	let calculatedSize = $derived.by(() => {
		if (request.size.width === 0) {
			const aspect = template.data.aspectRatio[1] / template.data.aspectRatio[0];
			return {
				width: request.size.height / aspect,
				height: request.size.height
			};
		} else if (request.size.height === 0) {
			const aspect = template.data.aspectRatio[0] / template.data.aspectRatio[1];
			return {
				width: request.size.width,
				height: request.size.width / aspect
			};
		} else {
			return request.size;
		}
	});
</script>

<div class="flex flex-col">
	<Label for="size"
		>Größe ({calculatedSize.width.toFixed(1)}x{calculatedSize.height.toFixed(1)})</Label
	>
	<div class="flex items-center gap-2">
		<Input
			type="number"
			bind:value={request.size.width}
			id="width"
			placeholder="Breite"
			disabled={request.size.height !== 0}
		/>
		<span class="-mt-1 text-xl">x</span>
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
<div class="flex items-center gap-2">
	<Checkbox bind:value={request.outline} id="outline" />
	<Label for="outline">Umrandung</Label>
</div>
