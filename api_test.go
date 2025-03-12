package boluobaoLib

import (
	"encoding/json"
	"fmt"
	"github.com/BUnipendix/boluobaoLib/boluobaomodel"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoluobao(t *testing.T) {
	cookie := os.Getenv("BOLUOBAO_COOKIE")
	if cookie == "" {
		t.Skip("cookie is empty")
	}
	client := NewClient(
		WithCookie(cookie),
		WithDebug(),
	)

	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "GetAccount",
			test: func(t *testing.T) {
				data, err := client.API().GetUserInfo()
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "GetAccountIp",
			test: func(t *testing.T) {
				data, err := client.API().GetCurreyIp()
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "GeContentTsukkomis",
			test: func(t *testing.T) {
				data, err := client.API().GeContentTsukkomis(1, 6208426, 1)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "GetMoney",
			test: func(t *testing.T) {
				data, err := client.API().GetUserMoney()
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "AuthorInfo",
			test: func(t *testing.T) {
				data, err := client.API().GetOtherUserInfo("184263")
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "BookInfo",
			test: func(t *testing.T) {
				data, err := client.API().GetBookInfo(165216)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "GetBookshelfList",
			test: func(t *testing.T) {
				data, err := client.API().GetBookShelfInfo()
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "Catalogue",
			test: func(t *testing.T) {
				data, err := client.API().GetCatalogue(165216)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "Content",
			test: func(t *testing.T) {
				data, err := client.API().GetChapterContent(6208426)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "Rank",
			test: func(t *testing.T) {
				data, err := client.API().GetRankWeekArray(boluobaomodel.RankQuery.SaleRank, 1)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "Search",
			test: func(t *testing.T) {
				data, err := client.API().GetSearch("诡秘之主", 1)
				assert.NoError(t, err)
				assert.Equal(t, http.StatusOK, data.GetCode())
				printJSON(data)
			},
		},
		{
			name: "DownloadImage",
			test: func(t *testing.T) {
				imageURL := "https://img.boluobaomo.com/upload/images/novel/2023/07/13/165216/64094721_202307131420129457.jpg"
				err := downloadImage(imageURL, "image.jpg")
				assert.NoError(t, err)
				fmt.Println("图片下载成功！")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func printJSON(data interface{}) {
	marshal, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("Failed to marshal data: %v\n", err)
		return
	}
	fmt.Println(string(marshal))
}

func downloadImage(url, filePath string) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download image: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Printf("failed to close body: %v\n", err)
		}
	}(res.Body)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Printf("failed to close file: %v\n", err)
		}
	}(file)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	return nil
}
