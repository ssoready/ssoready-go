// This file was auto-generated by Fern from our API Definition.

package samlconnections

import (
	context "context"
	ssoreadygo "github.com/ssoready/ssoready-go"
	core "github.com/ssoready/ssoready-go/core"
	management "github.com/ssoready/ssoready-go/management"
	option "github.com/ssoready/ssoready-go/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	if options.APIKey == "" {
		options.APIKey = os.Getenv("SSOREADY_API_KEY")
	}
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Lists SAML connections in an organization.
func (c *Client) ListSAMLConnections(
	ctx context.Context,
	request *management.SAMLConnectionsListSAMLConnectionsRequest,
	opts ...option.RequestOption,
) (*ssoreadygo.ListSAMLConnectionsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/saml-connections"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *ssoreadygo.ListSAMLConnectionsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Creates a SAML connection.
func (c *Client) CreateSAMLConnection(
	ctx context.Context,
	request *ssoreadygo.SAMLConnection,
	opts ...option.RequestOption,
) (*ssoreadygo.CreateSAMLConnectionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/saml-connections"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *ssoreadygo.CreateSAMLConnectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Gets a SAML connection.
func (c *Client) GetSAMLConnection(
	ctx context.Context,
	// ID of the SAML connection to get.
	id string,
	opts ...option.RequestOption,
) (*ssoreadygo.GetSAMLConnectionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/saml-connections/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *ssoreadygo.GetSAMLConnectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Updates a SAML connection.
func (c *Client) UpdateSAMLConnection(
	ctx context.Context,
	// The ID of the SAML connection to update.
	id string,
	request *ssoreadygo.SAMLConnection,
	opts ...option.RequestOption,
) (*ssoreadygo.UpdateSAMLConnectionResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := core.EncodeURL(baseURL+"/v1/saml-connections/%v", id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *ssoreadygo.UpdateSAMLConnectionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			MaxAttempts:     options.MaxAttempts,
			Headers:         headers,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}