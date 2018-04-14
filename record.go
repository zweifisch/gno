package gno

type Record map[string]interface{}

func (record Record) Unzip() ([]string, []interface{}) {
	keys := []string{}
	values := []interface{}{}
	for key, value := range record {
		keys = append(keys, key)
		values = append(values, value)
	}
	return keys, values
}

func (record Record) Keys() []string {
	keys := []string{}
	for key, _ := range record {
		keys = append(keys, key)
	}
	return keys
}
