// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/ssoready/ssoready-go/core"
	organizations "github.com/ssoready/ssoready-go/management/organizations"
	samlconnections "github.com/ssoready/ssoready-go/management/samlconnections"
	scimdirectories "github.com/ssoready/ssoready-go/management/scimdirectories"
	setupurls "github.com/ssoready/ssoready-go/management/setupurls"
	option "github.com/ssoready/ssoready-go/option"
	http "net/http"
	os "os"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Organizations   *organizations.Client
	SAMLConnections *samlconnections.Client
	SCIMDirectories *scimdirectories.Client
	SetupURLs       *setupurls.Client
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
		header:          options.ToHeader(),
		Organizations:   organizations.NewClient(opts...),
		SAMLConnections: samlconnections.NewClient(opts...),
		SCIMDirectories: scimdirectories.NewClient(opts...),
		SetupURLs:       setupurls.NewClient(opts...),
	}
}
