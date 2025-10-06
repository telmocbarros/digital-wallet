# Security settings

### 1. Cookie Security Flags (app.go:200)

Currently: c.SetCookie(..., false, true) - secure: false

- Should be true in production (requires HTTPS)
- Consider adding SameSite=Strict to prevent CSRF attacks

### 2. CORS Configuration (app.go:38)

Currently hardcoded to localhost:5173

- Fine for development
- In production: use environment variable for allowed origins

### 3. Error Information Leakage

Your errors are good, but watch for:

- Don't reveal "email exists" vs "wrong password" (helps attackers enumerate users)
- Currently you return "User not found" which is fine for a portfolio project

### 4. Token Secret Strength

Make sure ACCESS_TOKEN_SECRET and REFRESH_TOKEN_SECRET are:

- Long (32+ characters)
- Random
- Different from each other
- Stored securely (env vars, not hardcoded)

# Authorization (Beyond Roles)

### 5. Resource Ownership Checks

When you build wallet/transaction endpoints, ensure:
// User can only access THEIR wallet
userId := c.GetString("userId")
if wallet.OwnerId != userId {
c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
return
}

### 6. Action-Based Authorization

Even without roles, you'll need checks like:

- Can user withdraw if balance is insufficient?
- Can user transfer to themselves?
- Is the transaction amount valid?

---

Architecture Considerations

### 7. Audit Logging

For financial operations, track:

- Who did what, when
- IP addresses
- Failed authorization attempts

### 8. Token Blacklisting (Optional)

Currently when user logs out, access token is still valid until expiry (5 min)

- For high-security: maintain a blacklist of revoked access tokens
- Trade-off: adds complexity and DB lookups on every request
