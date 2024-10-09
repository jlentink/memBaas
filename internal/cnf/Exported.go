package cnf

// GetString is a wrapper for viper.GetString
func GetString(key string) string {
	return Get().GetString(key)
}

func GetBytes(key string) []byte {
	return []byte(GetString(key))
}

// GetInt is a wrapper for viper.GetInt
func GetInt(key string) int {
	return Get().GetInt(key)
}

func IsDebug() bool {
	return Get().GetString("server.mode") == "debug"
}

// GetBool is a wrapper for viper.GetBool
func GetBool(key string) bool {
	return Get().GetBool(key)
}

// GetStringSlice is a wrapper for viper.GetStringSlice
func GetStringSlice(key string) []string {
	return Get().GetStringSlice(key)
}
