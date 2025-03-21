<script lang="ts">
	import { createAvatar } from '@dicebear/core';
	import { initials } from '@dicebear/collection';
	import * as Sidebar from '$lib/components/ui/sidebar';
	import { useSidebar } from '$lib/components/ui/sidebar';
	import type { Component } from 'svelte';
	import { ChevronUp, type IconProps } from '@lucide/svelte';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { cn } from '$lib/utils';
	import TextOverflow from '../text-overflow/text-overflow.svelte';
	type MenuItem = {
		title: string;
		url: string;
		icon: Component<IconProps, {}, ''>;
	};
	let {
		menuItems,
		currentPath,
		userName
	}: { menuItems: MenuItem[]; currentPath: string; userName: string } = $props();
	const sidebar = useSidebar();
	const svg = $derived(createAvatar(initials, {
		seed: userName
	}).toDataUri());

</script>

<Sidebar.Root variant="inset" collapsible="icon">
	<Sidebar.Header />
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupLabel>Blogger Application</Sidebar.GroupLabel>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each menuItems as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton isActive={currentPath === item.url}>
								{#snippet child({ props })}
									<Tooltip.Provider>
										<Tooltip.Root>
											<Tooltip.Trigger class="w-full">
												<a href={item.url} {...props}>
													<item.icon />
													<span>{item.title}</span>
												</a>
											</Tooltip.Trigger>
											<Tooltip.Content side="right">
												<p>{item.title}</p>
											</Tooltip.Content>
										</Tooltip.Root>
									</Tooltip.Provider>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
	<Sidebar.Footer>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<DropdownMenu.Root>
					<DropdownMenu.Trigger>
						{#snippet child({ props })}
							<Sidebar.MenuButton
								{...props}
								class={cn(
									'data-[state=open]:bg-sidebar-accent data-[state=open]:text-sidebar-accent-foreground',
									!sidebar.open && 'h-6 w-6 p-0'
								)}
								size="lg"
								isActive={false}
							>
								<img src={svg} alt="avatar" class="w-7 rounded-full" />
								<TextOverflow text={userName} />
								<ChevronUp class="ml-auto" />
							</Sidebar.MenuButton>
						{/snippet}
					</DropdownMenu.Trigger>
					<DropdownMenu.Content side="top" class={sidebar.open ? 'w-[240px]' : 'w-full'}>
						<DropdownMenu.Item>
							<span>Account</span>
						</DropdownMenu.Item>
						<DropdownMenu.Item>
							<span>Billing</span>
						</DropdownMenu.Item>
						<DropdownMenu.Item>
							<span>Sign out</span>
						</DropdownMenu.Item>
					</DropdownMenu.Content>
				</DropdownMenu.Root>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Footer>
</Sidebar.Root>
