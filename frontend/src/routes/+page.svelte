<script lang="ts">
	import { base } from '$app/paths';
	import { templates } from '$lib/state.svelte';
</script>

<h1 class="text-2xl font-bold">Vorlagen</h1>

{#each templates as template}
	<a class="hover:underline" href="{base}/templates?id={template.id}">
		{template.name}
	</a>
{/each}
