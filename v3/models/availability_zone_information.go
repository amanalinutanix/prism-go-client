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

// AvailabilityZoneInformation Availability Zone information
//
// Availability Zone information.
//
// swagger:model availability_zone_information
type AvailabilityZoneInformation struct {

	// URL of the Availability Zone.
	//
	// Required: true
	AvailabilityZoneURL *string `json:"availability_zone_url"`

	// List of cluster references. This is applicable only in scenario where failed and recovery clusters both are managed by the same Availability Zone.
	ClusterReferenceList []*ClusterReference `json:"cluster_reference_list"`
}

// Validate validates this availability zone information
func (m *AvailabilityZoneInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityZoneURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterReferenceList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailabilityZoneInformation) validateAvailabilityZoneURL(formats strfmt.Registry) error {

	if err := validate.Required("availability_zone_url", "body", m.AvailabilityZoneURL); err != nil {
		return err
	}

	return nil
}

func (m *AvailabilityZoneInformation) validateClusterReferenceList(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterReferenceList) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterReferenceList); i++ {
		if swag.IsZero(m.ClusterReferenceList[i]) { // not required
			continue
		}

		if m.ClusterReferenceList[i] != nil {
			if err := m.ClusterReferenceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this availability zone information based on the context it is used
func (m *AvailabilityZoneInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterReferenceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailabilityZoneInformation) contextValidateClusterReferenceList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterReferenceList); i++ {

		if m.ClusterReferenceList[i] != nil {

			if swag.IsZero(m.ClusterReferenceList[i]) { // not required
				return nil
			}

			if err := m.ClusterReferenceList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_reference_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_reference_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AvailabilityZoneInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailabilityZoneInformation) UnmarshalBinary(b []byte) error {
	var res AvailabilityZoneInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
