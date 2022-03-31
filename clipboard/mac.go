package clipboard

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

type i_mac struct{}

func (that *i_mac) GetCopy() (ret string, ct ContentType, err error) {
	ct = that.get_copy_content_type()
	if ct == Text {
		ret, err = that.get_copy_content_string()
	} else if ct == Image {
		ret, err = that.get_copy_content_image()
	}
	return
}

func (that *i_mac) ToPaste(content []byte, contentType ContentType) error {
	if contentType != Text {
		return fmt.Errorf("Unsupport type %d", contentType)
	}
	cmd := exec.Command("pbcopy")
	if in, err := cmd.StdinPipe(); err != nil {
		return err
	} else {
		if err := cmd.Start(); err != nil {
			return err
		}
		if _, err := in.Write(content); err != nil {
			return err
		}
		if err := in.Close(); err != nil {
			return err
		}
	}
	return cmd.Wait()
}

func (that *i_mac) get_copy_content_type() ContentType {
	if res, err := exec.Command("osascript", "-e", "clipboard info").Output(); err != nil {
		fmt.Printf("get_copy_content_type error %+v\n", err)
		return Unknown
	} else {
		if strings.Index(string(res), "«class PNGf»") == 0 {
			return Image
		}
		return Text
	}
}
func (that *i_mac) get_copy_content_string() (string, error) {
	res, err := exec.Command("pbpaste").Output()
	return string(res), err
}
func (that *i_mac) get_copy_content_image() (string, error) {
	tempFile, err := ioutil.TempFile("", "*.png")
	if err != nil {
		return "", err
	}
	tempFile.Close()
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("write (the clipboard as «class PNGf») to (open for access \"%s\" with write permission)", tempFile.Name()))
	if _, err = cmd.CombinedOutput(); err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}
