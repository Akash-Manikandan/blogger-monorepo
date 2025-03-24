import { blogClient } from '$lib/grpc/blogs';
import type { PageServerLoad } from './$types';
import { marked } from 'marked';
export const load: PageServerLoad = async ({ cookies, setHeaders, params }) => {
	const token = cookies.get('token');
	const headers = new Headers();
	headers.append('Content-Type', 'application/json');
	headers.append('Accept', 'application/json');
	headers.append('cache-control', 'max-age=3600');
	headers.append('Authorization', `Bearer ${token}`);
	const posts = await blogClient.getBlog(
		{
			id: params.blogId
		},
		{ headers }
	);
	setHeaders({
		'cache-control': 'max-age=3600'
	});

    posts.blog.content = await marked(posts.blog.content);

	return {
		[params.blogId]: posts.blog
	};
};
