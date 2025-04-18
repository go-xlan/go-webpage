package gintestpage_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/go-xlan/go-webpage/gintestpage"
	"github.com/go-xlan/go-webpage/internal/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var caseService *gintestpage.Service

func TestMain(m *testing.M) {
	service := gintestpage.NewService()
	defer service.Close()

	caseService = service
	m.Run()
}

func TestService_SetPage(t *testing.T) {
	path := uuid.New().String()
	link := caseService.SetPage(path, []byte(utils.NewPage("标题1", "内容1")))
	t.Log("path:", path)
	t.Log("link:", link)

	response, err := resty.New().R().Get(link)
	require.NoError(t, err)
	page := response.Body()
	t.Log("page:", string(page))
}

func TestService_GetLink(t *testing.T) {
	path := uuid.New().String()
	caseService.SetPage(path, []byte(utils.NewPage("标题2", "内容2")))
	t.Log("path:", path)
	link := caseService.GetLink(path)
	t.Log("link:", link)

	response, err := resty.New().R().Get(link)
	require.NoError(t, err)
	page := response.Body()
	t.Log("page:", string(page))
}
