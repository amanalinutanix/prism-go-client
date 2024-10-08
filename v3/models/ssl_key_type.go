// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// SslKeyType SSL key type
//
// SSL key type. Key types with RSA_2048, ECDSA_256, ECDSA_384 and ECDSA_521
// are supported for key generation and importing.
//
// swagger:model ssl_key_type
type SslKeyType string

// Validate validates this ssl key type
func (m SslKeyType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ssl key type based on context it is used
func (m SslKeyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
