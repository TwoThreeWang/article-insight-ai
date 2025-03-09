# ArticleInsight - AI文章总结工具

一款基于人工智能技术的文章总结工具，能够快速提取文章核心内容，生成结构化摘要，帮助用户提高阅读效率。

## 功能特点

- 支持URL链接和直接文本输入两种方式
- 智能提取文章核心内容
- 生成结构化的Markdown格式摘要
- 支持预览和Markdown源码查看
- 提供历史记录功能
- 简洁美观的用户界面

## 技术栈

### 后端
- Go语言
- Google Gemini AI API

### 前端
- HTML5/CSS3/JavaScript
- Tailwind CSS
- Font Awesome 图标库
- Marked.js (Markdown解析)

## 快速开始

### 环境要求
- Go 1.16+
- Google Gemini API Key

### 安装步骤

1. 克隆项目
```bash
git clone https://github.com/yourusername/ArticleInsight.git
cd ArticleInsight
```

2. 安装依赖
```bash
go mod download
```

3. 配置环境
- 复制 `config/config_exp.json` 为 `config/config.json`
- 在 `config.json` 中填入你的 Gemini API Key

4. 运行项目
```bash
go run main.go
```

5. 访问应用
打开浏览器访问 `http://localhost:8080`

## 使用说明

1. 在首页输入框中粘贴文章URL或直接输入文章内容
2. 点击"开始总结"按钮
3. 等待AI分析完成
4. 查看生成的结构化摘要
5. 可以切换预览和Markdown视图
6. 点击"复制MD全文