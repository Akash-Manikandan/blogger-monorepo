<script lang="ts">
	import '../app.css';
	import { Toaster } from '$lib/components/ui/sonner';
	import LayoutTemplate from '$lib/components/custom/layout/layout-template.svelte';

	import { page } from '$app/state';
	import { PUBLIC_ROUTES } from '$lib/constants/routes';
	import { goto } from '$app/navigation';
	import type { User } from './types';
	import { type Snippet } from 'svelte';

	let { children, data }: { children: Snippet<[]>; data: User } = $props();
	let userName = $state('Unknown');
	$effect(() => {
		if (!data && !Object.values(PUBLIC_ROUTES).includes(page.url.pathname)) {
			goto(PUBLIC_ROUTES.LOGIN);
		} else if (!Object.values(PUBLIC_ROUTES).includes(page.url.pathname)) {
			userName = data?.username ?? 'Unknown';
		}
	});
</script>

{#if Object.values(PUBLIC_ROUTES).includes(page.url.pathname)}
	{@render children?.()}
{:else}
	<LayoutTemplate {children} currentPath={page.url.pathname} userName={userName} />
{/if}

<Toaster />
