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

// APIRuleV1 api rule v1
//
// swagger:model api.RuleV1
type APIRuleV1 struct {

	// action label
	// Required: true
	ActionLabel *string `json:"action_label"`

	// comment
	// Required: true
	Comment *string `json:"comment"`

	// committed on
	// Required: true
	// Format: date-time
	CommittedOn *strfmt.DateTime `json:"committed_on"`

	// created by
	// Required: true
	CreatedBy *string `json:"created_by"`

	// created on
	// Required: true
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on"`

	// customer id
	// Required: true
	CustomerID *string `json:"customer_id"`

	// deleted
	// Required: true
	Deleted *bool `json:"deleted"`

	// description
	// Required: true
	Description *string `json:"description"`

	// disposition id
	// Required: true
	DispositionID *int32 `json:"disposition_id"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// field values
	// Required: true
	FieldValues []*DomainFieldValue `json:"field_values"`

	// instance id
	// Required: true
	InstanceID *string `json:"instance_id"`

	// instance version
	// Required: true
	InstanceVersion *int64 `json:"instance_version"`

	// magic cookie
	// Required: true
	MagicCookie *int64 `json:"magic_cookie"`

	// modified by
	// Required: true
	ModifiedBy *string `json:"modified_by"`

	// modified on
	// Required: true
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on"`

	// name
	// Required: true
	Name *string `json:"name"`

	// pattern id
	// Required: true
	PatternID *string `json:"pattern_id"`

	// pattern severity
	// Required: true
	PatternSeverity *string `json:"pattern_severity"`

	// rulegroup id
	// Required: true
	RulegroupID *string `json:"rulegroup_id"`

	// ruletype id
	// Required: true
	RuletypeID *string `json:"ruletype_id"`

	// ruletype name
	// Required: true
	RuletypeName *string `json:"ruletype_name"`

	// version ids
	// Required: true
	VersionIds []string `json:"version_ids"`
}

// Validate validates this api rule v1
func (m *APIRuleV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommittedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDispositionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMagicCookie(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatternID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePatternSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRulegroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuletypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuletypeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIRuleV1) validateActionLabel(formats strfmt.Registry) error {

	if err := validate.Required("action_label", "body", m.ActionLabel); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateCommittedOn(formats strfmt.Registry) error {

	if err := validate.Required("committed_on", "body", m.CommittedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("committed_on", "body", "date-time", m.CommittedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("created_by", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateCreatedOn(formats strfmt.Registry) error {

	if err := validate.Required("created_on", "body", m.CreatedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateCustomerID(formats strfmt.Registry) error {

	if err := validate.Required("customer_id", "body", m.CustomerID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", m.Deleted); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateDispositionID(formats strfmt.Registry) error {

	if err := validate.Required("disposition_id", "body", m.DispositionID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateFieldValues(formats strfmt.Registry) error {

	if err := validate.Required("field_values", "body", m.FieldValues); err != nil {
		return err
	}

	for i := 0; i < len(m.FieldValues); i++ {
		if swag.IsZero(m.FieldValues[i]) { // not required
			continue
		}

		if m.FieldValues[i] != nil {
			if err := m.FieldValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("field_values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *APIRuleV1) validateInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("instance_id", "body", m.InstanceID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateInstanceVersion(formats strfmt.Registry) error {

	if err := validate.Required("instance_version", "body", m.InstanceVersion); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateMagicCookie(formats strfmt.Registry) error {

	if err := validate.Required("magic_cookie", "body", m.MagicCookie); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateModifiedBy(formats strfmt.Registry) error {

	if err := validate.Required("modified_by", "body", m.ModifiedBy); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateModifiedOn(formats strfmt.Registry) error {

	if err := validate.Required("modified_on", "body", m.ModifiedOn); err != nil {
		return err
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validatePatternID(formats strfmt.Registry) error {

	if err := validate.Required("pattern_id", "body", m.PatternID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validatePatternSeverity(formats strfmt.Registry) error {

	if err := validate.Required("pattern_severity", "body", m.PatternSeverity); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateRulegroupID(formats strfmt.Registry) error {

	if err := validate.Required("rulegroup_id", "body", m.RulegroupID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateRuletypeID(formats strfmt.Registry) error {

	if err := validate.Required("ruletype_id", "body", m.RuletypeID); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateRuletypeName(formats strfmt.Registry) error {

	if err := validate.Required("ruletype_name", "body", m.RuletypeName); err != nil {
		return err
	}

	return nil
}

func (m *APIRuleV1) validateVersionIds(formats strfmt.Registry) error {

	if err := validate.Required("version_ids", "body", m.VersionIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIRuleV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIRuleV1) UnmarshalBinary(b []byte) error {
	var res APIRuleV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
