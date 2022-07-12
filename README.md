# Ispsystems Vmmanager v6 JSON-RPC API in Golang

This is the Ispsystems Vmmanager v6 JSON-RPC API wrapper in Go. See a [documentation](https://docs.ispsystem.com/vmmanager-admin/developer-section/api/vmmanager-api#/) for more information.

## How to use

```go
package main

import (
    vmm "github.com/gadost/go-vmmanager-api"
    "fmt"
)

func main() {
    auth := &vmm.Auth{
        AuthByEmailAndPassword: vmm.AuthByEmailAndPassword{
            Email:    "admin@email.com",
            Password: "superPassworD",
        },
    }
    conn := vmm.New("domain.com").Auth(auth)
    resp, err := conn.ClusterList()
    fmt.Println(resp.LastNotify)
}
```

## Implemented

- Auth
- Backup
- Cluster

still in development
