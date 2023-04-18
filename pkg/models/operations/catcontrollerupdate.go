// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type CatControllerUpdateRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	ID          string                 `pathParam:"style=simple,explode=false,name=id"`
}

type CatControllerUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
