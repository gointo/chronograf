package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*GetTokenOK A JWT authentication token

swagger:response getTokenOK
*/
type GetTokenOK struct {

	// In: body
	Payload models.Token `json:"body,omitempty"`
}

// NewGetTokenOK creates GetTokenOK with default headers values
func NewGetTokenOK() *GetTokenOK {
	return &GetTokenOK{}
}

// WithPayload adds the payload to the get token o k response
func (o *GetTokenOK) WithPayload(payload models.Token) *GetTokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get token o k response
func (o *GetTokenOK) SetPayload(payload models.Token) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetTokenDefault Unexpected internal service error

swagger:response getTokenDefault
*/
type GetTokenDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTokenDefault creates GetTokenDefault with default headers values
func NewGetTokenDefault(code int) *GetTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &GetTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get token default response
func (o *GetTokenDefault) WithStatusCode(code int) *GetTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get token default response
func (o *GetTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get token default response
func (o *GetTokenDefault) WithPayload(payload *models.Error) *GetTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get token default response
func (o *GetTokenDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}