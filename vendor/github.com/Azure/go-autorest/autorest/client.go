package autorest

// Copyright 2017 Microsoft Corporation
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"

	"github.com/Azure/go-autorest/logger"
	"github.com/Azure/go-autorest/tracing"
)

const (
	// DefaultPollingDelay is a reasonable delay between polling requests.
	DefaultPollingDelay = 60 * time.Second

	// DefaultPollingDuration is a reasonable total polling duration.
	DefaultPollingDuration = 15 * time.Minute

	// DefaultRetryAttempts is number of attempts for retry status codes (5xx).
	DefaultRetryAttempts = 3

	// DefaultRetryDuration is the duration to wait between retries.
	DefaultRetryDuration = 30 * time.Second
)

var (
	// StatusCodesForRetry are a defined group of status code for which the client will retry
	StatusCodesForRetry = []int{
		http.StatusRequestTimeout,      // 408
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout,      // 504
	}
)

const (
	requestFormat = `HTTP Request Begin ===================================================
%s
===================================================== HTTP Request End
`
	responseFormat = `HTTP Response Begin ===================================================
%s
===================================================== HTTP Response End
`
)

// Response serves as the base for all responses from generated clients. It provides access to the
// last http.Response.
type Response struct {
	*http.Response `json:"-"`
}

// LoggingInspector implements request and response inspectors that log the full request and
// response to a supplied log.
type LoggingInspector struct {
	Logger *log.Logger
}

// WithInspection returns a PrepareDecorator that emits the http.Request to the supplied logger. The
// body is restored after being emitted.
//
// Note: Since it reads the entire Body, this decorator should not be used where body streaming is
// important. It is best used to trace JSON or similar body values.
func (li LoggingInspector) WithInspection() PrepareDecorator {
	return func(p Preparer) Preparer {
		return PreparerFunc(func(r *http.Request) (*http.Request, error) {
			var body, b bytes.Buffer

			defer r.Body.Close()

			r.Body = ioutil.NopCloser(io.TeeReader(r.Body, &body))
			if err := r.Write(&b); err != nil {
				return nil, fmt.Errorf("Failed to write response: %v", err)
			}

			li.Logger.Printf(requestFormat, b.String())

			r.Body = ioutil.NopCloser(&body)
			return p.Prepare(r)
		})
	}
}

// ByInspecting returns a RespondDecorator that emits the http.Response to the supplied logger. The
// body is restored after being emitted.
//
// Note: Since it reads the entire Body, this decorator should not be used where body streaming is
// important. It is best used to trace JSON or similar body values.
func (li LoggingInspector) ByInspecting() RespondDecorator {
	return func(r Responder) Responder {
		return ResponderFunc(func(resp *http.Response) error {
			var body, b bytes.Buffer
			defer resp.Body.Close()
			resp.Body = ioutil.NopCloser(io.TeeReader(resp.Body, &body))
			if err := resp.Write(&b); err != nil {
				return fmt.Errorf("Failed to write response: %v", err)
			}

			li.Logger.Printf(responseFormat, b.String())

			resp.Body = ioutil.NopCloser(&body)
			return r.Respond(resp)
		})
	}
}

// Client is the base for autorest generated clients. It provides default, "do nothing"
// implementations of an Authorizer, RequestInspector, and ResponseInspector. It also returns the
// standard, undecorated http.Client as a default Sender.
//
// Generated clients should also use Error (see NewError and NewErrorWithError) for errors and
// return responses that compose with Response.
//
// Most customization of generated clients is best achieved by supplying a custom Authorizer, custom
// RequestInspector, and / or custom ResponseInspector. Users may log requests, implement circuit
// breakers (see https://msdn.microsoft.com/en-us/library/dn589784.aspx) or otherwise influence
// sending the request by providing a decorated Sender.
type Client struct {
	Authorizer        Authorizer
	Sender            Sender
	RequestInspector  PrepareDecorator
	ResponseInspector RespondDecorator

	// PollingDelay sets the polling frequency used in absence of a Retry-After HTTP header
	PollingDelay time.Duration

	// PollingDuration sets the maximum polling time after which an error is returned.
	// Setting this to zero will use the provided context to control the duration.
	PollingDuration time.Duration

	// RetryAttempts sets the default number of retry attempts for client.
	RetryAttempts int

	// RetryDuration sets the delay duration for retries.
	RetryDuration time.Duration

	// UserAgent, if not empty, will be set as the HTTP User-Agent header on all requests sent
	// through the Do method.
	UserAgent string

	Jar http.CookieJar

	// Set to true to skip attempted registration of resource providers (false by default).
	SkipResourceProviderRegistration bool
}

// NewClientWithUserAgent returns an instance of a Client with the UserAgent set to the passed
// string.
func NewClientWithUserAgent(ua string) Client {
<<<<<<< HEAD
=======
	return newClient(ua, tls.RenegotiateNever)
}

// ClientOptions contains various Client configuration options.
type ClientOptions struct {
	// UserAgent is an optional user-agent string to append to the default user agent.
	UserAgent string

	// Renegotiation is an optional setting to control client-side TLS renegotiation.
	Renegotiation tls.RenegotiationSupport
}

// NewClientWithOptions returns an instance of a Client with the specified values.
func NewClientWithOptions(options ClientOptions) Client {
	return newClient(options.UserAgent, options.Renegotiation)
}

func newClient(ua string, renegotiation tls.RenegotiationSupport) Client {
>>>>>>> v0.0.4
	c := Client{
		PollingDelay:    DefaultPollingDelay,
		PollingDuration: DefaultPollingDuration,
		RetryAttempts:   DefaultRetryAttempts,
		RetryDuration:   DefaultRetryDuration,
		UserAgent:       UserAgent(),
	}
<<<<<<< HEAD
	c.Sender = c.sender()
=======
	c.Sender = c.sender(renegotiation)
>>>>>>> v0.0.4
	c.AddToUserAgent(ua)
	return c
}

// AddToUserAgent adds an extension to the current user agent
func (c *Client) AddToUserAgent(extension string) error {
	if extension != "" {
		c.UserAgent = fmt.Sprintf("%s %s", c.UserAgent, extension)
		return nil
	}
	return fmt.Errorf("Extension was empty, User Agent stayed as %s", c.UserAgent)
}

// Do implements the Sender interface by invoking the active Sender after applying authorization.
// If Sender is not set, it uses a new instance of http.Client. In both cases it will, if UserAgent
// is set, apply set the User-Agent header.
func (c Client) Do(r *http.Request) (*http.Response, error) {
	if r.UserAgent() == "" {
		r, _ = Prepare(r,
			WithUserAgent(c.UserAgent))
	}
	// NOTE: c.WithInspection() must be last in the list so that it can inspect all preceding operations
	r, err := Prepare(r,
		c.WithAuthorization(),
		c.WithInspection())
	if err != nil {
		var resp *http.Response
		if detErr, ok := err.(DetailedError); ok {
			// if the authorization failed (e.g. invalid credentials) there will
			// be a response associated with the error, be sure to return it.
			resp = detErr.Response
		}
		return resp, NewErrorWithError(err, "autorest/Client", "Do", nil, "Preparing request failed")
	}
	logger.Instance.WriteRequest(r, logger.Filter{
		Header: func(k string, v []string) (bool, []string) {
			// remove the auth token from the log
			if strings.EqualFold(k, "Authorization") || strings.EqualFold(k, "Ocp-Apim-Subscription-Key") {
				v = []string{"**REDACTED**"}
			}
			return true, v
		},
	})
<<<<<<< HEAD
	resp, err := SendWithSender(c.sender(), r)
=======
	resp, err := SendWithSender(c.sender(tls.RenegotiateNever), r)
>>>>>>> v0.0.4
	logger.Instance.WriteResponse(resp, logger.Filter{})
	Respond(resp, c.ByInspecting())
	return resp, err
}

// sender returns the Sender to which to send requests.
<<<<<<< HEAD
func (c Client) sender() Sender {
	if c.Sender == nil {
		j, _ := cookiejar.New(nil)
		tracing.Transport.Base = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		}
		client := &http.Client{Jar: j, Transport: tracing.Transport}
		return client
=======
func (c Client) sender(renengotiation tls.RenegotiationSupport) Sender {
	if c.Sender == nil {
		// Use behaviour compatible with DefaultTransport, but require TLS minimum version.
		var defaultTransport = http.DefaultTransport.(*http.Transport)
		transport := tracing.Transport
		// for non-default values of TLS renegotiation create a new tracing transport.
		// updating tracing.Transport affects all clients which is not what we want.
		if renengotiation != tls.RenegotiateNever {
			transport = tracing.NewTransport()
		}
		transport.Base = &http.Transport{
			Proxy:                 defaultTransport.Proxy,
			DialContext:           defaultTransport.DialContext,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
			TLSClientConfig: &tls.Config{
				MinVersion:    tls.VersionTLS12,
				Renegotiation: renengotiation,
			},
		}
		j, _ := cookiejar.New(nil)
		return &http.Client{Jar: j, Transport: transport}
>>>>>>> v0.0.4
	}

	return c.Sender
}

// WithAuthorization is a convenience method that returns the WithAuthorization PrepareDecorator
// from the current Authorizer. If not Authorizer is set, it uses the NullAuthorizer.
func (c Client) WithAuthorization() PrepareDecorator {
	return c.authorizer().WithAuthorization()
}

// authorizer returns the Authorizer to use.
func (c Client) authorizer() Authorizer {
	if c.Authorizer == nil {
		return NullAuthorizer{}
	}
	return c.Authorizer
}

// WithInspection is a convenience method that passes the request to the supplied RequestInspector,
// if present, or returns the WithNothing PrepareDecorator otherwise.
func (c Client) WithInspection() PrepareDecorator {
	if c.RequestInspector == nil {
		return WithNothing()
	}
	return c.RequestInspector
}

// ByInspecting is a convenience method that passes the response to the supplied ResponseInspector,
// if present, or returns the ByIgnoring RespondDecorator otherwise.
func (c Client) ByInspecting() RespondDecorator {
	if c.ResponseInspector == nil {
		return ByIgnoring()
	}
	return c.ResponseInspector
}
