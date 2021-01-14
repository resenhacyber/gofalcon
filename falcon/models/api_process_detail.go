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

// APIProcessDetail api process detail
//
// swagger:model api.ProcessDetail
type APIProcessDetail struct {

	// command line
	// Required: true
	CommandLine *string `json:"command_line"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// file name
	// Required: true
	FileName *string `json:"file_name"`

	// process id
	// Required: true
	ProcessID *string `json:"process_id"`

	// process id local
	// Required: true
	ProcessIDLocal *string `json:"process_id_local"`

	// start timestamp
	// Required: true
	StartTimestamp *string `json:"start_timestamp"`

	// start timestamp raw
	// Required: true
	StartTimestampRaw *string `json:"start_timestamp_raw"`

	// stop timestamp
	// Required: true
	StopTimestamp *string `json:"stop_timestamp"`

	// stop timestamp raw
	// Required: true
	StopTimestampRaw *string `json:"stop_timestamp_raw"`
}

// Validate validates this api process detail
func (m *APIProcessDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandLine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessIDLocal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTimestampRaw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopTimestampRaw(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIProcessDetail) validateCommandLine(formats strfmt.Registry) error {

	if err := validate.Required("command_line", "body", m.CommandLine); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("file_name", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateProcessID(formats strfmt.Registry) error {

	if err := validate.Required("process_id", "body", m.ProcessID); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateProcessIDLocal(formats strfmt.Registry) error {

	if err := validate.Required("process_id_local", "body", m.ProcessIDLocal); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateStartTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("start_timestamp", "body", m.StartTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateStartTimestampRaw(formats strfmt.Registry) error {

	if err := validate.Required("start_timestamp_raw", "body", m.StartTimestampRaw); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateStopTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("stop_timestamp", "body", m.StopTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *APIProcessDetail) validateStopTimestampRaw(formats strfmt.Registry) error {

	if err := validate.Required("stop_timestamp_raw", "body", m.StopTimestampRaw); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIProcessDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIProcessDetail) UnmarshalBinary(b []byte) error {
	var res APIProcessDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
