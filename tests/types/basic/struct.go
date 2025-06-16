//go:generate go run github.com/mchenriques22/pergolator github.com/mchenriques22/pergolator/tests/types/basic.Struct
package basic

type Struct struct {
	BasicString  string
	BasicInt     int
	BasicInt8    int8
	BasicInt16   int16
	BasicInt32   int32
	BasicInt64   int64
	BasicUint    uint
	BasicUint8   uint8
	BasicUint16  uint16
	BasicUint32  uint32
	BasicUint64  uint64
	BasicFloat32 float32
	BasicFloat64 float64
	BasicBool    bool
}
