package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

// InputDetails 用来存储用户通过命令行标签输入的详细信息
type InputDetails struct {
	Link     string // 用户指定的链接
	Username string // 用户名（可选）
	Password string // 密码（可选）
}

// HandleInput 用来处理用户输入
// 如果用户名和密码是空的，就认为不需要它们
func (i *InputDetails) HandleInput() {
	if i.Username == "" && i.Password == "" {
		fmt.Printf("开始处理链接：%s，无需用户名和密码\n", i.Link)
	} else {
		fmt.Printf("开始处理链接：%s，用户名：%s，密码：%s\n", i.Link, i.Username, i.Password)
	}
}

// 检查链接是否以http或https开头，并尝试解析以确保其格式正确
func validateLink(link string) (bool, error) {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		return false, fmt.Errorf("链接必须以'http://'或'https://'开头")
	}
	_, err := url.Parse(link)
	if err != nil {
		return false, fmt.Errorf("无法解析链接: %v", err)
	}
	return true, nil
}

func main() {
	// 定义标签，并设置默认值和帮助信息
	link := flag.String("link", "", "请输入链接")
	username := flag.String("username", "", "请输入用户名（如果不需要请留空）")
	password := flag.String("password", "", "请输入密码（如果不需要请留空）")

	// 解析命令行标签到相应变量
	flag.Parse()

	// 验证链接是否有效
	valid, err := validateLink(*link)
	if !valid {
		fmt.Println("错误:", err)
		return
	}

	// 使用命令行标签的值创建InputDetails结构
	details := InputDetails{
		Link:     *link,
		Username: *username,
		Password: *password,
	}

	// 处理输入
	details.HandleInput()
}
