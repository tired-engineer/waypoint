// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup hashicorp waypoint get job stream response terminal event step group
//
// swagger:model hashicorp.waypoint.GetJobStreamResponse.Terminal.Event.StepGroup
type HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup struct {

	// close
	Close bool `json:"close,omitempty"`
}

// Validate validates this hashicorp waypoint get job stream response terminal event step group
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp waypoint get job stream response terminal event step group based on context it is used
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointGetJobStreamResponseTerminalEventStepGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}