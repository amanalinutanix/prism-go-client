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

// RecoveryPlanStage Information about a stage in the Recovery Plan
//
// A stage specifies the work to be performed when the Recovery Plan is executed. A stage work can be execute an user script or recover entities in case of failover. If there are multiple entities to recover in a stage, all of them will be recovered in parallel.
//
// swagger:model recovery_plan_stage
type RecoveryPlanStage struct {

	// Amount of time in seconds to delay the execution of next stage after execution of current stage.
	//
	// Minimum: 0
	DelayTimeSecs *int64 `json:"delay_time_secs,omitempty"`

	// UUID of stage.
	// Pattern: ^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$
	StageUUID string `json:"stage_uuid,omitempty"`

	// stage work
	// Required: true
	StageWork *RecoveryPlanStageStageWork `json:"stage_work"`
}

// Validate validates this recovery plan stage
func (m *RecoveryPlanStage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayTimeSecs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStageWork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanStage) validateDelayTimeSecs(formats strfmt.Registry) error {
	if swag.IsZero(m.DelayTimeSecs) { // not required
		return nil
	}

	if err := validate.MinimumInt("delay_time_secs", "body", *m.DelayTimeSecs, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanStage) validateStageUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.StageUUID) { // not required
		return nil
	}

	if err := validate.Pattern("stage_uuid", "body", m.StageUUID, `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanStage) validateStageWork(formats strfmt.Registry) error {

	if err := validate.Required("stage_work", "body", m.StageWork); err != nil {
		return err
	}

	if m.StageWork != nil {
		if err := m.StageWork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage_work")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage_work")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recovery plan stage based on the context it is used
func (m *RecoveryPlanStage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStageWork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanStage) contextValidateStageWork(ctx context.Context, formats strfmt.Registry) error {

	if m.StageWork != nil {

		if err := m.StageWork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage_work")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage_work")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanStage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanStage) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanStage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RecoveryPlanStageStageWork Information about work to be performed as part of this stage on execution. Only one of recover_entities or script has to be provided.
//
// swagger:model RecoveryPlanStageStageWork
type RecoveryPlanStageStageWork struct {

	// Information about entities to be recovered. This can be either explicit entity list or entity filter to identify list of entities to be recovered.
	//
	RecoverEntities *RecoverEntities `json:"recover_entities,omitempty"`
}

// Validate validates this recovery plan stage stage work
func (m *RecoveryPlanStageStageWork) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecoverEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanStageStageWork) validateRecoverEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoverEntities) { // not required
		return nil
	}

	if m.RecoverEntities != nil {
		if err := m.RecoverEntities.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage_work" + "." + "recover_entities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage_work" + "." + "recover_entities")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recovery plan stage stage work based on the context it is used
func (m *RecoveryPlanStageStageWork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRecoverEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanStageStageWork) contextValidateRecoverEntities(ctx context.Context, formats strfmt.Registry) error {

	if m.RecoverEntities != nil {

		if swag.IsZero(m.RecoverEntities) { // not required
			return nil
		}

		if err := m.RecoverEntities.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stage_work" + "." + "recover_entities")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stage_work" + "." + "recover_entities")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanStageStageWork) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanStageStageWork) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanStageStageWork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
