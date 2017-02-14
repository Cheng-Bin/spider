package base

import (
	"net/http"
)

// Request .
type Request struct {
	httpReq *http.Request // HTTP请求的指针值
	depth   uint32        // 请求的深度
}

// NewRequest create a http request struct.
func NewRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq: httpReq, depth: depth}
}

// HTTPReq return http request.
func (req *Request) HTTPReq() *http.Request {
	return req.httpReq
}

// Depth return depth.
func (req *Request) Depth() uint32 {
	return req.depth
}

// Valid return http request is valid.
// Data interface impl.
func (req *Request) Valid() bool {
	return req.httpReq != nil && req.httpReq.URL != nil
}
