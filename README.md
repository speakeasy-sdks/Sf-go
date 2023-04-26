<div align="center">
        <img src="https://user-images.githubusercontent.com/6267663/232818997-c90f68a0-39a7-466c-85f0-a9d600c5daec.png" width="100" />
        <h1>Cats Go SDK</h1>
</div>
       
<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/Sf-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
    res, err := s.AppControllerGetHello(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [SDK](docs/sdk/README.md)

* [AppControllerGetHello](docs/sdk/README.md#appcontrollergethello)
* [CatControllerCreate](docs/sdk/README.md#catcontrollercreate)
* [CatControllerFindAll](docs/sdk/README.md#catcontrollerfindall)
* [CatControllerFindOne](docs/sdk/README.md#catcontrollerfindone)
* [CatControllerRemove](docs/sdk/README.md#catcontrollerremove)
* [CatControllerUpdate](docs/sdk/README.md#catcontrollerupdate)
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
