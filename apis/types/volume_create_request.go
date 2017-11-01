// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VolumeCreateRequest VolumeCreateRequest contains the response for the remote API:
// POST "/volumes/create
//
// swagger:model VolumeCreateRequest

type VolumeCreateRequest struct {

	// Driver is the Driver name used to create the volume.
	Driver string `json:"Driver,omitempty"`

	// DriverOpts holds the driver specific options to use for when creating the volume.
	DriverOpts map[string]string `json:"DriverOpts,omitempty"`

	// Labels is metadata specific to the volume.
	Labels map[string]string `json:"Labels,omitempty"`

	// Name is the name of the volume.
	Name string `json:"Name,omitempty"`
}

/* polymorph VolumeCreateRequest Driver false */

/* polymorph VolumeCreateRequest DriverOpts false */

/* polymorph VolumeCreateRequest Labels false */

/* polymorph VolumeCreateRequest Name false */

// Validate validates this volume create request
func (m *VolumeCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeCreateRequest) UnmarshalBinary(b []byte) error {
	var res VolumeCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}