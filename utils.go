package gno

func MapString(items []string, fn func(string) string) []string {
	ret := []string{}
	for _, item := range items {
		ret = append(ret, fn(item))
	}
	return ret
}

func MapKeysStringToAny(dict map[string]interface{}) []string {
	ret := []string{}
	for key, _ := range dict {
		ret = append(ret, key)
	}
	return ret
}

func MapValuesStringToAny(dict map[string]interface{}) []interface{} {
	ret := []interface{}{}
	for _, value := range dict {
		ret = append(ret, value)
	}
	return ret
}

func Fill(length int, value string) []string {
	var ret []string
	for i := 0; i < length; i++ {
		ret = append(ret, value)
	}
	return ret
}
