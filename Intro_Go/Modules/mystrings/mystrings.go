package mystrings

/*
	Reverse reverses a string left to right.
	Notice that we need to capitalize the
	first letter of the function.
	If we do not then we won't be able access
	this function outside of the 'mystrings' package.
*/

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
