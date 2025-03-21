import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';
import { userClient } from '$lib/grpc/users';
import { ConnectError } from '@connectrpc/connect';

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(formSchema))
	};
};

export const actions: Actions = {
	login: async (event) => {
		const form = await superValidate(event, zod(formSchema));
		if (!form.valid) {
			return fail(400, {
				form,
				status: 'error',
				error: 'Invalid data please check the form'
			});
		}
		try {
			const data = await userClient.login(form.data);
			const { token, ...user } = data;
			event.cookies.set('token', token, {
				path: '/',
				maxAge: 60 * 60 * 24 * 7
			});
			event.cookies.set('user', JSON.stringify(user), {
				path: '/',
				maxAge: 60 * 60 * 24 * 7
			});
			throw redirect(303, '/dashboard');
		} catch (error) {
			if (error && typeof error === 'object' && 'status' in error) {
				if (error.status === 303) throw error;
			}
			if (error instanceof ConnectError) {
				return fail(400, {
					form,
					error: error.message,
					status: 'error'
				});
			}
			return fail(500, {
				form,
				error: 'Unknown error occurred',
				status: 'error'
			});
		}
	}
};
