package main

import (
	"ArticleInsight/config"
	"ArticleInsight/models"
	"ArticleInsight/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type SummaryRequest struct {
	URL  string `json:"url" binding:"required"`
	TYPE string `json:"type" binding:"required"`
}

type SummaryResponse struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Error   string `json:"error,omitempty"`
}

func main() {
	// 设置日志输出
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fmt.Println("AI 文章总结工具启动:", time.Now().Format("2006-01-02 15:04:05"))
	// 初始化配置
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 注入配置到中间件
	r.Use(func(c *gin.Context) {
		c.Set("cfg", cfg)
		c.Next()
	})

	// 设置模板目录
	r.LoadHTMLGlob("templates/*")

	// 设置静态文件目录
	r.Static("/static", "static")

	// 主页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "AI文章总结工具",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// 文章总结接口
		api.POST("/summarize", handleSummarize)
	}

	// 创建必要的目录
	createDirs()

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// 处理文章总结请求
func handleSummarize(c *gin.Context) {
	// 从上下文获取配置
	cfg, exists := c.Get("cfg")
	if !exists {
		c.JSON(http.StatusBadRequest, SummaryResponse{
			Error: "配置未找到",
		})
		return
	}

	// 初始化服务
	aiService, err := services.NewAIService(cfg.(*config.Config))
	if err != nil {
		c.JSON(http.StatusBadRequest, SummaryResponse{
			Error: "AI 服务初始化报错：" + err.Error(),
		})
		return
	}

	// 获取请求参数
	var req SummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, SummaryResponse{
			Error: "Invalid request format",
		})
		return
	}

	// 获取页面内容
	var content string
	if req.TYPE == "url" {
		fetchService := services.NewFetchService(cfg.(*config.Config))
		content, err = fetchService.FetchStory(req.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, SummaryResponse{
				Error: "获取文章内容失败：" + err.Error(),
			})
			return
		}
	} else {
		content = fmt.Sprintf("\n<article>\n%s\n</article>\n", req.URL)
	}

	story := &models.Story{Content: content}

	// AI 文章总结
	err = aiService.GenerateSummary(story)
	if err != nil {
		c.JSON(http.StatusBadRequest, SummaryResponse{
			Error: "总结文章内容失败：" + err.Error(),
		})
		return
	}
	// 这里先返回模拟数据
	c.JSON(http.StatusOK, SummaryResponse{
		Summary: story.Summary,
	})
}

// 创建必要的目录
func createDirs() {
	dirs := []string{"templates", "static"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Printf("Failed to create directory %s: %v", dir, err)
		}
	}
}
