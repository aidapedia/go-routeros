package util

func ToQuery(path string, attr map[string]string) []string {
	result := []string{
		path,
	}
	for key, value := range attr {
		result = append(result, "="+key+"="+value)
	}
	return result
}
