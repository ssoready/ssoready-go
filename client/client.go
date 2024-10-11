// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/ssoready/ssoready-go/core"
	managementclient "github.com/ssoready/ssoready-go/management/client"
	option "github.com/ssoready/ssoready-go/option"
	saml "github.com/ssoready/ssoready-go/saml"
	scim "github.com/ssoready/ssoready-go/scim"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	SAML       *saml.Client
	Scim       *scim.Client
	Management *managementclient.Client
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
		header:     options.ToHeader(),
		SAML:       saml.NewClient(opts...),
		Scim:       scim.NewClient(opts...),
		Management: managementclient.NewClient(opts...),
	}
}