Current State Analysis

You have a monolithic approach with:

- All handlers in app.go (~378 lines)
- Basic /database folder with storage helpers
- Auth middleware separate
- No clear domain separation

Recommended Structure

Given your microservices goal and current phase, I recommend a hybrid
approach:

Phase 1: Refactor Monolith with Domain Boundaries (Now → 2 months)

/backend
/cmd
/api - main.go # Entry point with dependency wiring

    /internal                        # Private application code
      /user
        - handler.go                 # POST /register, POST /login
        - service.go                 # Registration, password hashing, KYC prep
        - repository.go              # User CRUD operations - models.go # User, UserDTO

      /auth
        - service.go                 # Token generation/validation/refresh
        - middleware.go              # AuthMiddleware (reuse existing)
        - repository.go              # Token storage (refresh tokens)
        - models.go                  # Token-related structs

      /wallet
        - handler.go                 # POST /wallets, GET /wallets/:id
        - service.go                 # Create wallet, balance checks
        - repository.go              # Wallet CRUD
        - models.go                  # Wallet struct

      /transaction
        - handler.go                 # POST /deposit, /withdraw, /transfer
        - service.go                 # Transaction orchestration
        - repository.go              # Transaction history
        - models.go                  # Transaction types
        - events.go                  # Event definitions (for future Kafka)

      /ledger
        - service.go                 # Double-entry bookkeeping logic
        - repository.go              # Ledger entries
        - models.go                  # LedgerEntry

    /pkg                             # Reusable packages
      - errors.go                    # Custom error types
      - logger.go                    # Logging utilities

    /config
      - config.go                    # Move secrets.go here

    /database
      - database.go                  # Connection setup
      /migrations                    # SQL migrations

Phase 2: Extract to Microservices (Months 3-4)

Once domains are clearly separated, extract services:

/services
/user-service
/cmd/api
/internal
/user
/auth - Dockerfile - go.mod

    /wallet-service
      /cmd/api
      /internal
        /wallet
      - Dockerfile
      - go.mod

    /transaction-service
      /cmd/api
      /internal
        /transaction
        /ledger
      - Dockerfile
      - go.mod

    /fraud-detection-service     # Python microservice
      - app.py
      - Dockerfile

    /notification-service
      - Dockerfile
      - go.mod

/shared # Shared libraries
/events # Event schemas
/proto # gRPC definitions (optional)

Migration Strategy

Step 1: Extract User Domain (Week 1)

1. Create /internal/user/models.go with User/UserDTO
2. Create /internal/user/repository.go - move database user functions
3. Create /internal/user/service.go - move registration/login logic from
   app.go:237-243
4. Create /internal/user/handler.go - move POST /login (app.go:108), POST
   /users (app.go:213)
5. Update cmd/api/main.go with dependency injection

Step 2: Extract Auth Domain (Week 1-2)

1. Move token generation from app.go:252-320 to /internal/auth/service.go
2. Move refresh logic from app.go:323-377 to auth service
3. Move middleware to /internal/auth/middleware.go
4. Create auth repository for token storage

Step 3: Add Wallet Domain (Week 2-3)

This is new functionality - start with proper structure:

- wallet/handler.go: API endpoints
- wallet/service.go: Business rules
- wallet/repository.go: DB operations

Step 4: Add Transaction Domain (Week 3-4)

- Implement with event-driven mindset (even if Kafka comes later)
- transaction/events.go defines events as structs now
- Later: wire to Kafka publisher

Key Principles

Dependency Flow:
Handler → Service → Repository → Database
↓ ↓
Response Events (future Kafka)

Why This Path?

| Decision                    | Reasoning                                         |
| --------------------------- | ------------------------------------------------- |
| Refactor monolith first     | Understand domain boundaries before splitting     |
| Domain folders in /internal | Prepares for microservice extraction              |
| Keep event definitions now  | Add Kafka integration later without restructuring |
| Share /pkg utilities        | Easy to extract to shared lib later               |

Would you like me to:

1. Start the refactor by extracting the user domain?
2. Show detailed code for one domain (user, auth, or wallet)?
3. Explain testing strategy for the layered architecture?
