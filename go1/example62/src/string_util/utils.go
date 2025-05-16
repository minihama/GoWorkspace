package string_util

// Stringer 인터페이스는 문자열형을 처리하는 인터페이스입니다.
type Stringer interface {
	String() string
}

// Reverse 함수는 문자열을 반대로 뒤집습니다.
// 예를 들어 abcd는 dcba가 됩니다.
func Reverse(s string) string {
	out := ""
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}
	return out
}
