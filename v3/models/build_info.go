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

// BuildInfo Cluster build details
//
// Cluster build details.
//
// swagger:model build_info
type BuildInfo struct {

	// Build type, one of {dbg, opt, release}.
	// Required: true
	BuildType *string `json:"build_type"`

	// Date/time of the last commit.
	// Format: date-time
	CommitDate strfmt.DateTime `json:"commit_date,omitempty"`

	// Last Git commit id which the build is based on.
	// Required: true
	CommitID *string `json:"commit_id"`

	// Full version name.
	FullVersion string `json:"full_version,omitempty"`

	// Flag to indicate if the AOS release is qualified as long term support.
	//
	IsLongTermSupport bool `json:"is_long_term_support,omitempty"`

	// First 6 characters of the last Git commit id.
	// Required: true
	ShortCommitID *string `json:"short_commit_id"`

	// Numeric version. e.g. "5.5"
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this build info
func (m *BuildInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortCommitID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildInfo) validateBuildType(formats strfmt.Registry) error {

	if err := validate.Required("build_type", "body", m.BuildType); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateCommitDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CommitDate) { // not required
		return nil
	}

	if err := validate.FormatOf("commit_date", "body", "date-time", m.CommitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateCommitID(formats strfmt.Registry) error {

	if err := validate.Required("commit_id", "body", m.CommitID); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateShortCommitID(formats strfmt.Registry) error {

	if err := validate.Required("short_commit_id", "body", m.ShortCommitID); err != nil {
		return err
	}

	return nil
}

func (m *BuildInfo) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this build info based on context it is used
func (m *BuildInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuildInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildInfo) UnmarshalBinary(b []byte) error {
	var res BuildInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
