package downloader

import "github.com/chapin/spider/base"

// PageDownloader interface.
type PageDownloader interface {
	Id() uint32
	Download(req base.Request) (*base.Response, error)
}

// IDGenerator interface.
type IDGenerator interface {
	GetUint32() uint32
}

// PageDownloaderPool interface.
type PageDownloaderPool interface {
	Take() (PageDownloader, error)  // take a PageDownloader from pool.
	Return(dl PageDownloader) error // return PageDownloader to pool.
	Total() uint32                  // PageDownloaderPool capcity.
	Used() uint32                   // PageDownloaderPool used.
}
