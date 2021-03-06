// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponsesPreventionSettingV1 A prevention policy setting
//
// swagger:model responses.PreventionSettingV1
type ResponsesPreventionSettingV1 struct {

	// The human readable description of the setting
	Description string `json:"description,omitempty"`

	// The id of the setting
	// Required: true
	ID *string `json:"id"`

	// The name of the setting
	// Required: true
	Name *string `json:"name"`

	// The type of the setting which can be used as a hint when displaying in the UI
	// Required: true
	// Enum: [toggle mlslider]
	Type *string `json:"type"`

	// The value for the setting. For a toggle this value will take the form {'enabled':true|false}. For an mlslider this will take the form {'detection':'DISABLED|CAUTIOUS|MODERATE|AGGRESSIVE|EXTRA_AGGRESSIVE','prevention':'DISABLED|CAUTIOUS|MODERATE|AGGRESSIVE|EXTRA_AGGRESSIVE'}
	// Required: true
	Value interface{} `json:"value"`
}

// Validate validates this responses prevention setting v1
func (m *ResponsesPreventionSettingV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResponsesPreventionSettingV1) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionSettingV1) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var responsesPreventionSettingV1TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["toggle","mlslider"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		responsesPreventionSettingV1TypeTypePropEnum = append(responsesPreventionSettingV1TypeTypePropEnum, v)
	}
}

const (

	// ResponsesPreventionSettingV1TypeToggle captures enum value "toggle"
	ResponsesPreventionSettingV1TypeToggle string = "toggle"

	// ResponsesPreventionSettingV1TypeMlslider captures enum value "mlslider"
	ResponsesPreventionSettingV1TypeMlslider string = "mlslider"
)

// prop value enum
func (m *ResponsesPreventionSettingV1) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, responsesPreventionSettingV1TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResponsesPreventionSettingV1) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ResponsesPreventionSettingV1) validateValue(formats strfmt.Registry) error {

	if m.Value == nil {
		return errors.Required("value", "body", nil)
	}

	return nil
}

// ContextValidate validates this responses prevention setting v1 based on context it is used
func (m *ResponsesPreventionSettingV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponsesPreventionSettingV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponsesPreventionSettingV1) UnmarshalBinary(b []byte) error {
	var res ResponsesPreventionSettingV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
