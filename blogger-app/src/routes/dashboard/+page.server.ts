import { blogClient } from '$lib/grpc/blogs';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get('token');
	const headers = new Headers();
	headers.append('Authorization', `Bearer ${token}`);
	const posts = await blogClient.listBlogs({}, { headers });
	return {
        blogs: posts.blogs
    }
};
