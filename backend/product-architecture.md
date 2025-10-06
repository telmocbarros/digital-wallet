# Product Architecture Guide

## Overview

This document outlines architectural patterns for structuring the Digital Wallet backend application in Go.

---

## Architectural Pattern Comparison

| Pattern | Best For | Complexity | Testability | Scalability |
|---------|----------|------------|-------------|-------------|
| **Controller-Service-Repository** | Teams familiar with MVC, medium-sized apps | Medium | High | Medium |
| **Domain-Driven Design (DDD)** | Complex business domains, large applications | High | High | High |
| **Flat Standard Go Layout** | Simple APIs, Go-idiomatic projects | Low | Medium | Medium |

---

## Pattern Details

### 1. Controller-Service-Repository (MVC-ish)

```
/handlers      - HTTP handlers (controllers)
/services      - Business logic
/repositories  - Database access
/models        - Data structures
/middleware    - Auth, logging, etc.
```

#### Pros & Cons

| Pros | Cons |
|------|------|
| ✅ Clear separation of concerns | ❌ Can feel over-engineered for simple apps |
| ✅ Easy to test each layer | ❌ Lots of boilerplate |
| ✅ Familiar to developers from other languages | ❌ All handlers/services mixed together |
| ✅ Well-understood pattern | ❌ Hard to locate feature-specific code |

---

### 2. Domain-Driven Design (DDD)

```
/user
  - handler.go
  - service.go
  - repository.go
  - model.go
/wallet
  - handler.go
  - service.go
  - repository.go
  - model.go
/transaction
  - handler.go
  - service.go
  - repository.go
  - model.go
```

#### Pros & Cons

| Pros | Cons |
|------|------|
| ✅ Everything related to a domain is in one place | ❌ Shared code location can be unclear |
| ✅ Scales well as features grow | ❌ More upfront planning required |
| ✅ Reflects business domains clearly | ❌ Can have code duplication across domains |
| ✅ Easy to find feature-specific code | ❌ Steeper learning curve |

---

### 3. Flat "Standard Go Layout"

```
/cmd           - Main applications
/internal      - Private application code
  /api         - HTTP handlers
  /service     - Business logic
  /store       - Database
/pkg           - Public libraries (reusable)
```

#### Pros & Cons

| Pros | Cons |
|------|------|
| ✅ Idiomatic Go | ❌ Can feel odd if new to Go |
| ✅ Clear public vs private boundaries | ❌ All business logic in one folder |
| ✅ Standard in Go community | ❌ Harder to navigate large codebases |
| ✅ Minimal boilerplate | ❌ Services can become monolithic |

---

## Recommended Architecture: Domain-Driven + Controller-Service-Repository Hybrid

For the Digital Wallet project, we recommend combining DDD with layered architecture:

```
/backend
  /cmd
    /api
      - main.go                    # Application entry point

  /internal
    /user
      - handler.go                 # HTTP endpoints (POST /register, POST /login)
      - service.go                 # Business logic (registration, KYC validation)
      - repository.go              # Database queries (GetByEmail, Create)
      - models.go                  # User, UserDTO structs

    /wallet
      - handler.go                 # HTTP endpoints (POST /wallets, GET /wallets/:id)
      - service.go                 # Business logic (create wallet, check balance)
      - repository.go              # Database queries
      - models.go                  # Wallet struct

    /transaction
      - handler.go                 # HTTP endpoints (POST /deposit, POST /withdraw, POST /transfer)
      - service.go                 # Business logic (validate, process transactions)
      - repository.go              # Database queries
      - models.go                  # Transaction struct
      - events.go                  # Event definitions (TransactionInitiated, etc.)

    /ledger
      - service.go                 # Double-entry bookkeeping logic
      - repository.go              # Ledger entry queries
      - models.go                  # LedgerEntry struct

    /auth
      - middleware.go              # AuthMiddleware
      - service.go                 # Token generation, validation, refresh
      - models.go                  # Token-related structs

  /config
    - config.go                    # Configuration, secrets, env vars

  /pkg
    - errors.go                    # Custom error types
    - events.go                    # Event bus interface
    - logger.go                    # Logging utilities

  /database
    - migrations/                  # SQL migration files
```

### Why This Structure?

| Benefit | Explanation |
|---------|-------------|
| **Clear Domain Boundaries** | Wallet code doesn't leak into user code; easy to reason about |
| **Service Layer for Business Rules** | "Can't withdraw more than balance" lives in `transaction/service.go` |
| **Repository Isolation** | Easy to swap databases or mock for testing |
| **Testable** | Mock services in handlers, mock repositories in services |
| **Event-Driven Ready** | `transaction/events.go` can publish to Kafka/message bus |
| **Scales with Complexity** | Each domain can grow independently |

---

## Layer Responsibilities

### Handler Layer (Controllers)

| Responsibility | What It Does | What It Doesn't Do |
|----------------|--------------|-------------------|
| Request parsing | Extract data from HTTP request | Business validation |
| Input validation | Check required fields present | Complex business rules |
| Response formatting | Convert service results to JSON | Database queries |
| Error handling | Map service errors to HTTP codes | Execute business logic |

**Example:**
```go
func (h *Handler) Register(c *gin.Context) {
    var req RegisterRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    user, err := h.service.Register(req.Email, req.Password)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, user)
}
```

---

### Service Layer (Business Logic)

| Responsibility | What It Does | What It Doesn't Do |
|----------------|--------------|-------------------|
| Business rules | Enforce constraints, policies | HTTP concerns |
| Orchestration | Coordinate multiple repositories | Parse HTTP requests |
| Transactions | Manage database transactions | Know about JSON/HTTP |
| Event publishing | Emit domain events | SQL queries |

**Example:**
```go
func (s *Service) Transfer(fromWalletID, toWalletID string, amount float64) error {
    // Business rule: Check sufficient balance
    fromWallet, _ := s.walletRepo.GetByID(fromWalletID)
    if fromWallet.Balance < amount {
        return ErrInsufficientFunds
    }

    // Orchestrate multiple operations
    s.walletRepo.Debit(fromWalletID, amount)
    s.walletRepo.Credit(toWalletID, amount)
    s.ledgerService.RecordTransfer(fromWalletID, toWalletID, amount)

    // Publish event
    s.eventBus.Publish(TransactionCompleted{...})

    return nil
}
```

---

### Repository Layer (Data Access)

| Responsibility | What It Does | What It Doesn't Do |
|----------------|--------------|-------------------|
| CRUD operations | Execute SQL queries | Business validation |
| Data mapping | Convert DB rows to structs | Complex calculations |
| Query building | Construct safe SQL statements | Enforce business rules |
| Connection management | Handle DB connections | Know about HTTP |

**Example:**
```go
func (r *Repository) GetByID(id string) (*User, error) {
    var user User
    err := r.db.QueryRow(
        "SELECT id, email FROM users WHERE id = $1",
        id,
    ).Scan(&user.ID, &user.Email)

    if err == sql.ErrNoRows {
        return nil, ErrUserNotFound
    }

    return &user, err
}
```

---

## Migration Path

### Current State

```
/backend
  app.go                           # ⚠️ Handlers + business logic mixed
  /database                        # Repository-ish code
  /middlewares                     # Auth middleware
  /config                          # Secrets
```

### Migration Steps

| Step | Action | Files Affected |
|------|--------|----------------|
| 1 | Create domain folders | `/internal/user`, `/internal/auth`, `/internal/wallet` |
| 2 | Extract user handlers | Move login/register from `app.go` to `user/handler.go` |
| 3 | Create user service | Move business logic to `user/service.go` |
| 4 | Refactor database code | Move to `user/repository.go` |
| 5 | Update main.go | Wire dependencies (repos → services → handlers) |
| 6 | Repeat for wallet domain | Create wallet domain structure |
| 7 | Add transaction domain | Implement transaction service with events |

---

## Example: User Domain Structure

### File: `internal/user/models.go`
```go
package user

type User struct {
    ID       string
    Email    string
    Password string // hashed
}

type UserDTO struct {
    ID    string `json:"id"`
    Email string `json:"email"`
}
```

### File: `internal/user/repository.go`
```go
package user

type Repository interface {
    GetByEmail(email string) (*User, error)
    GetByID(id string) (*User, error)
    Create(user *User) error
}

type postgresRepository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &postgresRepository{db: db}
}

func (r *postgresRepository) GetByEmail(email string) (*User, error) {
    // SQL query implementation
}
```

### File: `internal/user/service.go`
```go
package user

type Service struct {
    repo Repository
}

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) Register(email, password string) (*UserDTO, error) {
    // Check if user exists
    _, err := s.repo.GetByEmail(email)
    if err == nil {
        return nil, ErrUserAlreadyExists
    }

    // Hash password
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

    // Create user
    user := &User{
        ID:       uuid.New().String(),
        Email:    email,
        Password: string(hashedPassword),
    }

    if err := s.repo.Create(user); err != nil {
        return nil, err
    }

    return &UserDTO{ID: user.ID, Email: user.Email}, nil
}

func (s *Service) Login(email, password string) (string, string, error) {
    // Get user
    user, err := s.repo.GetByEmail(email)
    if err != nil {
        return "", "", ErrInvalidCredentials
    }

    // Verify password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", "", ErrInvalidCredentials
    }

    // Generate tokens (could be delegated to auth service)
    accessToken := generateAccessToken(user)
    refreshToken := generateRefreshToken(user)

    return accessToken, refreshToken, nil
}
```

### File: `internal/user/handler.go`
```go
package user

import "github.com/gin-gonic/gin"

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) Register(c *gin.Context) {
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    user, err := h.service.Register(req.Email, req.Password)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, user)
}

func (h *Handler) Login(c *gin.Context) {
    email := c.Query("email")
    password := c.Query("password")

    accessToken, refreshToken, err := h.service.Login(email, password)
    if err != nil {
        c.JSON(401, gin.H{"error": "Invalid credentials"})
        return
    }

    // Set cookies
    c.SetCookie("access_token", accessToken, 300, "/", "", false, true)
    c.SetCookie("refresh_token", refreshToken, 6400, "/", "", false, true)

    c.JSON(200, gin.H{"message": "Login successful"})
}
```

### File: `cmd/api/main.go`
```go
package main

import (
    "digitalwallet/backend/internal/user"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Initialize dependencies
    db := initDatabase()
    userRepo := user.NewRepository(db)
    userService := user.NewService(userRepo)
    userHandler := user.NewHandler(userService)

    // Register routes
    r.POST("/register", userHandler.Register)
    r.POST("/login", userHandler.Login)

    r.Run()
}
```

---

## Testing Strategy by Layer

| Layer | Test Type | What to Test | How to Mock |
|-------|-----------|--------------|-------------|
| **Handler** | Integration | HTTP request/response flow | Mock service layer |
| **Service** | Unit | Business logic, rules | Mock repository layer |
| **Repository** | Integration | Database queries | Use test database or mocks |

### Example: Testing Service Layer
```go
func TestService_Register(t *testing.T) {
    // Mock repository
    mockRepo := &MockRepository{
        GetByEmailFunc: func(email string) (*User, error) {
            return nil, ErrNotFound // User doesn't exist
        },
        CreateFunc: func(user *User) error {
            return nil // Success
        },
    }

    service := NewService(mockRepo)

    // Test
    user, err := service.Register("test@example.com", "password123")

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "test@example.com", user.Email)
}
```

---

## When to Refactor

| Current State | Action | Priority |
|---------------|--------|----------|
| Building initial prototype | Keep flat structure | Low |
| 3+ domains (user, wallet, transaction) | Start domain separation | Medium |
| Handlers >200 lines | Extract to services | High |
| Hard to test | Add service/repository layers | High |
| Multiple developers | Enforce architecture | High |

---

## Key Takeaways

1. **Start Simple**: Don't over-architect early. Flat structure is fine initially.
2. **Refactor When Needed**: When handlers grow or testing gets hard, refactor.
3. **Domain-Driven Works Here**: Your project has clear domains (user, wallet, transaction, ledger).
4. **Layer Separation**: Keep HTTP concerns in handlers, business logic in services, data access in repositories.
5. **Testability**: Each layer should be independently testable with mocked dependencies.
