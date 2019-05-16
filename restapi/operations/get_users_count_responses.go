// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "goswagger-race/models"
)

// GetUsersCountOKCode is the HTTP code returned for type GetUsersCountOK
const GetUsersCountOKCode int = 200

/*GetUsersCountOK The users were successfully generated

swagger:response getUsersCountOK
*/
type GetUsersCountOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserResponseModel `json:"body,omitempty"`
}

// NewGetUsersCountOK creates GetUsersCountOK with default headers values
func NewGetUsersCountOK() *GetUsersCountOK {

	return &GetUsersCountOK{}
}

// WithPayload adds the payload to the get users count o k response
func (o *GetUsersCountOK) WithPayload(payload *models.UserResponseModel) *GetUsersCountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users count o k response
func (o *GetUsersCountOK) SetPayload(payload *models.UserResponseModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersCountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetUsersCountDefault Error occurred

swagger:response getUsersCountDefault
*/
type GetUsersCountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewGetUsersCountDefault creates GetUsersCountDefault with default headers values
func NewGetUsersCountDefault(code int) *GetUsersCountDefault {
	if code <= 0 {
		code = 500
	}

	return &GetUsersCountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get users count default response
func (o *GetUsersCountDefault) WithStatusCode(code int) *GetUsersCountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get users count default response
func (o *GetUsersCountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get users count default response
func (o *GetUsersCountDefault) WithPayload(payload *models.ErrorModel) *GetUsersCountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users count default response
func (o *GetUsersCountDefault) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersCountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}