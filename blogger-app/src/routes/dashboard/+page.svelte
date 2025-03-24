<script lang="ts">
	import type { PageProps } from './$types';
	import * as Table from '$lib/components/ui/table';
	import { convertToUTCAndFormat } from '$lib/utils/date';
	import Trash from '@lucide/svelte/icons/trash-2';
	let { data }: PageProps = $props();
</script>

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
				<Table.Cell class="font-medium"><a href={`/dashboard/blog/${blog.id}`}>{blog.title}</a></Table.Cell>
				<Table.Cell>{blog.author.username}</Table.Cell>
				<Table.Cell>{convertToUTCAndFormat(blog.createdAt)}</Table.Cell>
				<Table.Cell>{convertToUTCAndFormat(blog.updatedAt)}</Table.Cell>
				<Table.Cell class="flex justify-center"><Trash size={16} /></Table.Cell>
			</Table.Row>
		{/each}
	</Table.Body>
</Table.Root>
