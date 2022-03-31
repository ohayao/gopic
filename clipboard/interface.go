package clipboard

import (
	"runtime"
	"strings"
)

type ContentType uint

const (
	Unknown ContentType = iota // 未知
	Text                       // 文本
	Image                      // 图片
)

type Clipboard interface {
	// 获取复制内容
	GetCopy() (string, ContentType, error)
	// 将内容写到粘贴板
	ToPaste([]byte, ContentType) error
}

func NewClipboard() Clipboard {
	var clipboard Clipboard
	platform := strings.ToLower(runtime.GOOS)
	if platform == "darwin" {
		clipboard = new(i_mac)
	} else if platform == "windows" {
		clipboard = new(i_win)
	} else if platform == "linux" {

	} else {

	}
	return clipboard
}
