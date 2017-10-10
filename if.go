package IF

const (
	_uint8      = uint8(0)
	_uint16     = uint16(0)
	_uint32     = uint32(0)
	_uint64     = uint64(0)
	_int8       = int8(0)
	_int16      = int16(0)
	_int32      = int32(0)
	_int64      = int64(0)
	_float32    = float32(0)
	_float64    = float64(0)
	_complex64  = complex64(0)
	_complex128 = complex128(0)
	_byte       = byte(0)
	_rune       = rune(0)
	_uint       = uint(0)
	_int        = int(0)
)

// Else returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is interface{}.
func Else(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// ElseSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []interface{}
func ElseSlice(condition bool, a []interface{}, b []interface{}) []interface{} {
	if condition {
		return a
	}
	return b
}

// Nil returns second argument if object is nil
// otherwise returns first argument.
// This function's return value's type is interface{}.
func Nil(object interface{}, a interface{}) interface{} {
	if isNil(object) {
		return a
	}
	return object
}

// NotNil returns second argument unless object is nil
// otherwise returns first argument.
// This function's return value's type is interface{}.
func NotNil(object interface{}, a interface{}) interface{} {
	if !isNil(object) {
		return a
	}
	return object
}

// ElseUint8 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uint8.
func ElseUint8(condition bool, a uint8, b uint8) uint8 {
	if condition {
		return a
	}
	return b
}

// ElseUint8Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uint8.
func ElseUint8Slice(condition bool, a []uint8, b []uint8) []uint8 {
	if condition {
		return a
	}
	return b
}

// Uint8IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is uint8.
func Uint8IsZero(n uint8, a uint8) uint8 {
	if n == _uint8 {
		return a
	}
	return n
}

// ElseUint16 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uint16.
func ElseUint16(condition bool, a uint16, b uint16) uint16 {
	if condition {
		return a
	}
	return b
}

// ElseUint16Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uint16.
func ElseUint16Slice(condition bool, a []uint16, b []uint16) []uint16 {
	if condition {
		return a
	}
	return b
}

// Uint16IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is uint16.
func Uint16IsZero(n uint16, a uint16) uint16 {
	if n == _uint16 {
		return a
	}
	return n
}

// ElseUint32 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uint32.
func ElseUint32(condition bool, a uint32, b uint32) uint32 {
	if condition {
		return a
	}
	return b
}

// ElseUint32Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uint32.
func ElseUint32Slice(condition bool, a []uint32, b []uint32) []uint32 {
	if condition {
		return a
	}
	return b
}

// Uint32IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is uint32.
func Uint32IsZero(n uint32, a uint32) uint32 {
	if n == _uint32 {
		return a
	}
	return n
}

// ElseUint64 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uint64.
func ElseUint64(condition bool, a uint64, b uint64) uint64 {
	if condition {
		return a
	}
	return b
}

// ElseUint64Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uint64.
func ElseUint64Slice(condition bool, a []uint64, b []uint64) []uint64 {
	if condition {
		return a
	}
	return b
}

// Uint64IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is uint64.
func Uint64IsZero(n uint64, a uint64) uint64 {
	if n == _uint64 {
		return a
	}
	return n
}

// ElseInt8 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is int8.
func ElseInt8(condition bool, a int8, b int8) int8 {
	if condition {
		return a
	}
	return b
}

// ElseInt8Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []int8.
func ElseInt8Slice(condition bool, a []int8, b []int8) []int8 {
	if condition {
		return a
	}
	return b
}

// Int8IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is int8.
func Int8IsZero(n int8, a int8) int8 {
	if n == _int8 {
		return a
	}
	return n
}

// ElseInt16 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is int16.
func ElseInt16(condition bool, a int16, b int16) int16 {
	if condition {
		return a
	}
	return b
}

// ElseInt16Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []int16.
func ElseInt16Slice(condition bool, a []int16, b []int16) []int16 {
	if condition {
		return a
	}
	return b
}

// Int16IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is int16.
func Int16IsZero(n int16, a int16) int16 {
	if n == _int16 {
		return a
	}
	return n
}

// ElseInt32 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is int32.
func ElseInt32(condition bool, a int32, b int32) int32 {
	if condition {
		return a
	}
	return b
}

// ElseInt32Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []int32.
func ElseInt32Slice(condition bool, a []int32, b []int32) []int32 {
	if condition {
		return a
	}
	return b
}

// Int32IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is int32.
func Int32IsZero(n int32, a int32) int32 {
	if n == _int32 {
		return a
	}
	return n
}

// ElseInt64 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is int64.
func ElseInt64(condition bool, a int64, b int64) int64 {
	if condition {
		return a
	}
	return b
}

// ElseInt64Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []int64.
func ElseInt64Slice(condition bool, a []int64, b []int64) []int64 {
	if condition {
		return a
	}
	return b
}

// Int64IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is int64.
func Int64IsZero(n int64, a int64) int64 {
	if n == _int64 {
		return a
	}
	return n
}

// ElseFloat32 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is float32.
func ElseFloat32(condition bool, a float32, b float32) float32 {
	if condition {
		return a
	}
	return b
}

// ElseFloat32Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []float32.
func ElseFloat32Slice(condition bool, a []float32, b []float32) []float32 {
	if condition {
		return a
	}
	return b
}

// Float32IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is float32.
func Float32IsZero(n float32, a float32) float32 {
	if n == _float32 {
		return a
	}
	return n
}

// ElseFloat64 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is float64.
func ElseFloat64(condition bool, a float64, b float64) float64 {
	if condition {
		return a
	}
	return b
}

// ElseFloat64Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []float64.
func ElseFloat64Slice(condition bool, a []float64, b []float64) []float64 {
	if condition {
		return a
	}
	return b
}

// Float64IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is float64.
func Float64IsZero(n float64, a float64) float64 {
	if n == _float64 {
		return a
	}
	return n
}

// ElseComplex64 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is complex64.
func ElseComplex64(condition bool, a complex64, b complex64) complex64 {
	if condition {
		return a
	}
	return b
}

// ElseComplex64Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []complex64.
func ElseComplex64Slice(condition bool, a []complex64, b []complex64) []complex64 {
	if condition {
		return a
	}
	return b
}

// Complex64IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is complex64.
func Complex64IsZero(n complex64, a complex64) complex64 {
	if n == _complex64 {
		return a
	}
	return n
}

// ElseComplex128 returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is complex128.
func ElseComplex128(condition bool, a complex128, b complex128) complex128 {
	if condition {
		return a
	}
	return b
}

// ElseComplex128Slice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []complex128.
func ElseComplex128Slice(condition bool, a []complex128, b []complex128) []complex128 {
	if condition {
		return a
	}
	return b
}

// Complex128IsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is complex128.
func Complex128IsZero(n complex128, a complex128) complex128 {
	if n == _complex128 {
		return a
	}
	return n
}

// ElseByte returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is byte.
func ElseByte(condition bool, a byte, b byte) byte {
	if condition {
		return a
	}
	return b
}

// ElseByteSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []byte.
func ElseByteSlice(condition bool, a []byte, b []byte) []byte {
	if condition {
		return a
	}
	return b
}

// ByteIsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is byte.
func ByteIsZero(n byte, a byte) byte {
	if n == _byte {
		return a
	}
	return n
}

// ElseRune returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is rune.
func ElseRune(condition bool, a rune, b rune) rune {
	if condition {
		return a
	}
	return b
}

// ElseRuneSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []rune.
func ElseRuneSlice(condition bool, a []rune, b []rune) []rune {
	if condition {
		return a
	}
	return b
}

// RuneIsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is rune.
func RuneIsZero(n rune, a rune) rune {
	if n == _rune {
		return a
	}
	return n
}

// ElseUint returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uint.
func ElseUint(condition bool, a uint, b uint) uint {
	if condition {
		return a
	}
	return b
}

// ElseUintSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uint.
func ElseUintSlice(condition bool, a []uint, b []uint) []uint {
	if condition {
		return a
	}
	return b
}

// UintIsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is uint.
func UintIsZero(n uint, a uint) uint {
	if n == _uint {
		return a
	}
	return n
}

// ElseInt returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is int.
func ElseInt(condition bool, a int, b int) int {
	if condition {
		return a
	}
	return b
}

// ElseIntSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []int.
func ElseIntSlice(condition bool, a []int, b []int) []int {
	if condition {
		return a
	}
	return b
}

// IntIsZero returns second argument if first argument is 0
// otherwise returns first argument.
// This function's return value's type is int.
func IntIsZero(n int, a int) int {
	if n == _int {
		return a
	}
	return n
}

// ElseUintptr returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is uintptr.
func ElseUintptr(condition bool, a uintptr, b uintptr) uintptr {
	if condition {
		return a
	}
	return b
}

// ElseUintptrSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []uintptr.
func ElseUintptrSlice(condition bool, a []uintptr, b []uintptr) []uintptr {
	if condition {
		return a
	}
	return b
}

// ElseString returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is string.
func ElseString(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

// ElseStringSlice returns second argument if condition is true
// otherwise returns third argument.
// This function's return value's type is []string.
func ElseStringSlice(condition bool, a []string, b []string) []string {
	if condition {
		return a
	}
	return b
}

// StringIsEmpty returns first argument if first argument is ""
// otherwise returns second argument.
// This function's return value's type is string.
func StringIsEmpty(str string, a string) string {
	if str == "" {
		return str
	}
	return str
}

// isNil is stolen from https://github.com/stretchr/testify/blob/master/assert/assertions.go
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if reflect.Chan <= kind && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}
