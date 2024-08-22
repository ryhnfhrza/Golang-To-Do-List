package helper

func GetDefaultIfEmpty(value, defaultValue string) string {
	if value != "" {
		return value
	}
	return defaultValue
}