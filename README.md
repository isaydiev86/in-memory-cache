Go in-cache-memory
================================

Go Example of caching local data

See it in action:

## Example #1

```go
package main

import (
  "fmt"
  "github.com/isaydiev86/in-cache-memory"
)

func main() {
	a := NewCache()
	a.Set("userId", 32)
	userId := a.Get("userId")
	fmt.Println(userId)
	a.Delete("userId")
	userId = a.Get("userId")
	fmt.Println(userId)
}
