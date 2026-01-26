# Moufube - YouTube Mock-up Project

## Overview

Moufube is a YouTube-like video sharing platform built to demonstrate large-scale application development practices. While developed by a single developer, the project follows enterprise-grade team workflows, architecture patterns, and best practices suitable for production environments.

### Architecture

This project uses a **microservices architecture** organized in a **monorepo** structure for simplified development. In production, this would typically be deployed as separate polyrepos, but the monorepo approach streamlines service management during development.

```
moufube/
├── frontend/              # Next.js web application
├── services/
│   ├── api-gateway/       # API Gateway (Go/Gin)
│   └── authentication/    # Authentication Service (Go)
├── deployment/            # Docker Compose orchestration
└── scripts/              # Development automation scripts
```

---

## 🚀 Tech Stack

### Frontend
- **Framework**: Next.js 15.5.4 with App Router
- **UI Library**: React 19.1.0
- **Language**: TypeScript 5
- **Styling**: CSS Modules
- **Build Tool**: Turbopack

### Backend Services
- **Language**: Go 1.25.4
- **Web Framework**: Gin (api-gateway), gRPC (authentication)
- **Database**: PostgreSQL with GORM ORM
- **Cache**: Redis
- **API Documentation**: Swagger/OpenAPI (api-gateway)

### DevOps & Infrastructure
- **Containerization**: Docker & Docker Compose
- **Development Environment**: Dev Containers (VS Code)
- **Hot Reload**: Air (Go services)
- **CI/CD**: GitHub Actions

---

## 📦 Services Overview

### API Gateway
- **Port**: 3000
- **Purpose**: Central entry point for all client requests
- **Features**:
  - Request routing and load balancing
  - Rate limiting and throttling
  - Identity verification with visitor tokens
  - Health check endpoints
  - Swagger API documentation
- **Tech**: Gin framework, Redis for token storage

### Authentication Service
- **Purpose**: User authentication and authorization
- **Features**:
  - User registration and login
  - Password hashing and verification
  - Session management
  - gRPC-based service communication
- **Tech**: GORM (PostgreSQL), gRPC

### Frontend
- **Port**: 3000 (development)
- **Purpose**: User-facing web application
- **Features**:
  - Server-side rendering with Next.js App Router
  - Optimized font loading with next/font
  - Responsive design
  - TypeScript for type safety

---

## 🛠️ Getting Started

### Prerequisites

- **Docker** & **Docker Compose** installed
- **Go** 1.25.4+ (for local development)
- **Node.js** 18+ (for frontend development)
- **PostgreSQL** & **Redis** (or use Docker)

### Quick Start with Docker Compose

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd moufube
   ```

2. **Start the API Gateway**
   ```bash
   cd deployment
   docker-compose up api-gateway
   ```

3. **Start the Frontend** (in a new terminal)
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

4. **Access the application**
   - Frontend: http://localhost:3000
   - API Gateway: http://localhost:3000
   - API Documentation: http://localhost:3000/swagger/index.html

### Development Setup with Dev Containers

Each Go service includes a `.devcontainer` configuration for a consistent development environment:

```bash
# Open API Gateway in Dev Container (VS Code)
code services/api-gateway

# Open Authentication Service in Dev Container (VS Code)
code services/authentication
```

The dev container includes:
- Go toolchain
- golangci-lint-v2
- Air for hot reload
- VS Code Go extension

### Environment Configuration

Copy the example environment files and configure as needed:

```bash
# API Gateway
cp services/api-gateway/.env.example services/api-gateway/.env

# Authentication Service
cp services/authentication/.env.example services/authentication/.env
```

#### API Gateway Environment Variables

```env
ENVIRONMENT=dev
HTTP_PORT=3000

# Timeouts
READ_TIMEOUT=990
WRITE_TIMEOUT=999
IDLE_TIMEOUT=999

# Token Configuration
SIZE_IDENTITY_TOKEN=32
VISITOR_TOKEN_EXPIRE_DAYS=365

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=admin
IDENTITY_DB=1
```

#### Authentication Service Environment Variables

```env
DB_HOST=localhost
DB_PORT=5173
DB_USER=user
DB_PASSWORD=password
DB_NAME=some_database
DB_SSL_MODE=disable
```

---

## 🧪 Development

### Frontend Development

```bash
cd frontend

# Install dependencies
npm install

# Start development server with Turbopack
npm run dev

# Build for production
npm run build

# Start production server
npm start

# Run linter
npm run lint

# Format code
npm run format
```

### Backend Services Development

#### API Gateway

```bash
cd services/api-gateway

# Download dependencies
go mod download

# Run with hot reload (Air)
air

# Run tests
go test ./...

# Run linter
golangci-lint-v2 run
```

#### Authentication Service

```bash
cd services/authentication

# Download dependencies
go mod download

# Generate proto files (if applicable)
# (protoc commands go here)

# Run with hot reload (Air)
air

# Run tests
go test ./...

# Run linter
golangci-lint-v2 run
```

### Running All Services

Use Docker Compose to orchestrate all services:

```bash
cd deployment
docker-compose up
```

---

## 📚 Documentation

### API Documentation
- **API Gateway Swagger**: Available at `/swagger/index.html` when the API Gateway is running
- **Service Documentation**: Located in each service's `documentation/` directory

### Project Structure
- **Frontend**: Next.js App Router structure under `frontend/src/app/`
- **API Gateway**: Clean architecture with `internal/` containing domain, application, and infrastructure layers
- **Authentication**: DDD-inspired structure with `domain/`, `application/`, `infrastructure/` layers

---

## 🔧 Developer Experience (DX)

### Linting & Formatting

**Frontend**:
- **Linter**: ESLint with TypeScript
- **Formatter**: Prettier
- **Pre-commit**: Husky with lint-staged

**Backend**:
- **Linter**: golangci-lint-v2
- **Formatter**: gofmt / goimports
- **Pre-commit**: Automated linting script

### Pre-commit Hooks

The project uses pre-commit hooks to ensure code quality:

```bash
# Lint all Go services
./scripts/lint.sh
```

This script runs `golangci-lint-v2` on:
- API Gateway
- Authentication Service

Failed linting will prevent commits.

### CI/CD Pipeline

GitHub Actions enforces:
1. **PR Title Validation**: Conventional commit format required
   - Format: `type(scope): description`
   - Types: `feat`, `fix`, `chore`, `docs`, `refactor`, `test`, `perf`
   - Example: `feat(api): add user authentication`

2. **Code Linting**: Runs golangci-lint-v2 on all Go services

### Dev Containers

Consistent development environments across machines:
- Pre-configured VS Code workspace
- Required extensions pre-installed
- Automatic hot reload with Air
- Integrated tooling (golangci-lint-v2, etc.)

---

## 🏗️ Architecture Decisions

### Why Monorepo?
While production typically favors polyrepos for independent deployments, a monorepo provides:
- Simplified service management during development
- Shared configuration and tooling
- Easier cross-service refactoring
- Unified CI/CD pipelines

### Clean Architecture
Services follow clean architecture principles:
- **Domain**: Business logic and entities
- **Application**: Use cases and application services
- **Infrastructure**: External dependencies (database, HTTP, etc.)
- **Interface**: Controllers, routers, and middleware

### API Gateway Pattern
Central API gateway provides:
- Single entry point for clients
- Cross-cutting concerns (auth, logging, rate limiting)
- Service abstraction and versioning
- Simplified frontend integration

---

## 🤖 AI Agents (OpenCode CLI)

This project leverages AI-powered development assistants:

1. **Documentation Agent**
   - Automated documentation generation
   - Code-to-documentation synchronization

2. **Error Summarizer & Guesser Agent**
   - Intelligent error analysis
   - Root cause suggestions
   - Debugging assistance

---

## 📊 Observability

### Logging
- **Go Services**: logrus structured logging
- **Frontend**: Console logging (extendable)

### Monitoring
- Health check endpoints on all services
- Structured logs for debugging and monitoring
- Extensible for metrics and tracing

---

## 🔐 Security

### Current Practices
- Environment-based configuration
- Pre-commit hooks preventing bad code
- Linting for security vulnerabilities
- Password hashing (authentication service)

### Future Enhancements
- Secret management with Vault
- JWT token rotation
- Rate limiting and DDoS protection
- CORS configuration
- HTTPS enforcement

---

## 🗺️ Roadmap

- [ ] Complete authentication flow
- [ ] Video upload and storage
- [ ] Video streaming functionality
- [ ] User profiles and preferences
- [ ] Comments and social features
- [ ] Recommendation engine
- [ ] Analytics dashboard
- [ ] Production deployment

---

## 🤝 Contributing

1. Follow the [Conventional Commits](https://www.conventionalcommits.org/) specification
2. Ensure all tests pass before submitting a PR
3. Run linters and fix any issues
4. Update documentation for new features
5. Use descriptive PR titles following the format: `type(scope): description`

---

## 📄 License

This project is licensed under the MIT License.

---

## 📞 Support

For questions or issues:
- Open an issue on GitHub
- Check existing documentation in service-specific READMEs
- Review API documentation at `/swagger/index.html`

---

**Built with ❤️ following enterprise-grade development practices**
