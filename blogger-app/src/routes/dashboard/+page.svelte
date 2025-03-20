<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
    $inspect(data);
</script>

{#each data.blogs as d}
	<div>
		<h1>{d.title}</h1>
		<p>{d.content}</p>
	</div>
{/each}
