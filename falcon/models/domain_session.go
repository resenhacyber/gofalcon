// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainSession domain session
//
// swagger:model domain.Session
type DomainSession struct {

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// cloud request ids
	// Required: true
	CloudRequestIds []string `json:"cloud_request_ids"`

	// commands
	// Required: true
	Commands map[string]DomainSessionCommands `json:"commands"`

	// commands queued
	// Required: true
	CommandsQueued *bool `json:"commands_queued"`

	// created at
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"created_at"`

	// deleted at
	// Required: true
	// Format: date-time
	DeletedAt *strfmt.DateTime `json:"deleted_at"`

	// device details
	// Required: true
	DeviceDetails *DomainDevice `json:"device_details"`

	// device id
	// Required: true
	DeviceID *string `json:"device_id"`

	// duration
	// Required: true
	Duration *float64 `json:"duration"`

	// hostname
	// Required: true
	Hostname *string `json:"hostname"`

	// id
	// Required: true
	ID *string `json:"id"`

	// logs
	// Required: true
	Logs []*ModelSessionLog `json:"logs"`

	// offline queued
	// Required: true
	OfflineQueued *bool `json:"offline_queued"`

	// origin
	// Required: true
	Origin *string `json:"origin"`

	// platform id
	// Required: true
	PlatformID *int32 `json:"platform_id"`

	// platform name
	// Required: true
	PlatformName *string `json:"platform_name"`

	// pwd
	// Required: true
	Pwd *string `json:"pwd"`

	// updated at
	// Required: true
	// Format: date-time
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`

	// user uuid
	// Required: true
	UserUUID *string `json:"user_uuid"`
}

// Validate validates this domain session
func (m *DomainSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudRequestIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommandsQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfflineQueued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePwd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainSession) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateCloudRequestIds(formats strfmt.Registry) error {

	if err := validate.Required("cloud_request_ids", "body", m.CloudRequestIds); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateCommands(formats strfmt.Registry) error {

	for k := range m.Commands {

		if err := validate.Required("commands"+"."+k, "body", m.Commands[k]); err != nil {
			return err
		}

		if err := validate.Required("commands"+"."+k, "body", m.Commands[k]); err != nil {
			return err
		}
		if val, ok := m.Commands[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DomainSession) validateCommandsQueued(formats strfmt.Registry) error {

	if err := validate.Required("commands_queued", "body", m.CommandsQueued); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateDeletedAt(formats strfmt.Registry) error {

	if err := validate.Required("deleted_at", "body", m.DeletedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("deleted_at", "body", "date-time", m.DeletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateDeviceDetails(formats strfmt.Registry) error {

	if err := validate.Required("device_details", "body", m.DeviceDetails); err != nil {
		return err
	}

	if m.DeviceDetails != nil {
		if err := m.DeviceDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_details")
			}
			return err
		}
	}

	return nil
}

func (m *DomainSession) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateLogs(formats strfmt.Registry) error {

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainSession) validateOfflineQueued(formats strfmt.Registry) error {

	if err := validate.Required("offline_queued", "body", m.OfflineQueued); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateOrigin(formats strfmt.Registry) error {

	if err := validate.Required("origin", "body", m.Origin); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validatePlatformID(formats strfmt.Registry) error {

	if err := validate.Required("platform_id", "body", m.PlatformID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validatePlatformName(formats strfmt.Registry) error {

	if err := validate.Required("platform_name", "body", m.PlatformName); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validatePwd(formats strfmt.Registry) error {

	if err := validate.Required("pwd", "body", m.Pwd); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

func (m *DomainSession) validateUserUUID(formats strfmt.Registry) error {

	if err := validate.Required("user_uuid", "body", m.UserUUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainSession) UnmarshalBinary(b []byte) error {
	var res DomainSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
