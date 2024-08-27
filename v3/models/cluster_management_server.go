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

// ClusterManagementServer Cluster Management Server
//
// Cluster Management server information.
//
// swagger:model cluster_management_server
type ClusterManagementServer struct {

	// Denotes if DRS is enabled or not.
	DrsEnabled bool `json:"drs_enabled,omitempty"`

	// ip
	// Required: true
	// Pattern: ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	IP *string `json:"ip"`

	// Array of management server status: - 'REGISTERED': Indicates whether the server is registered with
	//                 Nutanix or not.
	// - 'IN_USE': Indicates whether any host is managed by this server or
	//             not.
	//
	StatusList []string `json:"status_list"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this cluster management server
func (m *ClusterManagementServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterManagementServer) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	if err := validate.Pattern("ip", "body", *m.IP, `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		return err
	}

	return nil
}

func (m *ClusterManagementServer) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster management server based on context it is used
func (m *ClusterManagementServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterManagementServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterManagementServer) UnmarshalBinary(b []byte) error {
	var res ClusterManagementServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
