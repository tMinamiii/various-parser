package kwtest

import "strings"

func clean(path string) string {
	symbols := strings.Split(path, "/")
	var result []string
	for _, v := range symbols {
		if v == ".." {
			if result[len(result)-1] != ".." {
				result = result[:len(result)-1]
			} else {
				result = append(result, v)
			}
		} else {
			result = append(result, v)
		}

	}
	if len(result) == 0 {
		return ""
	}
	return strings.Join(result, "/")
}
