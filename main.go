package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ohayao/gopic/clipboard"
)

func main() {
	example()
}

func example() {
	clip := clipboard.NewClipboard()

	// 文本
	clip.ToPaste([]byte("Hello, ohayao gopic!"), clipboard.Text)
	result, contentType, err := clip.GetCopy()
	if err == nil {
		fmt.Printf("String[%d]:\t%s\n", contentType, result)
	}

	time.Sleep(time.Second * 15)
	// 图片
	// 使用截图等截一张图 然后去获取它
	if img, contentType, err := clip.GetCopy(); err == nil {
		if contentType == clipboard.Image {
			fmt.Printf("Image[%d] Path:\t%s\n", contentType, img)
			// ******使用完图片后请删除，避免占用空间*****
			os.Remove(img)
		} else {
			fmt.Printf("No picture!\t Expected:%d Actual:%d\n", clipboard.Image, contentType)
		}
	}
}
