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

// RecoveryPlanVMIPAssignment IP configuration for a VM
//
// IP configuration to be applied to a VM.
//
// swagger:model recovery_plan_vm_ip_assignment
type RecoveryPlanVMIPAssignment struct {

	// List of IP configurations for a VM.
	//
	// Required: true
	IPConfigList []*RecoveryPlanVMIPAssignmentIPConfigListItems0 `json:"ip_config_list"`

	// Reference to the VM.
	// Required: true
	VMReference *VMReference `json:"vm_reference"`
}

// Validate validates this recovery plan vm ip assignment
func (m *RecoveryPlanVMIPAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPConfigList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanVMIPAssignment) validateIPConfigList(formats strfmt.Registry) error {

	if err := validate.Required("ip_config_list", "body", m.IPConfigList); err != nil {
		return err
	}

	for i := 0; i < len(m.IPConfigList); i++ {
		if swag.IsZero(m.IPConfigList[i]) { // not required
			continue
		}

		if m.IPConfigList[i] != nil {
			if err := m.IPConfigList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_config_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_config_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanVMIPAssignment) validateVMReference(formats strfmt.Registry) error {

	if err := validate.Required("vm_reference", "body", m.VMReference); err != nil {
		return err
	}

	if m.VMReference != nil {
		if err := m.VMReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_reference")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recovery plan vm ip assignment based on the context it is used
func (m *RecoveryPlanVMIPAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPConfigList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanVMIPAssignment) contextValidateIPConfigList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPConfigList); i++ {

		if m.IPConfigList[i] != nil {

			if swag.IsZero(m.IPConfigList[i]) { // not required
				return nil
			}

			if err := m.IPConfigList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ip_config_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ip_config_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanVMIPAssignment) contextValidateVMReference(ctx context.Context, formats strfmt.Registry) error {

	if m.VMReference != nil {

		if err := m.VMReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_reference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_reference")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanVMIPAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanVMIPAssignment) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanVMIPAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RecoveryPlanVMIPAssignmentIPConfigListItems0 IP configuration.
//
// swagger:model RecoveryPlanVMIPAssignmentIPConfigListItems0
type RecoveryPlanVMIPAssignmentIPConfigListItems0 struct {

	// IP address.
	//
	// Required: true
	// Pattern: ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	IPAddress *string `json:"ip_address"`
}

// Validate validates this recovery plan VM IP assignment IP config list items0
func (m *RecoveryPlanVMIPAssignmentIPConfigListItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanVMIPAssignmentIPConfigListItems0) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ip_address", "body", m.IPAddress); err != nil {
		return err
	}

	if err := validate.Pattern("ip_address", "body", *m.IPAddress, `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recovery plan VM IP assignment IP config list items0 based on context it is used
func (m *RecoveryPlanVMIPAssignmentIPConfigListItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanVMIPAssignmentIPConfigListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanVMIPAssignmentIPConfigListItems0) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanVMIPAssignmentIPConfigListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
