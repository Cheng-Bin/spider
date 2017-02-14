package base

// ErrorType struct.
type ErrorType string

// const
const (
	// DownloaderError flag
	DownloaderError ErrorType = "Downloader Error"
	// AnalyzerError flag
	AnalyzerError ErrorType = "Analyzer Error"
	// ItemProcessorError flag
	ItemProcessorError ErrorType = "Item Processor Error"
)

// SpiderError base.
type SpiderError interface {
	Type() ErrorType // error type
	Error() string   // error message
}
