package clipboard

import (
	"testing"
)

func TestCopyPaste(t *testing.T) {

	data := []struct {
		input_type ContentType
		input_text string
	}{
		{
			Text,
			"0861008610010",
		},
		{
			Text,
			"ABcd6789#%#*$W(🥲⚡︎✓✕✦◆❢",
		},
		{
			Text,
			"中国加油！★★★★★❤︎❤︎❤︎",
		},
	}
	clip := NewClipboard()
	for _, v := range data {
		if err := clip.ToPaste([]byte(v.input_text), v.input_type); err != nil {
			t.Fatal(err)
			return
		} else {
			if res, ct, err := clip.GetCopy(); err != nil {
				t.Fatal(err)
				return
			} else {
				if res == v.input_text && ct == v.input_type {
					t.Logf("verify OK %d %s\n", ct, res)
				}
			}
		}
	}
}
