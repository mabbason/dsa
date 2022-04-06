package main

func main() {
	strSlc := []string{"a", "b", "c", "d", "c", "e", "f"}
	println(findDupStr(strSlc))
}

func findDupStr(strSlc []string) string {
	strHash := make(map[string]bool)
	for _, str := range strSlc {
		if !strHash[str] {
			strHash[str] = true
		} else {
			return str
		}
	}
	return ""
}
