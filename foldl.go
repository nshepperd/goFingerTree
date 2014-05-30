package fingerTree23

func Foldl(f FoldFunc, initial interface{}, s []Data, length int) interface{} {
	if length > 0 {
		return f(Foldl(f, initial, s[:length-1], length-1), s[length-1])
	}
	return initial
}
