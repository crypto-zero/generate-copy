package structs

import "github.com/crypto-zero/generate-copy/testdata/basic/base"

type MessageType int32

type MessageCategory int32

type C MessageType

type GC struct {
	CC string
	DG
}

type DG struct {
	W string
}

type model2 struct {
	Model2ID uint64
}

type Model3 struct {
	*Model4
}

type Model4 struct {
	Model4ID uint
}

type Message struct {
	base.Model
	Model3
	Model2      model2
	UserID      uint64
	MessageType MessageType
	TemplateID  int32
	Language    string
	Category    MessageCategory
	Title       string
	Content     string
	Redirect    string
	HasRead     bool
	CC          C
	GC
	MC         MessageCategory
	IntToUint  int
	FloatToInt int
}

type PMessage struct {
	ID         uint64
	CreatedAt  int64
	Category   int32
	Title      string
	Content    string
	Redirect   string
	HasRead    bool
	MC         uint32
	IntToUint  uint
	FloatToInt float64
	Model4ID   int
}
