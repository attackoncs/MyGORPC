package codec

import "io"

// 头
type Header struct {
	ServiceMethod string //服务吗和方法名，形如"Service.Method"
	Seq           uint64 //被客户端选择的序列号
	Error         string //错误信息
}

// 实现不同的Codec实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

type Type string

// 两种Codec，实际只用Gob,但实现类似，只需Gob缓存Json即可
const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
