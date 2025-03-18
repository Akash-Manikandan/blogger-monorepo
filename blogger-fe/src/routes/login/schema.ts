import { z } from 'zod';

export const formSchema = z.object({
	email: z.string().email().max(50).default(''),
	password: z.string().min(6).max(50).default('')
});

export type FormSchema = typeof formSchema;
