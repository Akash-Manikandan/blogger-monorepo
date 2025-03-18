# Blogger Monorepo

A monorepo architecture containing a SvelteKit application and a User Service built with Go.

## Project Structure

```
blogger-monorepo/
├── blogger-app/        # SvelteKit fullstack application
├── blogger-service/    # Go-based User Service
└── protobuf/           # Shared Protocol Buffers definitions
```

## Services

### Blogger App (blogger-app)
- **Framework**: SvelteKit - Full-stack web application framework
- **Features**:
  - Server-side rendering (SSR)
  - Client-side routing
  - API endpoints
  - TypeScript support
  - Protocol Buffers client integration
  - Tailwind CSS for styling
  - Authentication handling
  - Blog post management

### User Service (blogger-service)
- **Language**: Go
- **Features**:
  - gRPC service implementation
  - User management (registration, authentication)
  - JWT token handling
  - Database integration (PostgreSQL)
  - Validation middleware

### Shared Protobuf Definitions (protobuf)
- Protocol Buffers service definitions
- Shared between SvelteKit app and User Service
- Includes:
  - User service definitions
  - Validation rules
  - Request/Response types

## Prerequisites

- Node.js (v22+)
- Bun package manager
- Go 1.23+
- Protocol Buffers compiler (protoc)
- PostgreSQL
- Docker (optional)
- Make

## Setup Instructions

### 1. Install Global Dependencies

```bash
# Install Protocol Buffers compiler
brew install protobuf

# Install Bun
curl -fsSL https://bun.sh/install | bash

# Install buf for Protocol Buffers management
brew install bufbuild/buf/buf
```

### 2. Clone and Setup

```bash
# Clone repository
git clone https://github.com/yourusername/blogger-monorepo.git
cd blogger-monorepo

# Install dependencies
make install
```

### 3. Environment Configuration

#### SvelteKit App (.env)
```text
// filepath: /blogger-app/.env
PUBLIC_API_URL=http://localhost:8080
DATABASE_URL=postgresql://user:password@localhost:5432/bloggerdb
```

#### User Service (.env)
```text
// filepath: /blogger-service/.env
DB_CONNECTION=postgresql://user:password@localhost:5432/bloggerdb
JWT_SECRET=your-secret-key
GRPC_PORT=50051
```

### 4. Generate Protocol Buffer Code

```bash
make generate-proto
```

## Development Commands

```bash
# Start SvelteKit development server
make dev-app

# Start User Service
make dev-service

# Start both services
make dev

# Generate Protocol Buffer code
make generate-proto

# Run tests
make test

# Clean generated files
make clean
```

## Development URLs

- SvelteKit App: http://localhost:5173
  - Frontend UI
  - Server API routes
  - Server-side rendering
- User Service gRPC: localhost:50051
  - User management endpoints
  - Authentication service

## API Documentation

### User Service gRPC Endpoints

- `CreateUser`: User registration
- `Login`: User authentication
- `ValidateToken`: JWT token validation

### SvelteKit API Routes

- `/api/auth/*`: Authentication endpoints
- `/api/posts/*`: Blog post management
- `/api/users/*`: User management

## Database Schema

```sql
-- User table
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Contributing

1. Create feature branch from `main`
2. Make changes
3. Generate Protocol Buffer code if needed
4. Test changes
5. Submit pull request

## Common Issues

1. **Protocol Buffer Generation Fails**
   ```bash
   buf mod update
   buf generate
   ```

2. **Database Connection Issues**
   - Verify PostgreSQL is running
   - Check connection strings in .env files

3. **gRPC Connection Refused**
   - Ensure User Service is running
   - Check port configurations

## Learn More

- [SvelteKit Documentation](https://kit.svelte.dev/docs)
- [Protocol Buffers Guide](https://protobuf.dev/)
- [gRPC Documentation](https://grpc.io/docs/)
