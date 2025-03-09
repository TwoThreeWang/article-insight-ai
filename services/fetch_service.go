package services

import (
	"ArticleInsight/config"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FetchService struct {
	config *config.Config
	client *http.Client
}

func NewFetchService(cfg *config.Config) *FetchService {
	return &FetchService{
		config: cfg,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// FetchStory 获取单个文章的详细信息
func (s *FetchService) FetchStory(url string) (string, error) {
	// 设置请求头
	headers := make(http.Header)
	//headers.Set("X-Retain-Images", "none")

	// 获取文章内容
	req, err := http.NewRequest("GET", "https://r.jina.ai/"+url, nil)
	if err != nil {
		return "", err
	}
	req.Header = headers

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s %s", resp.Status, url)
	}

	articleBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 构建返回内容
	content := fmt.Sprintf("\n<article>\n%s\n</article>\n", string(articleBody))

	return content, nil
}
