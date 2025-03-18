<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input';
	import * as Card from '$lib/components/ui/card';
	import { formSchema, type FormSchema } from './schema';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';

	let { data }: { data: { form: SuperValidated<Infer<FormSchema>> } } = $props();
	const initialForm = data?.form ?? { email: '', password: '' };
	const form = superForm(initialForm, {
		validators: zodClient(formSchema)
	});

	const { form: formData, enhance } = form;
</script>

<main class="flex items-center justify-center h-screen">
    <Card.Root class="w-96">
        <Card.Header>
            <Card.Title>Login</Card.Title>
            <Card.Description>Enter your account details</Card.Description>
        </Card.Header>
        <Card.Content>
            <form method="POST" use:enhance>
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
