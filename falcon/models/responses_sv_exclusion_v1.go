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

// ResponsesSvExclusionV1 responses sv exclusion v1
//
// swagger:model responses.SvExclusionV1
type ResponsesSvExclusionV1 struct {

	// applied globally
	// Required: true
	AppliedGlobally *bool `json:"applied_globally"`

	// created by
	// Required: true
	CreatedBy *string `json:"created_by"`

	// created on
	// Required: true
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on"`

	// groups
	// Required: true
	Groups []*ResponsesHostGroupV1 `json:"groups"`

	// id
	// Required: true
	ID *string `json:"id"`

	// last modified
	// Required: true
	// Format: date-time
	LastModified *strfmt.DateTime `json:"last_modified"`

	// modified by
	// Required: true
	ModifiedBy *string `json:"modified_by"`

	// regexp value
	// Required: true
	RegexpValue *string `json:"regexp_value"`

	// value
	// Required: true
	Value *string `json:"value"`

	// value hash
	// Required: true
	ValueHash *string `json:"value_hash"`
}

// Validate validates this responses sv exclusion v1
func (m *ResponsesSvExclusionV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppliedGlobally(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegexpValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesSvExclusionV1) validateAppliedGlobally(formats strfmt.Registry) error {

	if err := validate.Required("applied_globally", "body", m.AppliedGlobally); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	for i := 0; i < len(m.Groups); i++ {
		if swag.IsZero(m.Groups[i]) { // not required
			continue
		}

		if m.Groups[i] != nil {
			if err := m.Groups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateLastModified(formats strfmt.Registry) error {

	if err := validate.Required("last_modified", "body", m.LastModified); err != nil {
		return err
	}

	if err := validate.FormatOf("last_modified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("modified_by", "body", m.ModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateRegexpValue(formats strfmt.Registry) error {

	if err := validate.Required("regexp_value", "body", m.RegexpValue); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesSvExclusionV1) validateValueHash(formats strfmt.Registry) error {

	if err := validate.Required("value_hash", "body", m.ValueHash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this responses sv exclusion v1 based on the context it is used
func (m *ResponsesSvExclusionV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesSvExclusionV1) contextValidateGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Groups); i++ {

		if m.Groups[i] != nil {
			if err := m.Groups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("groups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("groups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesSvExclusionV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesSvExclusionV1) UnmarshalBinary(b []byte) error {
	var res ResponsesSvExclusionV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
