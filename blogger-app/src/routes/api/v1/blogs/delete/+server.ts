import { json } from '@sveltejs/kit';
import { blogClient } from '$lib/grpc/blogs';

export async function POST({ request, cookies }) {
	const { id } = await request.json();
	const token = cookies.get('token');
	const headers = new Headers();
	headers.append('Content-Type', 'application/json');
	headers.append('Accept', 'application/json');
	headers.append('Authorization', `Bearer ${token}`);
	const { message } = await blogClient.deleteBlog({ id }, { headers });

	return json({ id, message }, { status: 200 });
}
