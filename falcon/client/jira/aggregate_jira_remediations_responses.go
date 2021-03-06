// Code generated by go-swagger; DO NOT EDIT.

package jira

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// AggregateJiraRemediationsReader is a Reader for the AggregateJiraRemediations structure.
type AggregateJiraRemediationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateJiraRemediationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateJiraRemediationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateJiraRemediationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateJiraRemediationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregateJiraRemediationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateJiraRemediationsOK creates a AggregateJiraRemediationsOK with default headers values
func NewAggregateJiraRemediationsOK() *AggregateJiraRemediationsOK {
	return &AggregateJiraRemediationsOK{}
}

/* AggregateJiraRemediationsOK describes a response with status code 200, with default header values.

OK
*/
type AggregateJiraRemediationsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

func (o *AggregateJiraRemediationsOK) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/jira-remediations/GET/v1][%d] aggregateJiraRemediationsOK  %+v", 200, o.Payload)
}
func (o *AggregateJiraRemediationsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateJiraRemediationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateJiraRemediationsForbidden creates a AggregateJiraRemediationsForbidden with default headers values
func NewAggregateJiraRemediationsForbidden() *AggregateJiraRemediationsForbidden {
	return &AggregateJiraRemediationsForbidden{}
}

/* AggregateJiraRemediationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateJiraRemediationsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *AggregateJiraRemediationsForbidden) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/jira-remediations/GET/v1][%d] aggregateJiraRemediationsForbidden  %+v", 403, o.Payload)
}
func (o *AggregateJiraRemediationsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateJiraRemediationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateJiraRemediationsTooManyRequests creates a AggregateJiraRemediationsTooManyRequests with default headers values
func NewAggregateJiraRemediationsTooManyRequests() *AggregateJiraRemediationsTooManyRequests {
	return &AggregateJiraRemediationsTooManyRequests{}
}

/* AggregateJiraRemediationsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateJiraRemediationsTooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *AggregateJiraRemediationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/jira-remediations/GET/v1][%d] aggregateJiraRemediationsTooManyRequests  %+v", 429, o.Payload)
}
func (o *AggregateJiraRemediationsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateJiraRemediationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateJiraRemediationsDefault creates a AggregateJiraRemediationsDefault with default headers values
func NewAggregateJiraRemediationsDefault(code int) *AggregateJiraRemediationsDefault {
	return &AggregateJiraRemediationsDefault{
		_statusCode: code,
	}
}

/* AggregateJiraRemediationsDefault describes a response with status code -1, with default header values.

OK
*/
type AggregateJiraRemediationsDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// Code gets the status code for the aggregate jira remediations default response
func (o *AggregateJiraRemediationsDefault) Code() int {
	return o._statusCode
}

func (o *AggregateJiraRemediationsDefault) Error() string {
	return fmt.Sprintf("[POST /falcon-complete-dashboards/aggregates/jira-remediations/GET/v1][%d] AggregateJiraRemediations default  %+v", o._statusCode, o.Payload)
}
func (o *AggregateJiraRemediationsDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateJiraRemediationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
