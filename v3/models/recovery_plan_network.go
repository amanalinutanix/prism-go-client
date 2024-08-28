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

// RecoveryPlanNetwork Network configuration for the Recovery Plan
//
// Network configuration to be used for performing network mapping and IP preservation/mapping on Recovery Plan execution.
//
// swagger:model recovery_plan_network
type RecoveryPlanNetwork struct {

	// Name of the network.
	//
	// Max Length: 64
	Name string `json:"name,omitempty"`

	// List of subnets for the network.
	//
	SubnetList []*RecoveryPlanSubnetConfig `json:"subnet_list"`

	// Client need to specify this field as true while using vpc_reference for specifying the VPC for the network. Without this values in vpc_reference will be ignored.
	//
	UseVpcReference bool `json:"use_vpc_reference,omitempty"`

	// Reference to the Virtual Network. This reference is deprecated, use vpc_reference instead.
	VirtualNetworkReference *VirtualNetworkReference `json:"virtual_network_reference,omitempty"`

	// Reference to the VPC.
	VpcReference *VpcReference `json:"vpc_reference,omitempty"`
}

// Validate validates this recovery plan network
func (m *RecoveryPlanNetwork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualNetworkReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpcReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanNetwork) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanNetwork) validateSubnetList(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetList) { // not required
		return nil
	}

	for i := 0; i < len(m.SubnetList); i++ {
		if swag.IsZero(m.SubnetList[i]) { // not required
			continue
		}

		if m.SubnetList[i] != nil {
			if err := m.SubnetList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnet_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnet_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanNetwork) validateVirtualNetworkReference(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualNetworkReference) { // not required
		return nil
	}

	if m.VirtualNetworkReference != nil {
		if err := m.VirtualNetworkReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_network_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_network_reference")
			}
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanNetwork) validateVpcReference(formats strfmt.Registry) error {
	if swag.IsZero(m.VpcReference) { // not required
		return nil
	}

	if m.VpcReference != nil {
		if err := m.VpcReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_reference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recovery plan network based on the context it is used
func (m *RecoveryPlanNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubnetList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualNetworkReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVpcReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanNetwork) contextValidateSubnetList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubnetList); i++ {

		if m.SubnetList[i] != nil {

			if swag.IsZero(m.SubnetList[i]) { // not required
				return nil
			}

			if err := m.SubnetList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnet_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnet_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanNetwork) contextValidateVirtualNetworkReference(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualNetworkReference != nil {

		if swag.IsZero(m.VirtualNetworkReference) { // not required
			return nil
		}

		if err := m.VirtualNetworkReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_network_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_network_reference")
			}
			return err
		}
	}

	return nil
}

func (m *RecoveryPlanNetwork) contextValidateVpcReference(ctx context.Context, formats strfmt.Registry) error {

	if m.VpcReference != nil {

		if swag.IsZero(m.VpcReference) { // not required
			return nil
		}

		if err := m.VpcReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vpc_reference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanNetwork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanNetwork) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
