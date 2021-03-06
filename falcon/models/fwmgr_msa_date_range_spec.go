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

// FwmgrMsaDateRangeSpec fwmgr msa date range spec
//
// swagger:model fwmgr.msa.DateRangeSpec
type FwmgrMsaDateRangeSpec struct {

	// from
	// Required: true
	From *string `json:"from"`

	// to
	// Required: true
	To *string `json:"to"`
}

// Validate validates this fwmgr msa date range spec
func (m *FwmgrMsaDateRangeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FwmgrMsaDateRangeSpec) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *FwmgrMsaDateRangeSpec) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fwmgr msa date range spec based on context it is used
func (m *FwmgrMsaDateRangeSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FwmgrMsaDateRangeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FwmgrMsaDateRangeSpec) UnmarshalBinary(b []byte) error {
	var res FwmgrMsaDateRangeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
