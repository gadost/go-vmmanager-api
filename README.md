# Ispsystems Vmmanager v6 JSON-RPC API in Golang

This is the Ispsystems Vmmanager v6 JSON-RPC API wrapper in Go. See a [documentation](https://docs.ispsystem.com/vmmanager-admin/developer-section/api/vmmanager-api#/) for more information.

## How to use

```go
package main

import (
    vmm "github.com/gadost/go-vmmanager-api"
    "github.com/gadost/go-vmmanager-api/types"
    "fmt"
)

func main() {
    // Define auth params
    auth := &vmm.Auth{
        AuthByEmailAndPassword: vmm.AuthByEmailAndPassword{
            Email:    "admin@email.com",
            Password: "superPassworD",
        },
    }

    // Connect to VMmanager6
    conn := vmm.New("domain.com").Auth(auth)

    // Get Cluster list
    resp, err := conn.ClusterList()
    fmt.Println(resp.LastNotify)

    //Migrate host
    s := &types.HostMigrateRequest{
        Plain: false,
        Node:  0,
    }

    resp2, err := conn.HostMigrate(1,s)
    if err != nil {
        fmt.Println(resp2.Task)
    }
}
```

## Implemented

- Auth
- Backup
- Cluster
- Disk
- Form
- Host
- Image
- Import
- License
- Network
- Node
- Script
- OS
still in development
