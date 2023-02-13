// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointJobAuthResultResult hashicorp waypoint job auth result result
//
// swagger:model hashicorp.waypoint.Job.AuthResult.Result
type HashicorpWaypointJobAuthResultResult struct {

	// this is true if the component was authenticated using the Auth
	// callback. If false, then no attempt was made to authenticate. This
	// can be on purpose for example if "check_only" is set to true on
	// the op.
	AuthCompleted bool `json:"auth_completed,omitempty"`

	// auth error
	AuthError *GoogleRPCStatus `json:"auth_error,omitempty"`

	// auth supported is true if this component implemented the auth
	// interface.
	AuthSupported bool `json:"auth_supported,omitempty"`

	// check error
	CheckError *GoogleRPCStatus `json:"check_error,omitempty"`

	// result of the auth check. If the component didn't implement the
	// auth interface this will be set to true. You can check for interface
	// implementation using auth_supported. If auth is attempted, the auth
	// operation will recheck the status and this value will reflect the
	// check post-auth attempt. You can use this to verify if the auth
	// succeeded.
	CheckResult bool `json:"check_result,omitempty"`

	// component that was checked
	Component *HashicorpWaypointComponent `json:"component,omitempty"`
}

// Validate validates this hashicorp waypoint job auth result result
func (m *HashicorpWaypointJobAuthResultResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) validateAuthError(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthError) { // not required
		return nil
	}

	if m.AuthError != nil {
		if err := m.AuthError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) validateCheckError(formats strfmt.Registry) error {
	if swag.IsZero(m.CheckError) { // not required
		return nil
	}

	if m.CheckError != nil {
		if err := m.CheckError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("check_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("check_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) validateComponent(formats strfmt.Registry) error {
	if swag.IsZero(m.Component) { // not required
		return nil
	}

	if m.Component != nil {
		if err := m.Component.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint job auth result result based on the context it is used
func (m *HashicorpWaypointJobAuthResultResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCheckError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComponent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) contextValidateAuthError(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthError != nil {
		if err := m.AuthError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) contextValidateCheckError(ctx context.Context, formats strfmt.Registry) error {

	if m.CheckError != nil {
		if err := m.CheckError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("check_error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("check_error")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointJobAuthResultResult) contextValidateComponent(ctx context.Context, formats strfmt.Registry) error {

	if m.Component != nil {
		if err := m.Component.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("component")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("component")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobAuthResultResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobAuthResultResult) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobAuthResultResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}