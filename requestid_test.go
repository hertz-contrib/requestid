package requestid

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	hzconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/cloudwego/hertz/pkg/route"
	"net/http"
	"testing"
)

const testXRequestID = "test-request-id"
const customHeaderKey = "customKey"

func emptySuccessResponse(ctx context.Context, c *app.RequestContext) {
	c.String(http.StatusOK, "")
}

func hertzHandler(middleware app.HandlerFunc) *route.Engine {
	r := route.NewEngine(hzconfig.NewOptions([]hzconfig.Option{}))
	r.Use(middleware)
	r.GET("/", emptySuccessResponse)

	return r
}

func TestCreateNewRequestID(t *testing.T) {
	r := hertzHandler(New())
	w := ut.PerformRequest(r, http.MethodGet, "/", nil)

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.NotEqual(t, "", string(w.Header().Peek(headerXRequestID)))
}

func TestPassThruRequestID(t *testing.T) {
	r := hertzHandler(New())
	w := ut.PerformRequest(r, http.MethodGet, "/", nil, ut.Header{
		Key:   headerXRequestID,
		Value: testXRequestID,
	})

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.DeepEqual(t, testXRequestID, string(w.Header().Peek(headerXRequestID)))
}

func TestRequestIDWithCustomGenerator(t *testing.T) {
	r := hertzHandler(New(
		WithGenerator(func() string {
			return testXRequestID
		}),
	))

	w := ut.PerformRequest(r, http.MethodGet, "/", nil)

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.DeepEqual(t, testXRequestID, string(w.Header().Peek(headerXRequestID)))
}

func TestRequestIDWithCustomHeaderKey(t *testing.T) {
	r := hertzHandler(New(
		WithCustomHeaderStrKey(customHeaderKey),
	))

	w := ut.PerformRequest(r, http.MethodGet, "/", nil, ut.Header{
		Key:   customHeaderKey,
		Value: testXRequestID,
	})

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.DeepEqual(t, testXRequestID, string(w.Header().Peek(customHeaderKey)))
}

func TestRequestIDWithHandler(t *testing.T) {

	called := false

	r := hertzHandler(New(
		WithHandler(func(ctx context.Context, c *app.RequestContext, requestID string) {
			called = true
			assert.DeepEqual(t, testXRequestID, requestID)
		})))

	w := ut.PerformRequest(r, http.MethodGet, "/", nil, ut.Header{
		Key:   "X-Request-ID",
		Value: testXRequestID,
	})

	assert.DeepEqual(t, http.StatusOK, w.Code)
	assert.True(t, called)
}

func TestGetRequestID(t *testing.T) {

	r := route.NewEngine(hzconfig.NewOptions([]hzconfig.Option{}))
	r.Use(New())
	r.GET("/", func(ctx context.Context, c *app.RequestContext) {
		assert.DeepEqual(t, testXRequestID, Get(c))
	})

	_ = ut.PerformRequest(r, http.MethodGet, "/", nil, ut.Header{
		Key:   "X-Request-ID",
		Value: testXRequestID,
	})
}
