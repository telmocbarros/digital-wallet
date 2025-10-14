# Ledger Service API Examples

The ledger service provides read-only HTTP endpoints for querying balances, transaction history, and verifying integrity.

## Endpoints

### 1. Get Account Balance

Retrieves the current balance for an account.

```bash
GET /api/ledger/balance/:accountId
```

**Example:**
```bash
curl -X GET http://localhost:8080/api/ledger/balance/alice-wallet-123 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "account_id": "alice-wallet-123",
  "balance": {
    "account_id": "alice-wallet-123",
    "balance": 50.00,
    "currency": "USD",
    "updated_at": 1697299200
  }
}
```

---

### 2. Get Account Statement

Retrieves all ledger entries for an account (transaction history).

```bash
GET /api/ledger/statement/:accountId
```

**Example:**
```bash
curl -X GET http://localhost:8080/api/ledger/statement/alice-wallet-123 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "account_id": "alice-wallet-123",
  "count": 3,
  "entries": [
    {
      "id": "entry-001",
      "account_id": "alice-wallet-123",
      "amount": 100.00,
      "currency": "USD",
      "entry_type": "CREDIT",
      "transaction_id": "txn-001",
      "transaction_type": "DEPOSIT",
      "created_at": 1697299100,
      "description": "Deposit from external_bank: Initial deposit"
    },
    {
      "id": "entry-002",
      "account_id": "alice-wallet-123",
      "amount": -50.00,
      "currency": "USD",
      "entry_type": "DEBIT",
      "transaction_id": "txn-002",
      "transaction_type": "TRANSFER",
      "created_at": 1697299200,
      "description": "Transfer to bob-wallet-456: Payment for services"
    }
  ]
}
```

---

### 3. Get Transaction Details

Retrieves all ledger entries for a specific transaction (shows double-entry breakdown).

```bash
GET /api/ledger/transaction/:transactionId
```

**Example:**
```bash
curl -X GET http://localhost:8080/api/ledger/transaction/txn-002 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "transaction_id": "txn-002",
  "count": 2,
  "entries": [
    {
      "id": "entry-002",
      "account_id": "alice-wallet-123",
      "amount": -50.00,
      "currency": "USD",
      "entry_type": "DEBIT",
      "transaction_id": "txn-002",
      "transaction_type": "TRANSFER",
      "created_at": 1697299200,
      "description": "Transfer to bob-wallet-456: Payment for services"
    },
    {
      "id": "entry-003",
      "account_id": "bob-wallet-456",
      "amount": 50.00,
      "currency": "USD",
      "entry_type": "CREDIT",
      "transaction_id": "txn-002",
      "transaction_type": "TRANSFER",
      "created_at": 1697299200,
      "description": "Transfer from alice-wallet-123: Payment for services"
    }
  ]
}
```

---

### 4. Verify Account Balance (Admin/Debug)

Verifies that the cached balance matches the calculated balance from ledger entries.

```bash
POST /api/ledger/verify/account/:accountId
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/ledger/verify/account/alice-wallet-123 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response (Success):**
```json
{
  "account_id": "alice-wallet-123",
  "verified": true,
  "message": "Account balance verified successfully"
}
```

**Response (Mismatch):**
```json
{
  "account_id": "alice-wallet-123",
  "verified": false,
  "message": "Cached balance does not match calculated balance"
}
```

---

### 5. Verify Transaction (Admin/Debug)

Verifies that all entries for a transaction sum to zero (double-entry integrity check).

```bash
POST /api/ledger/verify/transaction/:transactionId
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/ledger/verify/transaction/txn-002 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response (Success):**
```json
{
  "transaction_id": "txn-002",
  "verified": true,
  "message": "Transaction verified successfully"
}
```

**Response (Failed):**
```json
{
  "transaction_id": "txn-002",
  "verified": false,
  "message": "Transaction entries do not sum to zero"
}
```

---

## Integration with Transaction Service (Event-Driven)

**Note:** Write operations (transfers, deposits, withdrawals) are NOT exposed via HTTP. Instead, they are triggered by events from the Transaction Service:

```
Transaction Service → Kafka Event → Ledger Service
```

Example flow:
1. User initiates transfer via Transaction Service
2. Transaction Service emits `TransactionInitiated` event
3. Ledger Service consumes event and creates ledger entries
4. Ledger Service emits `TransactionCompleted` event
5. Transaction Service updates transaction status

## Use Cases

### For User Dashboards
- **Get Balance**: Show user's current wallet balance
- **Get Statement**: Display transaction history / account activity

### For Admin Panels
- **Verify Account**: Run integrity checks on user accounts
- **Verify Transaction**: Audit specific transactions
- **Get Transaction Details**: Debug transaction issues

### For Reporting
- **Get Statement**: Generate account statements for compliance
- **Get Transaction Details**: Analyze transaction patterns

## Error Responses

### 400 Bad Request
```json
{
  "error": "Missing required parameter: accountId"
}
```

### 404 Not Found
```json
{
  "error": "Account not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```
