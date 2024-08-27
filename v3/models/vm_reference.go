// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMReference Reference to a vm
//
// # The reference to a vm
//
// swagger:model vm_reference
type VMReference struct {

	// The kind name
	// Required: true
	// Read Only: true
	Kind string `json:"kind"`

	// name
	// Read Only: true
	// Max Length: 1024
	Name string `json:"name,omitempty"`

	// uuid
	// Required: true
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	UUID *string `json:"uuid"`
}

// Validate validates this vm reference
func (m *VMReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMReference) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *VMReference) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 1024); err != nil {
		return err
	}

	return nil
}

func (m *VMReference) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	if err := validate.Pattern("uuid", "body", *m.UUID, `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vm reference based on the context it is used
func (m *VMReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMReference) contextValidateKind(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "kind", "body", string(m.Kind)); err != nil {
		return err
	}

	return nil
}

func (m *VMReference) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMReference) UnmarshalBinary(b []byte) error {
	var res VMReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
