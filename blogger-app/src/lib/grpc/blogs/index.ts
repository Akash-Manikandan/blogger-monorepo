import { CONNECTION_URL } from '$lib/constants/grpc';
import { BlogService } from '$lib/proto/blog/v1/blog_pb';
import { createClient } from '@connectrpc/connect';

import { createGrpcTransport } from '@connectrpc/connect-node';

const transport = createGrpcTransport({
    baseUrl: CONNECTION_URL,
});

export const blogClient = {
    listBlogs: createClient(BlogService, transport).listBlogs,
    getBlog: createClient(BlogService, transport).getBlog
};
