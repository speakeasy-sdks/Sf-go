# SDK

## Overview

The cats API description

### Available Operations

* [AppControllerGetHello](#appcontrollergethello)
* [CatControllerCreate](#catcontrollercreate)
* [CatControllerFindAll](#catcontrollerfindall)
* [CatControllerFindOne](#catcontrollerfindone)
* [CatControllerRemove](#catcontrollerremove)
* [CatControllerUpdate](#catcontrollerupdate)

## AppControllerGetHello

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.AppControllerGetHello(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CatControllerCreate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := shared.CreateCatDto{
        Age: 5488.14,
        Breed: "provident",
        Name: "Ellis Mitchell",
    }

    res, err := s.SDK.CatControllerCreate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CatControllerFindAll

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.CatControllerFindAll(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CatControllerFindOne

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := operations.CatControllerFindOneRequest{
        ID: "d69a674e-0f46-47cc-8796-ed151a05dfc2",
    }

    res, err := s.SDK.CatControllerFindOne(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CatControllerRemove

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := operations.CatControllerRemoveRequest{
        ID: "ddf7cc78-ca1b-4a92-8fc8-16742cb73920",
    }

    res, err := s.SDK.CatControllerRemove(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CatControllerUpdate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := operations.CatControllerUpdateRequest{
        RequestBody: map[string]interface{}{
            "natus": "sed",
            "iste": "dolor",
        },
        ID: "96fea759-6eb1-40fa-aa23-52c5955907af",
    }

    res, err := s.SDK.CatControllerUpdate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
