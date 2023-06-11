// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package generated

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Project defines model for Project.
type Project struct {
	Code      string             `json:"code"`
	CretaedAt openapi_types.Date `json:"cretaed_at"`
	Id        string             `json:"id"`
	Name      string             `json:"name"`
}

// User defines model for User.
type User struct {
	// CreatedAt The date that the user was created.
	CreatedAt   *openapi_types.Date `json:"created_at,omitempty"`
	DateOfBirth *openapi_types.Date `json:"date_of_birth,omitempty"`
	Email       openapi_types.Email `json:"email"`

	// EmailVerified Set to true if the user's email has been verified.
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"first_name"`

	// Id Unique identifier for the given user.
	Id       int    `json:"id"`
	LastName string `json:"last_name"`
}

// PostUserJSONBody defines parameters for PostUser.
type PostUserJSONBody struct {
	DateOfBirth openapi_types.Date `json:"date_of_birth"`
	Email       string             `json:"email"`
	FirstName   string             `json:"first_name"`
	LastName    string             `json:"last_name"`
}

// PatchUsersUserIdJSONBody defines parameters for PatchUsersUserId.
type PatchUsersUserIdJSONBody struct {
	DateOfBirth *string `json:"dateOfBirth,omitempty"`

	// Email If a new email is given, the user's email verified property will be set to false.
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody PostUserJSONBody

// PatchUsersUserIdJSONRequestBody defines body for PatchUsersUserId for application/json ContentType.
type PatchUsersUserIdJSONRequestBody PatchUsersUserIdJSONBody
