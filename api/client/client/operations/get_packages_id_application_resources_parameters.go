// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetPackagesIDApplicationResourcesParams creates a new GetPackagesIDApplicationResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPackagesIDApplicationResourcesParams() *GetPackagesIDApplicationResourcesParams {
	return &GetPackagesIDApplicationResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPackagesIDApplicationResourcesParamsWithTimeout creates a new GetPackagesIDApplicationResourcesParams object
// with the ability to set a timeout on a request.
func NewGetPackagesIDApplicationResourcesParamsWithTimeout(timeout time.Duration) *GetPackagesIDApplicationResourcesParams {
	return &GetPackagesIDApplicationResourcesParams{
		timeout: timeout,
	}
}

// NewGetPackagesIDApplicationResourcesParamsWithContext creates a new GetPackagesIDApplicationResourcesParams object
// with the ability to set a context for a request.
func NewGetPackagesIDApplicationResourcesParamsWithContext(ctx context.Context) *GetPackagesIDApplicationResourcesParams {
	return &GetPackagesIDApplicationResourcesParams{
		Context: ctx,
	}
}

// NewGetPackagesIDApplicationResourcesParamsWithHTTPClient creates a new GetPackagesIDApplicationResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPackagesIDApplicationResourcesParamsWithHTTPClient(client *http.Client) *GetPackagesIDApplicationResourcesParams {
	return &GetPackagesIDApplicationResourcesParams{
		HTTPClient: client,
	}
}

/* GetPackagesIDApplicationResourcesParams contains all the parameters to send to the API endpoint
   for the get packages ID application resources operation.

   Typically these are written to a http.Request.
*/
type GetPackagesIDApplicationResourcesParams struct {

	// ID.
	ID string

	/* Page.

	   Page number of the query
	*/
	Page int64

	/* PageSize.

	   Maximum items to return
	*/
	PageSize int64

	// ReportingSBOMAnalyzersContainElements.
	ReportingSBOMAnalyzersContainElements []string

	// ReportingSBOMAnalyzersDoesntContainElements.
	ReportingSBOMAnalyzersDoesntContainElements []string

	// ResourceHashContains.
	ResourceHashContains []string

	// ResourceHashEnd.
	ResourceHashEnd *string

	// ResourceHashIsNot.
	ResourceHashIsNot []string

	// ResourceHashIs.
	ResourceHashIs []string

	// ResourceHashStart.
	ResourceHashStart *string

	// ResourceNameContains.
	ResourceNameContains []string

	// ResourceNameEnd.
	ResourceNameEnd *string

	// ResourceNameIsNot.
	ResourceNameIsNot []string

	// ResourceNameIs.
	ResourceNameIs []string

	// ResourceNameStart.
	ResourceNameStart *string

	/* SortDir.

	   Sorting direction

	   Default: "ASC"
	*/
	SortDir *string

	/* SortKey.

	   Sort key
	*/
	SortKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get packages ID application resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackagesIDApplicationResourcesParams) WithDefaults() *GetPackagesIDApplicationResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get packages ID application resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPackagesIDApplicationResourcesParams) SetDefaults() {
	var (
		sortDirDefault = string("ASC")
	)

	val := GetPackagesIDApplicationResourcesParams{
		SortDir: &sortDirDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithTimeout(timeout time.Duration) *GetPackagesIDApplicationResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithContext(ctx context.Context) *GetPackagesIDApplicationResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithHTTPClient(client *http.Client) *GetPackagesIDApplicationResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithID(id string) *GetPackagesIDApplicationResourcesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetID(id string) {
	o.ID = id
}

// WithPage adds the page to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithPage(page int64) *GetPackagesIDApplicationResourcesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetPage(page int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithPageSize(pageSize int64) *GetPackagesIDApplicationResourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetPageSize(pageSize int64) {
	o.PageSize = pageSize
}

// WithReportingSBOMAnalyzersContainElements adds the reportingSBOMAnalyzersContainElements to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithReportingSBOMAnalyzersContainElements(reportingSBOMAnalyzersContainElements []string) *GetPackagesIDApplicationResourcesParams {
	o.SetReportingSBOMAnalyzersContainElements(reportingSBOMAnalyzersContainElements)
	return o
}

// SetReportingSBOMAnalyzersContainElements adds the reportingSBOMAnalyzersContainElements to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetReportingSBOMAnalyzersContainElements(reportingSBOMAnalyzersContainElements []string) {
	o.ReportingSBOMAnalyzersContainElements = reportingSBOMAnalyzersContainElements
}

// WithReportingSBOMAnalyzersDoesntContainElements adds the reportingSBOMAnalyzersDoesntContainElements to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithReportingSBOMAnalyzersDoesntContainElements(reportingSBOMAnalyzersDoesntContainElements []string) *GetPackagesIDApplicationResourcesParams {
	o.SetReportingSBOMAnalyzersDoesntContainElements(reportingSBOMAnalyzersDoesntContainElements)
	return o
}

// SetReportingSBOMAnalyzersDoesntContainElements adds the reportingSBOMAnalyzersDoesntContainElements to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetReportingSBOMAnalyzersDoesntContainElements(reportingSBOMAnalyzersDoesntContainElements []string) {
	o.ReportingSBOMAnalyzersDoesntContainElements = reportingSBOMAnalyzersDoesntContainElements
}

// WithResourceHashContains adds the resourceHashContains to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceHashContains(resourceHashContains []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceHashContains(resourceHashContains)
	return o
}

// SetResourceHashContains adds the resourceHashContains to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceHashContains(resourceHashContains []string) {
	o.ResourceHashContains = resourceHashContains
}

// WithResourceHashEnd adds the resourceHashEnd to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceHashEnd(resourceHashEnd *string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceHashEnd(resourceHashEnd)
	return o
}

// SetResourceHashEnd adds the resourceHashEnd to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceHashEnd(resourceHashEnd *string) {
	o.ResourceHashEnd = resourceHashEnd
}

// WithResourceHashIsNot adds the resourceHashIsNot to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceHashIsNot(resourceHashIsNot []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceHashIsNot(resourceHashIsNot)
	return o
}

// SetResourceHashIsNot adds the resourceHashIsNot to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceHashIsNot(resourceHashIsNot []string) {
	o.ResourceHashIsNot = resourceHashIsNot
}

// WithResourceHashIs adds the resourceHashIs to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceHashIs(resourceHashIs []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceHashIs(resourceHashIs)
	return o
}

// SetResourceHashIs adds the resourceHashIs to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceHashIs(resourceHashIs []string) {
	o.ResourceHashIs = resourceHashIs
}

// WithResourceHashStart adds the resourceHashStart to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceHashStart(resourceHashStart *string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceHashStart(resourceHashStart)
	return o
}

// SetResourceHashStart adds the resourceHashStart to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceHashStart(resourceHashStart *string) {
	o.ResourceHashStart = resourceHashStart
}

// WithResourceNameContains adds the resourceNameContains to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceNameContains(resourceNameContains []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceNameContains(resourceNameContains)
	return o
}

// SetResourceNameContains adds the resourceNameContains to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceNameContains(resourceNameContains []string) {
	o.ResourceNameContains = resourceNameContains
}

// WithResourceNameEnd adds the resourceNameEnd to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceNameEnd(resourceNameEnd *string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceNameEnd(resourceNameEnd)
	return o
}

// SetResourceNameEnd adds the resourceNameEnd to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceNameEnd(resourceNameEnd *string) {
	o.ResourceNameEnd = resourceNameEnd
}

// WithResourceNameIsNot adds the resourceNameIsNot to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceNameIsNot(resourceNameIsNot []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceNameIsNot(resourceNameIsNot)
	return o
}

// SetResourceNameIsNot adds the resourceNameIsNot to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceNameIsNot(resourceNameIsNot []string) {
	o.ResourceNameIsNot = resourceNameIsNot
}

// WithResourceNameIs adds the resourceNameIs to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceNameIs(resourceNameIs []string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceNameIs(resourceNameIs)
	return o
}

// SetResourceNameIs adds the resourceNameIs to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceNameIs(resourceNameIs []string) {
	o.ResourceNameIs = resourceNameIs
}

// WithResourceNameStart adds the resourceNameStart to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithResourceNameStart(resourceNameStart *string) *GetPackagesIDApplicationResourcesParams {
	o.SetResourceNameStart(resourceNameStart)
	return o
}

// SetResourceNameStart adds the resourceNameStart to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetResourceNameStart(resourceNameStart *string) {
	o.ResourceNameStart = resourceNameStart
}

// WithSortDir adds the sortDir to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithSortDir(sortDir *string) *GetPackagesIDApplicationResourcesParams {
	o.SetSortDir(sortDir)
	return o
}

// SetSortDir adds the sortDir to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetSortDir(sortDir *string) {
	o.SortDir = sortDir
}

// WithSortKey adds the sortKey to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) WithSortKey(sortKey string) *GetPackagesIDApplicationResourcesParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the get packages ID application resources params
func (o *GetPackagesIDApplicationResourcesParams) SetSortKey(sortKey string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetPackagesIDApplicationResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param page
	qrPage := o.Page
	qPage := swag.FormatInt64(qrPage)
	if qPage != "" {

		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	// query param pageSize
	qrPageSize := o.PageSize
	qPageSize := swag.FormatInt64(qrPageSize)
	if qPageSize != "" {

		if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
			return err
		}
	}

	if o.ReportingSBOMAnalyzersContainElements != nil {

		// binding items for reportingSBOMAnalyzers[containElements]
		joinedReportingSBOMAnalyzersContainElements := o.bindParamReportingSBOMAnalyzersContainElements(reg)

		// query array param reportingSBOMAnalyzers[containElements]
		if err := r.SetQueryParam("reportingSBOMAnalyzers[containElements]", joinedReportingSBOMAnalyzersContainElements...); err != nil {
			return err
		}
	}

	if o.ReportingSBOMAnalyzersDoesntContainElements != nil {

		// binding items for reportingSBOMAnalyzers[doesntContainElements]
		joinedReportingSBOMAnalyzersDoesntContainElements := o.bindParamReportingSBOMAnalyzersDoesntContainElements(reg)

		// query array param reportingSBOMAnalyzers[doesntContainElements]
		if err := r.SetQueryParam("reportingSBOMAnalyzers[doesntContainElements]", joinedReportingSBOMAnalyzersDoesntContainElements...); err != nil {
			return err
		}
	}

	if o.ResourceHashContains != nil {

		// binding items for resourceHash[contains]
		joinedResourceHashContains := o.bindParamResourceHashContains(reg)

		// query array param resourceHash[contains]
		if err := r.SetQueryParam("resourceHash[contains]", joinedResourceHashContains...); err != nil {
			return err
		}
	}

	if o.ResourceHashEnd != nil {

		// query param resourceHash[end]
		var qrResourceHashEnd string

		if o.ResourceHashEnd != nil {
			qrResourceHashEnd = *o.ResourceHashEnd
		}
		qResourceHashEnd := qrResourceHashEnd
		if qResourceHashEnd != "" {

			if err := r.SetQueryParam("resourceHash[end]", qResourceHashEnd); err != nil {
				return err
			}
		}
	}

	if o.ResourceHashIsNot != nil {

		// binding items for resourceHash[isNot]
		joinedResourceHashIsNot := o.bindParamResourceHashIsNot(reg)

		// query array param resourceHash[isNot]
		if err := r.SetQueryParam("resourceHash[isNot]", joinedResourceHashIsNot...); err != nil {
			return err
		}
	}

	if o.ResourceHashIs != nil {

		// binding items for resourceHash[is]
		joinedResourceHashIs := o.bindParamResourceHashIs(reg)

		// query array param resourceHash[is]
		if err := r.SetQueryParam("resourceHash[is]", joinedResourceHashIs...); err != nil {
			return err
		}
	}

	if o.ResourceHashStart != nil {

		// query param resourceHash[start]
		var qrResourceHashStart string

		if o.ResourceHashStart != nil {
			qrResourceHashStart = *o.ResourceHashStart
		}
		qResourceHashStart := qrResourceHashStart
		if qResourceHashStart != "" {

			if err := r.SetQueryParam("resourceHash[start]", qResourceHashStart); err != nil {
				return err
			}
		}
	}

	if o.ResourceNameContains != nil {

		// binding items for resourceName[contains]
		joinedResourceNameContains := o.bindParamResourceNameContains(reg)

		// query array param resourceName[contains]
		if err := r.SetQueryParam("resourceName[contains]", joinedResourceNameContains...); err != nil {
			return err
		}
	}

	if o.ResourceNameEnd != nil {

		// query param resourceName[end]
		var qrResourceNameEnd string

		if o.ResourceNameEnd != nil {
			qrResourceNameEnd = *o.ResourceNameEnd
		}
		qResourceNameEnd := qrResourceNameEnd
		if qResourceNameEnd != "" {

			if err := r.SetQueryParam("resourceName[end]", qResourceNameEnd); err != nil {
				return err
			}
		}
	}

	if o.ResourceNameIsNot != nil {

		// binding items for resourceName[isNot]
		joinedResourceNameIsNot := o.bindParamResourceNameIsNot(reg)

		// query array param resourceName[isNot]
		if err := r.SetQueryParam("resourceName[isNot]", joinedResourceNameIsNot...); err != nil {
			return err
		}
	}

	if o.ResourceNameIs != nil {

		// binding items for resourceName[is]
		joinedResourceNameIs := o.bindParamResourceNameIs(reg)

		// query array param resourceName[is]
		if err := r.SetQueryParam("resourceName[is]", joinedResourceNameIs...); err != nil {
			return err
		}
	}

	if o.ResourceNameStart != nil {

		// query param resourceName[start]
		var qrResourceNameStart string

		if o.ResourceNameStart != nil {
			qrResourceNameStart = *o.ResourceNameStart
		}
		qResourceNameStart := qrResourceNameStart
		if qResourceNameStart != "" {

			if err := r.SetQueryParam("resourceName[start]", qResourceNameStart); err != nil {
				return err
			}
		}
	}

	if o.SortDir != nil {

		// query param sortDir
		var qrSortDir string

		if o.SortDir != nil {
			qrSortDir = *o.SortDir
		}
		qSortDir := qrSortDir
		if qSortDir != "" {

			if err := r.SetQueryParam("sortDir", qSortDir); err != nil {
				return err
			}
		}
	}

	// query param sortKey
	qrSortKey := o.SortKey
	qSortKey := qrSortKey
	if qSortKey != "" {

		if err := r.SetQueryParam("sortKey", qSortKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPackagesIDApplicationResources binds the parameter reportingSBOMAnalyzers[containElements]
func (o *GetPackagesIDApplicationResourcesParams) bindParamReportingSBOMAnalyzersContainElements(formats strfmt.Registry) []string {
	reportingSBOMAnalyzersContainElementsIR := o.ReportingSBOMAnalyzersContainElements

	var reportingSBOMAnalyzersContainElementsIC []string
	for _, reportingSBOMAnalyzersContainElementsIIR := range reportingSBOMAnalyzersContainElementsIR { // explode []string

		reportingSBOMAnalyzersContainElementsIIV := reportingSBOMAnalyzersContainElementsIIR // string as string
		reportingSBOMAnalyzersContainElementsIC = append(reportingSBOMAnalyzersContainElementsIC, reportingSBOMAnalyzersContainElementsIIV)
	}

	// items.CollectionFormat: ""
	reportingSBOMAnalyzersContainElementsIS := swag.JoinByFormat(reportingSBOMAnalyzersContainElementsIC, "")

	return reportingSBOMAnalyzersContainElementsIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter reportingSBOMAnalyzers[doesntContainElements]
func (o *GetPackagesIDApplicationResourcesParams) bindParamReportingSBOMAnalyzersDoesntContainElements(formats strfmt.Registry) []string {
	reportingSBOMAnalyzersDoesntContainElementsIR := o.ReportingSBOMAnalyzersDoesntContainElements

	var reportingSBOMAnalyzersDoesntContainElementsIC []string
	for _, reportingSBOMAnalyzersDoesntContainElementsIIR := range reportingSBOMAnalyzersDoesntContainElementsIR { // explode []string

		reportingSBOMAnalyzersDoesntContainElementsIIV := reportingSBOMAnalyzersDoesntContainElementsIIR // string as string
		reportingSBOMAnalyzersDoesntContainElementsIC = append(reportingSBOMAnalyzersDoesntContainElementsIC, reportingSBOMAnalyzersDoesntContainElementsIIV)
	}

	// items.CollectionFormat: ""
	reportingSBOMAnalyzersDoesntContainElementsIS := swag.JoinByFormat(reportingSBOMAnalyzersDoesntContainElementsIC, "")

	return reportingSBOMAnalyzersDoesntContainElementsIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceHash[contains]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceHashContains(formats strfmt.Registry) []string {
	resourceHashContainsIR := o.ResourceHashContains

	var resourceHashContainsIC []string
	for _, resourceHashContainsIIR := range resourceHashContainsIR { // explode []string

		resourceHashContainsIIV := resourceHashContainsIIR // string as string
		resourceHashContainsIC = append(resourceHashContainsIC, resourceHashContainsIIV)
	}

	// items.CollectionFormat: ""
	resourceHashContainsIS := swag.JoinByFormat(resourceHashContainsIC, "")

	return resourceHashContainsIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceHash[isNot]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceHashIsNot(formats strfmt.Registry) []string {
	resourceHashIsNotIR := o.ResourceHashIsNot

	var resourceHashIsNotIC []string
	for _, resourceHashIsNotIIR := range resourceHashIsNotIR { // explode []string

		resourceHashIsNotIIV := resourceHashIsNotIIR // string as string
		resourceHashIsNotIC = append(resourceHashIsNotIC, resourceHashIsNotIIV)
	}

	// items.CollectionFormat: ""
	resourceHashIsNotIS := swag.JoinByFormat(resourceHashIsNotIC, "")

	return resourceHashIsNotIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceHash[is]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceHashIs(formats strfmt.Registry) []string {
	resourceHashIsIR := o.ResourceHashIs

	var resourceHashIsIC []string
	for _, resourceHashIsIIR := range resourceHashIsIR { // explode []string

		resourceHashIsIIV := resourceHashIsIIR // string as string
		resourceHashIsIC = append(resourceHashIsIC, resourceHashIsIIV)
	}

	// items.CollectionFormat: ""
	resourceHashIsIS := swag.JoinByFormat(resourceHashIsIC, "")

	return resourceHashIsIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceName[contains]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceNameContains(formats strfmt.Registry) []string {
	resourceNameContainsIR := o.ResourceNameContains

	var resourceNameContainsIC []string
	for _, resourceNameContainsIIR := range resourceNameContainsIR { // explode []string

		resourceNameContainsIIV := resourceNameContainsIIR // string as string
		resourceNameContainsIC = append(resourceNameContainsIC, resourceNameContainsIIV)
	}

	// items.CollectionFormat: ""
	resourceNameContainsIS := swag.JoinByFormat(resourceNameContainsIC, "")

	return resourceNameContainsIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceName[isNot]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceNameIsNot(formats strfmt.Registry) []string {
	resourceNameIsNotIR := o.ResourceNameIsNot

	var resourceNameIsNotIC []string
	for _, resourceNameIsNotIIR := range resourceNameIsNotIR { // explode []string

		resourceNameIsNotIIV := resourceNameIsNotIIR // string as string
		resourceNameIsNotIC = append(resourceNameIsNotIC, resourceNameIsNotIIV)
	}

	// items.CollectionFormat: ""
	resourceNameIsNotIS := swag.JoinByFormat(resourceNameIsNotIC, "")

	return resourceNameIsNotIS
}

// bindParamGetPackagesIDApplicationResources binds the parameter resourceName[is]
func (o *GetPackagesIDApplicationResourcesParams) bindParamResourceNameIs(formats strfmt.Registry) []string {
	resourceNameIsIR := o.ResourceNameIs

	var resourceNameIsIC []string
	for _, resourceNameIsIIR := range resourceNameIsIR { // explode []string

		resourceNameIsIIV := resourceNameIsIIR // string as string
		resourceNameIsIC = append(resourceNameIsIC, resourceNameIsIIV)
	}

	// items.CollectionFormat: ""
	resourceNameIsIS := swag.JoinByFormat(resourceNameIsIC, "")

	return resourceNameIsIS
}
