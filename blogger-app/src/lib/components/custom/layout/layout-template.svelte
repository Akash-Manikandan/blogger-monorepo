<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar';
	import AppSidebar from '$lib/components/custom/sidebar/app-sidebar.svelte';
	import House from '@lucide/svelte/icons/house';
	import Search from '@lucide/svelte/icons/search';
	import Settings from '@lucide/svelte/icons/settings';
	import TrendingUp from '@lucide/svelte/icons/trending-up';

	const menuItems = [
		{
			title: 'Dashboard',
			url: '/dashboard',
			icon: House
		},
		{
			title: 'Trending',
			url: '/trending',
			icon: TrendingUp
		},
		{
			title: 'Search',
			url: '/search',
			icon: Search
		},
		{
			title: 'Settings',
			url: '/settings',
			icon: Settings
		}
	];

	let { children, currentPath, userName } = $props();
</script>

<Sidebar.Provider class="font-display h-screen overflow-hidden">
	<AppSidebar {menuItems} {currentPath} {userName} />
	<main class="bg-primary-foreground m-2 w-full h-screen rounded-xl">
		<div class="h-full overflow-scroll">
			<div class="mx-2 my-4">
				<Sidebar.Trigger class="m-2" />
				{@render children?.()}
			</div>
		</div>
	</main>
</Sidebar.Provider>
