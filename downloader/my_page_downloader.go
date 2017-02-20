package downloader

import (
	"net/http"

	"github.com/chapin/spider/base"
	mdw "github.com/chapin/spider/middleware"
)

var downloaderIDGenertor mdw.IDGenerator = mdw.NewIDGenertor()

func genDownloaderID() uint32 {
	return downloaderIDGenertor.GetUint32()
}

type myPageDownloader struct {
	httpClient http.Client
	id         uint32
}

func NewPageDownloader(client *http.Client) PageDownloader {
	id := genDownloaderID()
	if client == nil {
		client = &http.Client{}
	}
	return &myPageDownloader{
		id:         id,
		httpClient: *client,
	}
}

func (dl *myPageDownloader) Id() uint32 {
	return dl.id
}

func (dl *myPageDownloader) Download(req base.Request) (*base.Response, error) {
	httpReq := req.HTTPReq()
	httpResp, err := dl.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return base.NewResponse(httpResp, req.Depth()), nil
}
