package base

import (
	"net/http"
)

// Response struct.
type Response struct {
	httpResp *http.Response
	depth    uint32
}

// NewResponse create a Response struct.
func NewResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp: httpResp, depth: depth}
}

// HTTPResp return a Response struct.
func (resp *Response) HTTPResp() *http.Response {
	return resp.httpResp
}

// Depth return spider depth.
func (resp *Response) Depth() uint32 {
	return resp.depth
}

// Valid return Response valid .
// Data interface impl .
func (resp *Response) Valid() bool {
	return resp.httpResp != nil && resp.httpResp.Body != nil
}
