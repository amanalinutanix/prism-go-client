// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CitrixConnectorConfigDetailsSpec Citrix Connector details.
//
// Citrix Connector details.
//
// swagger:model citrix_connector_config_details_spec
type CitrixConnectorConfigDetailsSpec struct {

	// Reference to the list of vm ids registered with citrix cloud.
	CitrixVMReferenceList []*VMReference `json:"citrix_vm_reference_list"`

	// The client id for the Citrix Cloud.
	// Max Length: 200
	ClientID string `json:"client_id,omitempty"`

	// The client secret for the Citrix Cloud.
	ClientSecret string `json:"client_secret,omitempty"`

	// The customer id registered with Citrix Cloud.
	// Max Length: 200
	CustomerID string `json:"customer_id,omitempty"`

	// resource location
	ResourceLocation *CitrixResourceLocationSpec `json:"resource_location,omitempty"`
}

// Validate validates this citrix connector config details spec
func (m *CitrixConnectorConfigDetailsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCitrixVMReferenceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) validateCitrixVMReferenceList(formats strfmt.Registry) error {
	if swag.IsZero(m.CitrixVMReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.CitrixVMReferenceList); i++ {
		if swag.IsZero(m.CitrixVMReferenceList[i]) { // not required
			continue
		}

		if m.CitrixVMReferenceList[i] != nil {
			if err := m.CitrixVMReferenceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("citrix_vm_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("citrix_vm_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) validateClientID(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientID) { // not required
		return nil
	}

	if err := validate.MaxLength("client_id", "body", m.ClientID, 200); err != nil {
		return err
	}

	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) validateCustomerID(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerID) { // not required
		return nil
	}

	if err := validate.MaxLength("customer_id", "body", m.CustomerID, 200); err != nil {
		return err
	}

	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) validateResourceLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceLocation) { // not required
		return nil
	}

	if m.ResourceLocation != nil {
		if err := m.ResourceLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this citrix connector config details spec based on the context it is used
func (m *CitrixConnectorConfigDetailsSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCitrixVMReferenceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) contextValidateCitrixVMReferenceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CitrixVMReferenceList); i++ {

		if m.CitrixVMReferenceList[i] != nil {

			if swag.IsZero(m.CitrixVMReferenceList[i]) { // not required
				return nil
			}

			if err := m.CitrixVMReferenceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("citrix_vm_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("citrix_vm_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CitrixConnectorConfigDetailsSpec) contextValidateResourceLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceLocation != nil {

		if swag.IsZero(m.ResourceLocation) { // not required
			return nil
		}

		if err := m.ResourceLocation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource_location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CitrixConnectorConfigDetailsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CitrixConnectorConfigDetailsSpec) UnmarshalBinary(b []byte) error {
	var res CitrixConnectorConfigDetailsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
