package testmap

func someFunction(key string) (value string) {
	test := map[string]string{
		"test": "value",
	}
	return test[key]
}
