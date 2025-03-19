import { CONNECTION_URL } from '$lib/constants/grpc';
import { UserService } from '$lib/proto/user/v1/user_pb';
import { createClient } from '@connectrpc/connect';

import { createGrpcTransport } from '@connectrpc/connect-node';

const transport = createGrpcTransport({
	baseUrl: CONNECTION_URL,
});

export const userClient = {
	login: createClient(UserService, transport).login,
    createUser: createClient(UserService, transport).createUser
};
