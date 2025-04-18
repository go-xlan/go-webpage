package gintestpage

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/yyle88/must"
	"github.com/yyle88/must/muststrings"
	"github.com/yyle88/rese"
	"github.com/yyle88/syncmap"
)

type Service struct {
	srvTest *httptest.Server
	pageMap *syncmap.Map[string, []byte]
}

func NewService() *Service {
	pageMap := syncmap.New[string, []byte]()

	engine := gin.New()
	engine.GET(":path", func(c *gin.Context) {
		path := c.Param("path")

		page, ok := pageMap.Load(path)
		if !ok {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Data(200, "text/html; charset=utf-8", page)
	})

	return &Service{
		srvTest: httptest.NewServer(engine), //contains start() function
		pageMap: pageMap,
	}
}

func (service *Service) Close() {
	service.srvTest.Close()
}

func (service *Service) SetPage(path string, page []byte) string {
	muststrings.NotContains(must.Nice(path), "/")
	service.pageMap.Store(path, page)
	return service.GetLink(path)
}

func (service *Service) GetLink(path string) string {
	muststrings.NotContains(must.Nice(path), "/")
	return rese.C1(url.JoinPath(service.srvTest.URL, path))
}
