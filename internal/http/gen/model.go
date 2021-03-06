// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package gen

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// User defines model for user.
type User struct {
	Id       *string `json:"id,omitempty"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
}

// LoginJSONBody defines parameters for Login.
type LoginJSONBody User

// SignupJSONBody defines parameters for Signup.
type SignupJSONBody User

// LoginJSONRequestBody defines body for Login for application/json ContentType.
type LoginJSONRequestBody LoginJSONBody

// SignupJSONRequestBody defines body for Signup for application/json ContentType.
type SignupJSONRequestBody SignupJSONBody
