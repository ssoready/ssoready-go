# SSOReady Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)

The SSOReady Go library provides convenient access to the SSOReady API from Go.

## Requirements

This module requires Go version >= 1.18.

# Installation

Run the following command to use the ssoready Go library in your module:

```sh
go get github.com/ssoready/ssoready-go
```

## Usage

```go
import (
  "github.com/ssoready/ssoready-go"
  ssoreadyclient "github.com/ssoready/ssoready-go/client"
  "github.com/ssoready/ssoready-go/option"
)

client := ssoreadyclient.NewClient(
  option.WithAPIKey("<YOUR_API_KEY>"),
)

response, err := client.SAML.RedeemSAMLAccessCode(
  context.TODO(),
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
)
```

## Optional Parameters

This library models optional primitives and enum types as pointers. This is primarily meant to distinguish
default zero values from explicit values (e.g. `false` for `bool` and `""` for `string`). A collection of
helper functions are provided to easily map a primitive or enum to its pointer-equivalent (e.g. `ssoready.String`).

## Timeouts

Setting a timeout for each individual request is as simple as using the standard `context` library. Setting
a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
)
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := ssoreadyclient.NewClient(
  option.WithAPIKey("<YOUR_API_KEY>"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

These request options can either be specified on the client so that they're applied on _every_
request (shown above), or for an individual request like so:

```go
response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
  option.WithAPIKey("<YOUR_API_KEY>"),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Automatic Retries

The ssoready Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := ssoreadyclient.NewClient(
  option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
  option.WithMaxAttempts(1),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to an unauthorized request (i.e. status code 401) with the following:

```go
response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
)
if err != nil {
  if unauthorizedErr, ok := err.(*ssoready.UnauthorizedError);
    // Do something with the unauthorized request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
)
if err != nil {
  var unauthorizedErr *ssoready.UnauthorizedError
  if errors.As(err, unauthorizedErr) {
    // Do something with the unauthorized request ...
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability
to access the type with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
response, err := client.SAML.RedeemSAMLAccessCode(
  ctx,
  &ssoready.RedeemSAMLAccessCodeRequest{
    SAMLAccessCode: "saml_access_code...",
  },
)
if err != nil {
  return fmt.Errorf("failed to create application: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
