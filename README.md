# Go CSRF Token Generator and Validator 
[![GoDoc](https://godoc.org/github.com/jor-go/csrf?status.svg)](https://godoc.org/github.com/jor-go/csrf)

### Import
``` golang
import "github.com/jor-go/csrf"
```

### Usage
``` golang
package main

import (
	"github.com/jor-go/csrf"
	"fmt"
	"time"
)


func main() {
	// Set Secret and Token Duration
	csrf.Secret = "cool-secret"
	csrf.MaxTokenAge = 12 * time.Hour
	
	// Get Session ID
	sessionID := "1234"

	// Generate Token
	token := csrf.CreateToken(sessionID)
	fmt.Println(token)

	// Validate Token
	isValid := csrf.ValidToken(token, sessionID)
	fmt.Println("Token is Valid:", isValid)
}
```
