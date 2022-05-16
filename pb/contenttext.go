package pb

import "encoding/json"

type TextContent struct {
	Text string
}

func (t *TextContent) Bytes() []byte {
	buf, _ := json.Marshal(t)
	return buf
}

func (t *TextContent) String() string {
	return t.Text
}

func NewTextContent(content []byte) Content {
	text := &TextContent{}
	_ = json.Unmarshal(content, text)
	return text
}
