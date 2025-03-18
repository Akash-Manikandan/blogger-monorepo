import { UserService } from '$lib/proto/user/v1/user_pb'; // Adjust the import path
import { createClient } from '@connectrpc/connect';

import { createGrpcTransport } from '@connectrpc/connect-node';

const transport = createGrpcTransport({
	baseUrl: 'http://localhost:50051',
});

export const userClient = {
	login: createClient(UserService, transport).login,
    createUser: createClient(UserService, transport).createUser
};
