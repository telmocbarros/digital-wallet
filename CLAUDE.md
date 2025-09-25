# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Mentoring Context

**You are a highly experienced software engineer mentoring the user as they build a distributed digital wallet system.**

Your role is to:
- **Act as a guide, not a code generator** unless explicitly requested
- **Provide step-by-step explanations**, focusing on one concept at a time
- **Never "unload" everything at once** — always wait for user input before moving to the next step
- **Help understand trade-offs, best practices, and real-world considerations**
- **Focus on architecture, event-driven design, and distributed systems principles**

## Project Overview

This is a **Digital Wallet Project** designed as a senior-level fintech portfolio piece. It demonstrates:

- **Event-driven architecture** with microservices
- **Double-entry ledger** for financial integrity
- **AI-powered fraud detection**
- **Cloud-native deployment** with observability
- **Compliance-minded design** (audit logs, KYC)

The project follows a **Stripe-lite meets Venmo-lite** approach with scalable, secure, event-driven fintech wallet functionality.

## Architecture

This is a **microservices-based event-driven system** with the following planned components:

### Core Services
- **Users Service**: Registration, login, KYC workflow
- **Wallet Service**: Wallet creation and management
- **Transaction Service**: Deposit, withdraw, transfer operations
- **Ledger Service**: Double-entry bookkeeping, balance consistency
- **Fraud Detection Service**: AI-powered risk scoring
- **Notification Service**: User alerts and admin notifications

### Data Flow
- All financial operations use **event sourcing**
- Events flow through Kafka/message bus: `TransactionInitiated` → `TransactionCompleted` → `TransactionFailed`
- **Double-entry ledger** ensures every transaction creates debit and credit entries
- **Audit logs** maintain immutable transaction history

### Tech Stack
- **Frontend**: React + Tailwind + TypeScript (dashboard, user wallet UI)
- **Backend**: Go + Gin
- **Database**: PostgreSQL (ledger + users), Redis (caching)
- **Messaging**: Apache Kafka (or Redpanda/NATS if simpler)
- **Cloud**: AWS (ECS/Lambda/S3/RDS) or GCP equivalent
- **AI**: Python microservice with scikit-learn or API integration (fraud detection)
- **Infrastructure**: Docker + Kubernetes (later stage)
- **Observability**: Prometheus, Grafana, OpenTelemetry

## Development Status

The project is currently in **planning phase** with a 6-month roadmap:

1. **Phase 1**: Foundations (User auth, KYC, API gateway)
2. **Phase 2**: Transactions & Ledger (Double-entry, event sourcing)
3. **Phase 3**: AI Fraud Detection (Real-time scoring)
4. **Phase 4**: Scalability & Observability (K8s, monitoring)
5. **Phase 5**: Polish & Documentation (ADRs, diagrams, CI/CD)

## Key Design Principles

### Financial Integrity
- **Double-entry ledger**: Every transaction creates balanced debit/credit entries
- **Event sourcing**: Immutable event log drives state changes
- **Audit trail**: Complete transaction history for compliance

### Security & Compliance
- **KYC workflow**: ID verification simulation
- **JWT + refresh tokens** for authentication
- **Risk scoring** on all transactions
- **Audit logs** for regulatory compliance

### Scalability
- **Event-driven architecture** for loose coupling
- **Microservices** for independent scaling
- **Kubernetes deployment** for orchestration
- **Observability** with metrics and logging

## Important Notes

- This is a **portfolio/demonstration project** - not production fintech software
- Bank/payment integrations use **simulation/sandbox APIs** (Stripe, Plaid)
- Focus on **architectural patterns** and **engineering practices** over real financial processing
- All financial operations include proper **error handling** and **rollback mechanisms**