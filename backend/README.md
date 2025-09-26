# Set up JWT token with GoLang

https://www.youtube.com/watch?v=XFxm6NEWakQ&t=621s
https://github.com/golang-jwt/jwt

# Parsing the incoming body of a request as JSON

https://stackoverflow.com/questions/61919830/go-gin-get-request-body-json

# GoLang "if" pattern

The pattern on line 34 is a Go idiom for error handling with immediate assignment and checking.

Here's what's happening:

```go
if err := c.BindJSON(&user); err != nil {
return
}
```

Breaking it down:

1. := (short variable declaration): Creates and assigns the err variable in one step
2. c.BindJSON(&user): Attempts to parse the JSON request body into the user struct
3. err != nil: Checks if an error occurred
4. return: Exits the function early if there was an error

This pattern combines:

- Variable declaration
- Function call
- Error checking
- Early return

It's equivalent to writing:

```go
err := c.BindJSON(&user)
if err != nil {
return
}
```

Why this pattern is common in Go:

- Concise: Handles the common "call function, check error" pattern in one line
- Scoped: The err variable only exists within the if block
- Idiomatic: Very standard Go style for error handling

Note: Your current code just returns without sending an error response. You might want to add:

```go
if err := c.BindJSON(&user); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
return
}
```

# Read environment variables

```go
import (
    "os"
    "fmt"
)

var ENV_VARIABLE = os.GetEnv("my-env-variable")
fmt.Println(ENV_VARIABLE)
```
