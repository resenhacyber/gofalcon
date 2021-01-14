// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationPolicyExtV1 registration policy ext v1
//
// swagger:model registration.PolicyExtV1
type RegistrationPolicyExtV1 struct {

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// policy id
	// Required: true
	PolicyID *int32 `json:"policy_id"`

	// severity
	// Required: true
	Severity *string `json:"severity"`
}

// Validate validates this registration policy ext v1
func (m *RegistrationPolicyExtV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationPolicyExtV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationPolicyExtV1) validatePolicyID(formats strfmt.Registry) error {

	if err := validate.Required("policy_id", "body", m.PolicyID); err != nil {
		return err
	}

	return nil
}

func (m *RegistrationPolicyExtV1) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationPolicyExtV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationPolicyExtV1) UnmarshalBinary(b []byte) error {
	var res RegistrationPolicyExtV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
