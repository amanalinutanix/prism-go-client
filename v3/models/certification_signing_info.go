// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificationSigningInfo Customer information in Certificate Signing Request
//
// Customer information used in Certificate Signing Request for
// creating digital certificates.
//
// swagger:model certification_signing_info
type CertificationSigningInfo struct {

	// The Town or City where customer's business is located.
	City string `json:"city,omitempty"`

	// Common name of the organization or host server
	CommonName string `json:"common_name,omitempty"`

	// Common name is by default <node_uuid>.nutanix.com, but if a customer
	// wants something instead of nutanix.com they can specify it here.
	//
	CommonNameSuffix string `json:"common_name_suffix,omitempty"`

	// Two-letter ISO code for Country where customer's organization is
	// located.
	//
	CountryCode string `json:"country_code,omitempty"`

	// Email address of the certificate administrator.
	EmailAddress string `json:"email_address,omitempty"`

	// Name of the customer business.
	Organization string `json:"organization,omitempty"`

	// The Province, Region, County or State where customer business is
	// is located.
	//
	State string `json:"state,omitempty"`
}

// Validate validates this certification signing info
func (m *CertificationSigningInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certification signing info based on context it is used
func (m *CertificationSigningInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificationSigningInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificationSigningInfo) UnmarshalBinary(b []byte) error {
	var res CertificationSigningInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
