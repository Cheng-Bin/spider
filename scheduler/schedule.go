package scheduler

import (
	"net/http"

	"github.com/chapin/spider/analyzer"
	"github.com/chapin/spider/itempipeline"
)

// GenHTTPClient .
type GenHTTPClient func() *http.Client

// SchedSummary interface .
type SchedSummary interface {
	String() string
	Detail() string
	Same(other SchedSummary) bool
}

// Scheduler interface .
type Scheduler interface {
	Start(channelLen uint,
		poolSize uint32,
		crawlDepth uint32,
		httpClientGenerator GenHTTPClient,
		respParsers []analyzer.ParseResponse,
		itemProcessors []itempipeline.ProcessItem,
		firstHTTPReq *http.Request) (err error)

	// Stop method.
	Stop() bool

	// Running state.
	Running() bool

	// ErrorChan
	ErrorChan() <-chan error

	// Idle
	Idle() bool

	// Summary
	Summary(prefix string) SchedSummary
}
