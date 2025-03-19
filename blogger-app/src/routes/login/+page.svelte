<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();
	const form = superForm(data.form, {
		validators: zodClient(formSchema),
		onResult({ result }) {
            const { status, data } = result as any;
			console.log(data)
			if (status >= 400 && status < 600) {
				toast.error(data?.error ?? "An error occurred");
			} else {
                toast.success("Login successful");
                goto("/dashboard");
            }
		}
	});

	const { form: formData, enhance } = form;
</script>

<main class="flex h-screen items-center justify-center">
	<Card.Root class="w-96">
		<Card.Header>
			<Card.Title>Login</Card.Title>
			<Card.Description>Enter your account details</Card.Description>
		</Card.Header>
		<Card.Content>
			<form method="POST" action="?/login" use:enhance>
				<Form.Field {form} name="email">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Email</Form.Label>
							<Input {...props} type="email" bind:value={$formData.email} />
						{/snippet}
					</Form.Control>
					<Form.FieldErrors />
				</Form.Field>
				<Form.Field {form} name="password">
					<Form.Control>
						{#snippet children({ props })}
							<Form.Label>Password</Form.Label>
							<Input {...props} type="password" bind:value={$formData.password} />
						{/snippet}
					</Form.Control>
					<Form.FieldErrors />
				</Form.Field>
				<Form.Button class="mt-4">Submit</Form.Button>
			</form>
		</Card.Content>
	</Card.Root>
</main>
