<script lang="ts">
	import type { PageProps } from './$types';
	import * as Table from '$lib/components/ui/table';
	import { Button } from '$lib/components/ui/button';
	import { convertToUTCAndFormat } from '$lib/utils/date';
	import Trash from '@lucide/svelte/icons/trash-2';
	import PlusIcon from '@lucide/svelte/icons/plus';
	import { goto, invalidateAll } from '$app/navigation';
	import { toast } from 'svelte-sonner';
	let { data }: PageProps = $props();

	const createBlog = () => {
		goto('/dashboard/blog/');
	};

	const deleteBlog = async (blogId: string) => {
		const request = await fetch(`/api/v1/blogs/delete`, {
			method: 'POST',
			body: JSON.stringify({ id: blogId })
		});
		const response = await request.json();
		toast.success(response.message);
		invalidateAll();
	};

</script>

<main class="flex flex-col gap-8 mt-8">
	<Button class="w-fit" onclick={createBlog}>
		<PlusIcon class="mr-2" />
		Create Blog
	</Button>
	<Table.Root>
		<Table.Caption>A list of your recent blogs.</Table.Caption>
		<Table.Header>
			<Table.Row>
				<Table.Head>Blog Title</Table.Head>
				<Table.Head>Author</Table.Head>
				<Table.Head>Created Date</Table.Head>
				<Table.Head>Updated Date</Table.Head>
				<Table.Head class="flex justify-center">Action</Table.Head>
			</Table.Row>
		</Table.Header>
		<Table.Body>
			{#each data.blogs as blog}
				<Table.Row>
					<Table.Cell class="font-medium"
						><a href={`/dashboard/blog/${blog.id}`}>{blog.title}</a></Table.Cell
					>
					<Table.Cell>{blog.author.username}</Table.Cell>
					<Table.Cell>{convertToUTCAndFormat(blog.createdAt)}</Table.Cell>
					<Table.Cell>{convertToUTCAndFormat(blog.updatedAt)}</Table.Cell>
					<Table.Cell class="flex justify-center"><Trash size={16} class="cursor-pointer" onclick={() => deleteBlog(blog.id)} /></Table.Cell>
				</Table.Row>
			{/each}
		</Table.Body>
	</Table.Root>
</main>
