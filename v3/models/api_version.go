// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// APIVersion API Version
//
// API Version of the Nutanix v3 API framework.
//
// swagger:model api_version
type APIVersion string

// Validate validates this api version
func (m APIVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this api version based on the context it is used
func (m APIVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", APIVersion(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
