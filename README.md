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

## Extended token lifetime

```go
// Standart token lifetime = 1h inactivity . If you need extended tokenlife time:
expiresAt := time.Now().Add(30*time.Day)
conn := vmm.New("domain.com").Auth(auth).WithExtendedTokenLifetime(expiresAt, "extended token desc")
```

## Errors handling

```go
resp, err := conn.Hosts(1,s)
/* related API errors: resp contains field Error 
type Error struct {
    Code int    `json:"code,omitempty"`
    Msg  string `json:"msg,omitempty"`
}
*/
fmt.Println(resp.Error.Msg)

// Http errors : err
fmt.Println(err)
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
- Platform Backups
- Preset
- Recipe
- Repository
- Snapshot
- Schedule
- Storage
- Settings
- Tag
- Task
- VXLan
- User
- VMmmanager
