// Code generated by go-swagger; DO NOT EDIT.

package falconx_sandbox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUploadSampleV2Params creates a new UploadSampleV2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadSampleV2Params() *UploadSampleV2Params {
	return &UploadSampleV2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadSampleV2ParamsWithTimeout creates a new UploadSampleV2Params object
// with the ability to set a timeout on a request.
func NewUploadSampleV2ParamsWithTimeout(timeout time.Duration) *UploadSampleV2Params {
	return &UploadSampleV2Params{
		timeout: timeout,
	}
}

// NewUploadSampleV2ParamsWithContext creates a new UploadSampleV2Params object
// with the ability to set a context for a request.
func NewUploadSampleV2ParamsWithContext(ctx context.Context) *UploadSampleV2Params {
	return &UploadSampleV2Params{
		Context: ctx,
	}
}

// NewUploadSampleV2ParamsWithHTTPClient creates a new UploadSampleV2Params object
// with the ability to set a custom HTTPClient for a request.
func NewUploadSampleV2ParamsWithHTTPClient(client *http.Client) *UploadSampleV2Params {
	return &UploadSampleV2Params{
		HTTPClient: client,
	}
}

/* UploadSampleV2Params contains all the parameters to send to the API endpoint
   for the upload sample v2 operation.

   Typically these are written to a http.Request.
*/
type UploadSampleV2Params struct {

	/* XCSUSERUUID.

	   User UUID
	*/
	XCSUSERUUID *string

	/* Body.

	     Content of the uploaded sample in binary format. For example, use `--data-binary @$FILE_PATH` when using cURL. Max file size: 100 MB.

	Accepted file formats:

	- Portable executables: `.exe`, `.scr`, `.pif`, `.dll`, `.com`, `.cpl`, etc.
	- Office documents: `.doc`, `.docx`, `.ppt`, `.pps`, `.pptx`, `.ppsx`, `.xls`, `.xlsx`, `.rtf`, `.pub`
	- PDF
	- APK
	- Executable JAR
	- Windows script component: `.sct`
	- Windows shortcut: `.lnk`
	- Windows help: `.chm`
	- HTML application: `.hta`
	- Windows script file: `.wsf`
	- Javascript: `.js`
	- Visual Basic: `.vbs`,  `.vbe`
	- Shockwave Flash: `.swf`
	- Perl: `.pl`
	- Powershell: `.ps1`, `.psd1`, `.psm1`
	- Scalable vector graphics: `.svg`
	- Python: `.py`
	- Linux ELF executables
	- Email files: MIME RFC 822 `.eml`, Outlook `.msg`.
	*/
	Body []int64

	/* Comment.

	   A descriptive comment to identify the file for other users.
	*/
	Comment *string

	/* FileName.

	   Name of the file.
	*/
	FileName string

	/* IsConfidential.

	     Defines visibility of this file in Falcon MalQuery, either via the API or the Falcon console.

	- `true`: File is only shown to users within your customer account
	- `false`: File can be seen by other CrowdStrike customers

	Default: `true`.

	     Default: true
	*/
	IsConfidential *bool

	/* Upfile.

	   The binary file.
	*/
	Upfile runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload sample v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadSampleV2Params) WithDefaults() *UploadSampleV2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload sample v2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadSampleV2Params) SetDefaults() {
	var (
		isConfidentialDefault = bool(true)
	)

	val := UploadSampleV2Params{
		IsConfidential: &isConfidentialDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upload sample v2 params
func (o *UploadSampleV2Params) WithTimeout(timeout time.Duration) *UploadSampleV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload sample v2 params
func (o *UploadSampleV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload sample v2 params
func (o *UploadSampleV2Params) WithContext(ctx context.Context) *UploadSampleV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload sample v2 params
func (o *UploadSampleV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload sample v2 params
func (o *UploadSampleV2Params) WithHTTPClient(client *http.Client) *UploadSampleV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload sample v2 params
func (o *UploadSampleV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the upload sample v2 params
func (o *UploadSampleV2Params) WithXCSUSERUUID(xCSUSERUUID *string) *UploadSampleV2Params {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the upload sample v2 params
func (o *UploadSampleV2Params) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithBody adds the body to the upload sample v2 params
func (o *UploadSampleV2Params) WithBody(body []int64) *UploadSampleV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload sample v2 params
func (o *UploadSampleV2Params) SetBody(body []int64) {
	o.Body = body
}

// WithComment adds the comment to the upload sample v2 params
func (o *UploadSampleV2Params) WithComment(comment *string) *UploadSampleV2Params {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the upload sample v2 params
func (o *UploadSampleV2Params) SetComment(comment *string) {
	o.Comment = comment
}

// WithFileName adds the fileName to the upload sample v2 params
func (o *UploadSampleV2Params) WithFileName(fileName string) *UploadSampleV2Params {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the upload sample v2 params
func (o *UploadSampleV2Params) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithIsConfidential adds the isConfidential to the upload sample v2 params
func (o *UploadSampleV2Params) WithIsConfidential(isConfidential *bool) *UploadSampleV2Params {
	o.SetIsConfidential(isConfidential)
	return o
}

// SetIsConfidential adds the isConfidential to the upload sample v2 params
func (o *UploadSampleV2Params) SetIsConfidential(isConfidential *bool) {
	o.IsConfidential = isConfidential
}

// WithUpfile adds the upfile to the upload sample v2 params
func (o *UploadSampleV2Params) WithUpfile(upfile runtime.NamedReadCloser) *UploadSampleV2Params {
	o.SetUpfile(upfile)
	return o
}

// SetUpfile adds the upfile to the upload sample v2 params
func (o *UploadSampleV2Params) SetUpfile(upfile runtime.NamedReadCloser) {
	o.Upfile = upfile
}

// WriteToRequest writes these params to a swagger request
func (o *UploadSampleV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XCSUSERUUID != nil {

		// header param X-CS-USERUUID
		if err := r.SetHeaderParam("X-CS-USERUUID", *o.XCSUSERUUID); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {

		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
	}

	if o.IsConfidential != nil {

		// query param is_confidential
		var qrIsConfidential bool

		if o.IsConfidential != nil {
			qrIsConfidential = *o.IsConfidential
		}
		qIsConfidential := swag.FormatBool(qrIsConfidential)
		if qIsConfidential != "" {

			if err := r.SetQueryParam("is_confidential", qIsConfidential); err != nil {
				return err
			}
		}
	}
	// form file param upfile
	if err := r.SetFileParam("upfile", o.Upfile); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
