// This file was auto-generated by Fern from our API Definition.

package management

type SAMLConnectionsListSAMLConnectionsRequest struct {
	// The organization the SAML connections belong to.
	OrganizationID *string `json:"-" url:"organizationId,omitempty"`
	// Pagination token. Leave empty to get the first page of results.
	PageToken *string `json:"-" url:"pageToken,omitempty"`
}
