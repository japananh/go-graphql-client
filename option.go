package graphql

import (
	"net/http"
	"time"
)

// ClientOption is used to configure client with options.
type ClientOption func(c *Client)

// WithRetry creates an option to indicate the number of retries.
func WithRetry(maxRetries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = maxRetries
	}
}

// WithRetryBaseDelay creates an option to indicate the base delay factor of retries.
func WithRetryBaseDelay(delay time.Duration) ClientOption {
	return func(c *Client) {
		c.retryBaseDelay = delay
	}
}

// WithRetryExponentialRate creates an option to indicate the exponential rate of retries.
func WithRetryExponentialRate(rate float64) ClientOption {
	return func(c *Client) {
		c.retryExponentialRate = rate
	}
}

// WithRetryHTTPStatus creates an option to retry if the HTTP response status is in the status slice.
func WithRetryHTTPStatus(status []int) ClientOption {
	return func(c *Client) {
		c.retryHttpStatus = status
	}
}

// WithRetryOnGraphQLError creates a callback option to check if the graphql error is retryable.
func WithRetryOnGraphQLError(callback func(errs Errors) bool) ClientOption {
	return func(c *Client) {
		c.retryOnGraphQLError = callback
	}
}

// OptionType represents the logic of graphql query construction.
type OptionType string

const (
	OptionTypeOperationDirective OptionType = "operation_directive"
)

// They are optional parts. By default GraphQL queries can request data without them.
type Option interface {
	// Type returns the supported type of the renderer
	// available types: operation_name and operation_directive
	Type() OptionType
}

// operationNameOption represents the operation name render component.
type operationNameOption struct {
	name string
}

func (ono operationNameOption) Type() OptionType {
	return "operation_name"
}

func (ono operationNameOption) String() string {
	return ono.name
}

// OperationName creates the operation name option.
func OperationName(name string) Option {
	return operationNameOption{name}
}

// bind the struct pointer to decode extensions from response.
type bindExtensionsOption struct {
	value any
}

func (ono bindExtensionsOption) Type() OptionType {
	return "bind_extensions"
}

// BindExtensions bind the struct pointer to decode extensions from json response.
func BindExtensions(value any) Option {
	return bindExtensionsOption{value: value}
}

// bind the struct pointer to return headers from response.
type bindResponseHeadersOption struct {
	value *http.Header
}

func (ono bindResponseHeadersOption) Type() OptionType {
	return "bind_response_headers"
}

// BindExtensionsBindResponseHeaders bind the header response to the pointer.
func BindResponseHeaders(value *http.Header) Option {
	return bindResponseHeadersOption{value: value}
}
