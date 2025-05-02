# Backend Development Guidelines

## Technology Stack

### Language: Golang
- Version: Latest stable version (Recommended 1.24 or later)
- Features:
  - High performance
  - Static typing
  - Concurrency support
  - Compile-time type safety

### HTTP Server Framework: Gin
- Reasons for selection:
  - Fast routing
  - Lightweight and flexible
  - Middleware support
  - Excellent performance

### ORM: Gorm
- Reasons for selection:
  - Rich feature set
  - Database abstraction
  - Migration support
  - Easy implementation of complex queries

## Onion Architecture

### Folder Structure

```
backend/
│
├── cmd/                  # Application entry points
│   └── server/
│       └── main.go
│
├── internal/             # Internal packages (not accessible from outside)
│   ├── domain/           # Domain models and business logic
│   │   ├── models/       # Entity definitions
│   │   └── interfaces/   # Interface definitions
│   │
│   ├── infrastructure/   # External connection implementations
│   │   ├── database/     # Database connections
│   │   ├── repositories/ # Data access layer
│   │   ├── repomodels/   # DAO implementation
│   │   └── api/          # External API connections
│   │
│   ├── application/      # Application services
│   │   ├── services/     # Use case implementations
│   │   └── dto/          # Data transfer objects
│   │
│   └── presentation/     # HTTP handlers
│       ├── handlers/     # Request processing
│       └── middleware/   # Common middleware
│
├── pkg/                  # Common packages accessible from outside
│   └── utils/            # Generic utilities
│
├── configs/              # Configuration files
│
└── migrations/           # Database migrations
```

### Architecture Advantages
- Separation of concerns
- Improved testability
- Modularity
- Clear dependency control

## Development Guidelines

1. Dependency Direction
   - Inner layers must not depend on outer layers
   - Domain layer is placed in the innermost layer and remains independent of other layers

2. Interface-Driven Development
   - Depend on interfaces rather than concrete implementations
   - Adhere to the Dependency Inversion Principle

3. Error Handling
   - Utilize custom error types
   - Wrap errors and provide detailed information

4. Security
   - Input validation
   - Appropriate authentication and authorization mechanisms
   - Secure database operations

5. Where to start implementation
   - Try to proceed with implementation from the less dependent modules

6. ALways create unit test
   - Make sure to write unit test after implementation

7. ALways create integration test for apis
   - Make sure to create integration test after implementing an api

8. Run commands inside docker container
    - Make sure to run commands (like installing package) inside docker compose's api container

9. Implement repository models to correspond with tables
   - let repository handle repository models, and use it to create record, select record, create domain models ...

## Local Development

### Hot Reloading
- We use [Reflex](https://github.com/cespare/reflex) for hot reloading during development
- Reflex automatically restarts the Go application when source files change
- To use reflex in development:
  ```bash
  # Inside the API docker container
  reflex -r '\.go$' -s -- sh -c 'go run /app/cmd/server/main.go'
  ```
- Configuration can be customized to watch specific file patterns or exclude certain directories

## Architectural points

- API path design
   - Please make sure to treat `docs/openapi.yml` spec file as an absolute truth
   - Make them based on the idea of restful for consistent api path rules.
   - Make sure to prefix the path with `/api/`
- ID for domain models
   * Use uuid v4 for the identifier of domain models