# Struct field copy generator

Installation:

```shell
go install github.com/crypto-zero/generate-copy@main
```

How to use: 

Declare generate type variable, for example:

```go
//go:generate generate-copy
var GenerateCopyStructs = [][]any{
	{structs.Message{}, structs.PMessage{}, "MessageToProtoMessage"},
	{structs.PMessage{}, structs.Message{}, "ProtoMessageToMessage", true},
}
```

Then run go generate command to generate code.

```shell
go generate .
```