import { blogClient } from '$lib/grpc/blogs';
import type { PageServerLoad } from './$types';
import { marked } from 'marked';
export const load: PageServerLoad = async ({ cookies, params }) => {
	const token = cookies.get('token');
	const headers = new Headers();
	headers.append('Content-Type', 'application/json');
	headers.append('Accept', 'application/json');
	headers.append('Authorization', `Bearer ${token}`);
	const posts = await blogClient.getBlog(
		{
			id: params.blogId
		},
		{ headers }
	);

    posts.blog.content = await marked(posts.blog.content);

	return {
		[params.blogId]: posts.blog
	};
};
