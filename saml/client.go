// This file was auto-generated by Fern from our API Definition.

package saml

import (
	context "context"
	ssoreadygo "github.com/ssoready/ssoready-go"
	core "github.com/ssoready/ssoready-go/core"
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

// Exchanges a SAML access code for details about your user's SAML login details.
func (c *Client) RedeemSAMLAccessCode(
	ctx context.Context,
	request *ssoreadygo.RedeemSAMLAccessCodeRequest,
	opts ...option.RequestOption,
) (*ssoreadygo.RedeemSAMLAccessCodeResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/saml/redeem"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *ssoreadygo.RedeemSAMLAccessCodeResponse
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

// Gets a SAML initiation URL to redirect your users to.
func (c *Client) GetSAMLRedirectURL(
	ctx context.Context,
	request *ssoreadygo.GetSAMLRedirectURLRequest,
	opts ...option.RequestOption,
) (*ssoreadygo.GetSAMLRedirectURLResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.ssoready.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/v1/saml/redirect"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())
	headers.Set("Content-Type", "application/json")

	var response *ssoreadygo.GetSAMLRedirectURLResponse
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