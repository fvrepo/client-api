// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllPortsParams creates a new GetAllPortsParams object
// with the default values initialized.
func NewGetAllPortsParams() GetAllPortsParams {

	var (
		// initialize parameters with default values

		limitDefault = int64(10)
		skipDefault  = int64(0)
	)

	return GetAllPortsParams{
		Limit: limitDefault,

		Skip: skipDefault,
	}
}

// GetAllPortsParams contains all the bound params for the get all ports operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAllPorts
type GetAllPortsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*number of elements to return.
	  In: query
	  Default: 10
	*/
	Limit int64
	/*number of elements to skip
	  In: query
	  Default: 0
	*/
	Skip int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllPortsParams() beforehand.
func (o *GetAllPortsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qSkip, qhkSkip, _ := qs.GetOK("skip")
	if err := o.bindSkip(qSkip, qhkSkip, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetAllPortsParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: true
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllPortsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = value

	return nil
}

// bindSkip binds and validates parameter Skip from query.
func (o *GetAllPortsParams) bindSkip(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: true
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetAllPortsParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("skip", "query", "int64", raw)
	}
	o.Skip = value

	return nil
}
