import { blogClient } from '$lib/grpc/blogs';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies, setHeaders }) => {
	const token = cookies.get('token');
	const headers = new Headers();
	headers.append('Content-Type', 'application/json');
	headers.append('Accept', 'application/json');
	headers.append('cache-control', 'max-age=3600');
	headers.append('Authorization', `Bearer ${token}`);
	const posts = await blogClient.listBlogs({}, { headers });
	setHeaders({
		'cache-control': 'max-age=3600'
	});

	return {
		blogs: posts.blogs
	};
};
