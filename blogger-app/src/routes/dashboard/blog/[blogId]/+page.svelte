<script lang="ts">
	import { page } from "$app/state";
	import type { PageProps } from "./$types";

	let { data }: PageProps = $props();
    const { blogId } = page.params;
    const blog = data?.[blogId];
</script>

<main class="flex flex-col gap-8 mt-8">
    <h1>{blog.title}</h1>
    {@html blog.content}
</main>

