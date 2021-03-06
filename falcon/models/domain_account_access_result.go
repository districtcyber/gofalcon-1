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

// DomainAccountAccessResult domain account access result
//
// swagger:model domain.AccountAccessResult
type DomainAccountAccessResult struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// reason
	Reason string `json:"reason,omitempty"`

	// successful
	// Required: true
	Successful *bool `json:"successful"`
}

// Validate validates this domain account access result
func (m *DomainAccountAccessResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccessful(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainAccountAccessResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainAccountAccessResult) validateSuccessful(formats strfmt.Registry) error {

	if err := validate.Required("successful", "body", m.Successful); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain account access result based on context it is used
func (m *DomainAccountAccessResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainAccountAccessResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainAccountAccessResult) UnmarshalBinary(b []byte) error {
	var res DomainAccountAccessResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
