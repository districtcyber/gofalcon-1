// Code generated by go-swagger; DO NOT EDIT.

package sample_uploads

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

// NewUploadSampleV3Params creates a new UploadSampleV3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadSampleV3Params() *UploadSampleV3Params {
	return &UploadSampleV3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadSampleV3ParamsWithTimeout creates a new UploadSampleV3Params object
// with the ability to set a timeout on a request.
func NewUploadSampleV3ParamsWithTimeout(timeout time.Duration) *UploadSampleV3Params {
	return &UploadSampleV3Params{
		timeout: timeout,
	}
}

// NewUploadSampleV3ParamsWithContext creates a new UploadSampleV3Params object
// with the ability to set a context for a request.
func NewUploadSampleV3ParamsWithContext(ctx context.Context) *UploadSampleV3Params {
	return &UploadSampleV3Params{
		Context: ctx,
	}
}

// NewUploadSampleV3ParamsWithHTTPClient creates a new UploadSampleV3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUploadSampleV3ParamsWithHTTPClient(client *http.Client) *UploadSampleV3Params {
	return &UploadSampleV3Params{
		HTTPClient: client,
	}
}

/* UploadSampleV3Params contains all the parameters to send to the API endpoint
   for the upload sample v3 operation.

   Typically these are written to a http.Request.
*/
type UploadSampleV3Params struct {

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

// WithDefaults hydrates default values in the upload sample v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadSampleV3Params) WithDefaults() *UploadSampleV3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload sample v3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadSampleV3Params) SetDefaults() {
	var (
		isConfidentialDefault = bool(true)
	)

	val := UploadSampleV3Params{
		IsConfidential: &isConfidentialDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the upload sample v3 params
func (o *UploadSampleV3Params) WithTimeout(timeout time.Duration) *UploadSampleV3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload sample v3 params
func (o *UploadSampleV3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload sample v3 params
func (o *UploadSampleV3Params) WithContext(ctx context.Context) *UploadSampleV3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload sample v3 params
func (o *UploadSampleV3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload sample v3 params
func (o *UploadSampleV3Params) WithHTTPClient(client *http.Client) *UploadSampleV3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload sample v3 params
func (o *UploadSampleV3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCSUSERUUID adds the xCSUSERUUID to the upload sample v3 params
func (o *UploadSampleV3Params) WithXCSUSERUUID(xCSUSERUUID *string) *UploadSampleV3Params {
	o.SetXCSUSERUUID(xCSUSERUUID)
	return o
}

// SetXCSUSERUUID adds the xCSUSERUuid to the upload sample v3 params
func (o *UploadSampleV3Params) SetXCSUSERUUID(xCSUSERUUID *string) {
	o.XCSUSERUUID = xCSUSERUUID
}

// WithBody adds the body to the upload sample v3 params
func (o *UploadSampleV3Params) WithBody(body []int64) *UploadSampleV3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload sample v3 params
func (o *UploadSampleV3Params) SetBody(body []int64) {
	o.Body = body
}

// WithComment adds the comment to the upload sample v3 params
func (o *UploadSampleV3Params) WithComment(comment *string) *UploadSampleV3Params {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the upload sample v3 params
func (o *UploadSampleV3Params) SetComment(comment *string) {
	o.Comment = comment
}

// WithFileName adds the fileName to the upload sample v3 params
func (o *UploadSampleV3Params) WithFileName(fileName string) *UploadSampleV3Params {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the upload sample v3 params
func (o *UploadSampleV3Params) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithIsConfidential adds the isConfidential to the upload sample v3 params
func (o *UploadSampleV3Params) WithIsConfidential(isConfidential *bool) *UploadSampleV3Params {
	o.SetIsConfidential(isConfidential)
	return o
}

// SetIsConfidential adds the isConfidential to the upload sample v3 params
func (o *UploadSampleV3Params) SetIsConfidential(isConfidential *bool) {
	o.IsConfidential = isConfidential
}

// WithUpfile adds the upfile to the upload sample v3 params
func (o *UploadSampleV3Params) WithUpfile(upfile runtime.NamedReadCloser) *UploadSampleV3Params {
	o.SetUpfile(upfile)
	return o
}

// SetUpfile adds the upfile to the upload sample v3 params
func (o *UploadSampleV3Params) SetUpfile(upfile runtime.NamedReadCloser) {
	o.Upfile = upfile
}

// WriteToRequest writes these params to a swagger request
func (o *UploadSampleV3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
