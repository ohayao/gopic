package clipboard

type i_win struct{}

func (that *i_win) GetCopy() (string, ContentType, error) {
	return "", 0, nil
}

func (that *i_win) ToPaste(content []byte, contentType ContentType) error {
	return nil
}
