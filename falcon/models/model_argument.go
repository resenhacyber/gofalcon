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

// ModelArgument model argument
//
// swagger:model model.Argument
type ModelArgument struct {

	// arg name
	// Required: true
	ArgName *string `json:"arg_name"`

	// arg type
	// Required: true
	ArgType *string `json:"arg_type"`

	// command level
	// Required: true
	CommandLevel *string `json:"command_level"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// data type
	// Required: true
	DataType *string `json:"data_type"`

	// default value
	// Required: true
	DefaultValue *string `json:"default_value"`

	// description
	// Required: true
	Description *string `json:"description"`

	// encoding
	// Required: true
	Encoding *string `json:"encoding"`

	// id
	// Required: true
	ID *int64 `json:"id"`

	// options
	// Required: true
	Options []string `json:"options"`

	// required
	// Required: true
	Required *bool `json:"required"`

	// requires value
	// Required: true
	RequiresValue *bool `json:"requires_value"`

	// script id
	// Required: true
	ScriptID *int64 `json:"script_id"`

	// sequence
	// Required: true
	Sequence *int64 `json:"sequence"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this model argument
func (m *ModelArgument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArgType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommandLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncoding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiresValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScriptID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelArgument) validateArgName(formats strfmt.Registry) error {

	if err := validate.Required("arg_name", "body", m.ArgName); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateArgType(formats strfmt.Registry) error {

	if err := validate.Required("arg_type", "body", m.ArgType); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateCommandLevel(formats strfmt.Registry) error {

	if err := validate.Required("command_level", "body", m.CommandLevel); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("data_type", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateDefaultValue(formats strfmt.Registry) error {

	if err := validate.Required("default_value", "body", m.DefaultValue); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateEncoding(formats strfmt.Registry) error {

	if err := validate.Required("encoding", "body", m.Encoding); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateRequired(formats strfmt.Registry) error {

	if err := validate.Required("required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateRequiresValue(formats strfmt.Registry) error {

	if err := validate.Required("requires_value", "body", m.RequiresValue); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Required("script_id", "body", m.ScriptID); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateSequence(formats strfmt.Registry) error {

	if err := validate.Required("sequence", "body", m.Sequence); err != nil {
		return err
	}

	return nil
}

func (m *ModelArgument) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelArgument) UnmarshalBinary(b []byte) error {
	var res ModelArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
