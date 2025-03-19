import type { LayoutServerLoad } from './$types';
import type { User } from './types';

export const load = (async ({ cookies }) => {
	const user: User = JSON.parse(cookies.get('user') ?? '{}');
	return user;
}) satisfies LayoutServerLoad;
