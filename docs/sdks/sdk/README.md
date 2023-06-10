# SDK

## Overview

Cats example: The cats API description

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.AppControllerGetHelloResponse](../../models/operations/appcontrollergethelloresponse.md), error**


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
    res, err := s.SDK.CatControllerCreate(ctx, shared.CreateCatDto{
        Age: 5488.14,
        Breed: "provident",
        Name: "Ellis Mitchell",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.CreateCatDto](../../models/shared/createcatdto.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CatControllerCreateResponse](../../models/operations/catcontrollercreateresponse.md), error**


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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.CatControllerFindAllResponse](../../models/operations/catcontrollerfindallresponse.md), error**


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
    res, err := s.SDK.CatControllerFindOne(ctx, operations.CatControllerFindOneRequest{
        ID: "d69a674e-0f46-47cc-8796-ed151a05dfc2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CatControllerFindOneRequest](../../models/operations/catcontrollerfindonerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.CatControllerFindOneResponse](../../models/operations/catcontrollerfindoneresponse.md), error**


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
    res, err := s.SDK.CatControllerRemove(ctx, operations.CatControllerRemoveRequest{
        ID: "ddf7cc78-ca1b-4a92-8fc8-16742cb73920",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CatControllerRemoveRequest](../../models/operations/catcontrollerremoverequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CatControllerRemoveResponse](../../models/operations/catcontrollerremoveresponse.md), error**


## CatControllerUpdate

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/Sf-go"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/operations"
	"github.com/speakeasy-sdks/Sf-go/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.CatControllerUpdate(ctx, operations.CatControllerUpdateRequest{
        UpdateCatDto: shared.UpdateCatDto{},
        ID: "5929396f-ea75-496e-b10f-aaa2352c5955",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CatControllerUpdateRequest](../../models/operations/catcontrollerupdaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.CatControllerUpdateResponse](../../models/operations/catcontrollerupdateresponse.md), error**

