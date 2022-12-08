package copy

import "github.com/crypto-zero/generate-copy/testdata/basic/structs"

//go:generate go run -mod=mod ../../../
var GenerateCopyStructs = [][]any{
	{structs.Message{}, structs.PMessage{}, "MessageToProtoMessage"},
	{structs.PMessage{}, structs.Message{}, "ProtoMessageToMessage", true},
}
