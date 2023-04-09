package Gee_RPC_Study

import "github.com/djz03/Gee-RPC-Study/codec"

const MagicNumber = 0x3be5fc

type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}
