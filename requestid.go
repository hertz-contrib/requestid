/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package requestid

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/google/uuid"
)

var headerXRequestID string

// Option for request id generator
type Option func(*config)

type (
	Generator func() string
	Handler   func(ctx context.Context, c *app.RequestContext, requestID string)
)

// New initializes the RequestID middleware.
func New(opts ...Option) app.HandlerFunc {
	cfg := &config{
		generator: func() string {
			return uuid.New().String()
		},
		headerKey: "X-Request-ID",
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return func(ctx context.Context, c *app.RequestContext) {
		// Get id from request
		rid := c.Request.Header.Get(string(cfg.headerKey))
		if rid == "" {
			rid = cfg.generator()
		}
		headerXRequestID = string(cfg.headerKey)
		if cfg.handler != nil {
			cfg.handler(ctx, c, rid)
		}
		// Set the id to ensure that the request id is in the response
		c.Header(headerXRequestID, rid)
	}
}

type HeaderStrKey string

// WithGenerator set generator function
func WithGenerator(g Generator) Option {
	return func(cfg *config) {
		cfg.generator = g
	}
}

// WithCustomHeaderStrKey set custom header key for request id
func WithCustomHeaderStrKey(s HeaderStrKey) Option {
	return func(cfg *config) {
		cfg.headerKey = s
	}
}

// WithHandler set handler function for request id with context
func WithHandler(handler Handler) Option {
	return func(cfg *config) {
		cfg.handler = handler
	}
}

// Config defines the config for RequestID middleware
type config struct {
	// Generator defines a function to generate an ID.
	// Optional. Default: func() string {
	//   return uuid.New().String()
	// }
	generator Generator
	headerKey HeaderStrKey
	handler   Handler
}

// Get returns the request identifier
func Get(c *app.RequestContext) string {
	return c.Response.Header.Get(headerXRequestID)
}
