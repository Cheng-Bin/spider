package analyzer

import (
	"net/http"

	"github.com/chapin/spider/base"
)

// ParseResponse struct .
type ParseResponse func(httpResp *http.Response, respDepth uint32) ([]base.Data, []error)

// Analyzer interface .
type Analyzer interface {
	Id() uint32
	Analyze(respParsers []ParseResponse, resp base.Response) ([]base.Data, []error)
}

// AnalyzerPool struct.
type AnalyzerPool interface {
	Take() (Analyzer, error)        // Take an Analyzer from AnalyzerPool.
	Return(analyzer Analyzer) error // Return Analyzer to pool.
	Total() uint32                  // Return Analyzer cacity.
	Used() uint32                   // Return Analyzer used.
}
