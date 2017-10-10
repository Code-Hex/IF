package IF

func Else(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func ElseSlice(condition bool, a []interface{}, b []interface{}) []interface{} {
	if condition {
		return a
	}
	return b
}

func Nil(object interface{}, a interface{}) interface{} {
	if isNil(object) {
		return a
	}
	return object
}

func NotNil(object interface{}, a interface{}) interface{} {
	if !isNil(object) {
		return a
	}
	return object
}

func ElseUint8(condition bool, a uint8, b uint8) uint8 {
	if condition {
		return a
	}
	return b
}

func ElseUint8Slice(condition bool, a []uint8, b []uint8) []uint8 {
	if condition {
		return a
	}
	return b
}

func ZeroUint8(n uint8, a uint8) uint8 {
	if n == uint8(0) {
		return a
	}
	return n
}

func ElseUint16(condition bool, a uint16, b uint16) uint16 {
	if condition {
		return a
	}
	return b
}

func ElseUint16Slice(condition bool, a []uint16, b []uint16) []uint16 {
	if condition {
		return a
	}
	return b
}

func ZeroUint16(n uint16, a uint16) uint16 {
	if n == uint16(0) {
		return a
	}
	return n
}

func ElseUint32(condition bool, a uint32, b uint32) uint32 {
	if condition {
		return a
	}
	return b
}

func ElseUint32Slice(condition bool, a []uint32, b []uint32) []uint32 {
	if condition {
		return a
	}
	return b
}

func ZeroUint32(n uint32, a uint32) uint32 {
	if n == uint32(0) {
		return a
	}
	return n
}

func ElseUint64(condition bool, a uint64, b uint64) uint64 {
	if condition {
		return a
	}
	return b
}

func ElseUint64Slice(condition bool, a []uint64, b []uint64) []uint64 {
	if condition {
		return a
	}
	return b
}

func ZeroUint64(n uint64, a uint64) uint64 {
	if n == uint64(0) {
		return a
	}
	return n
}

func ElseInt8(condition bool, a int8, b int8) int8 {
	if condition {
		return a
	}
	return b
}

func ElseInt8Slice(condition bool, a []int8, b []int8) []int8 {
	if condition {
		return a
	}
	return b
}

func ZeroInt8(n int8, a int8) int8 {
	if n == int8(0) {
		return a
	}
	return n
}

func ElseInt16(condition bool, a int16, b int16) int16 {
	if condition {
		return a
	}
	return b
}

func ElseInt16Slice(condition bool, a []int16, b []int16) []int16 {
	if condition {
		return a
	}
	return b
}

func ZeroInt16(n int16, a int16) int16 {
	if n == int16(0) {
		return a
	}
	return n
}

func ElseInt32(condition bool, a int32, b int32) int32 {
	if condition {
		return a
	}
	return b
}

func ElseInt32Slice(condition bool, a []int32, b []int32) []int32 {
	if condition {
		return a
	}
	return b
}

func ZeroInt32(n int32, a int32) int32 {
	if n == int32(0) {
		return a
	}
	return n
}

func ElseInt64(condition bool, a int64, b int64) int64 {
	if condition {
		return a
	}
	return b
}

func ElseInt64Slice(condition bool, a []int64, b []int64) []int64 {
	if condition {
		return a
	}
	return b
}

func ZeroInt64(n int64, a int64) int64 {
	if n == int64(0) {
		return a
	}
	return n
}

func ElseFloat32(condition bool, a float32, b float32) float32 {
	if condition {
		return a
	}
	return b
}

func ElseFloat32Slice(condition bool, a []float32, b []float32) []float32 {
	if condition {
		return a
	}
	return b
}

func ZeroFloat32(n float32, a float32) float32 {
	if n == float32(0) {
		return a
	}
	return n
}

func ElseFloat64(condition bool, a float64, b float64) float64 {
	if condition {
		return a
	}
	return b
}

func ElseFloat64Slice(condition bool, a []float64, b []float64) []float64 {
	if condition {
		return a
	}
	return b
}

func ZeroFloat64(n float64, a float64) float64 {
	if n == float64(0) {
		return a
	}
	return n
}

func ElseComplex64(condition bool, a complex64, b complex64) complex64 {
	if condition {
		return a
	}
	return b
}

func ElseComplex64Slice(condition bool, a []complex64, b []complex64) []complex64 {
	if condition {
		return a
	}
	return b
}

func ZeroComplex64(n complex64, a complex64) complex64 {
	if n == complex64(0) {
		return a
	}
	return n
}

func ElseComplex128(condition bool, a complex128, b complex128) complex128 {
	if condition {
		return a
	}
	return b
}

func ElseComplex128Slice(condition bool, a []complex128, b []complex128) []complex128 {
	if condition {
		return a
	}
	return b
}

func ZeroComplex128(n complex128, a complex128) complex128 {
	if n == complex128(0) {
		return a
	}
	return n
}

func ElseByte(condition bool, a byte, b byte) byte {
	if condition {
		return a
	}
	return b
}

func ElseByteSlice(condition bool, a []byte, b []byte) []byte {
	if condition {
		return a
	}
	return b
}

func ZeroByte(n byte, a byte) byte {
	if n == byte(0) {
		return a
	}
	return n
}

func ElseRune(condition bool, a rune, b rune) rune {
	if condition {
		return a
	}
	return b
}

func ElseRuneSlice(condition bool, a []rune, b []rune) []rune {
	if condition {
		return a
	}
	return b
}

func ZeroRune(n rune, a rune) rune {
	if n == rune(0) {
		return a
	}
	return n
}

func ElseUint(condition bool, a uint, b uint) uint {
	if condition {
		return a
	}
	return b
}

func ElseUintSlice(condition bool, a []uint, b []uint) []uint {
	if condition {
		return a
	}
	return b
}

func ZeroUint(n uint, a uint) uint {
	if n == uint(0) {
		return a
	}
	return n
}

func ElseInt(condition bool, a int, b int) int {
	if condition {
		return a
	}
	return b
}

func ElseIntSlice(condition bool, a []int, b []int) []int {
	if condition {
		return a
	}
	return b
}

func ZeroInt(n int, a int) int {
	if n == int(0) {
		return a
	}
	return n
}

func ElseUintptr(condition bool, a uintptr, b uintptr) uintptr {
	if condition {
		return a
	}
	return b
}

func ElseUintptrSlice(condition bool, a []uintptr, b []uintptr) []uintptr {
	if condition {
		return a
	}
	return b
}

func ElseString(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func ElseStringSlice(condition bool, a []string, b []string) []string {
	if condition {
		return a
	}
	return b
}

func Empty(str string, a string) string {
	if str == "" {
		return str
	}
	return str
}

// isNil code stolen from https://github.com/stretchr/testify/blob/master/assert/assertions.go
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
