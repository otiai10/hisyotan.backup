package conf

var instance ConfSet

// Name ...
func Name(args ...bool) string {
	var prefix string
	if len(args) > 0 && args[0] {
		prefix = "@"
	}
	return prefix + instance.Name()
}

// RedisNS ...
func RedisNS(key string) string {
	return instance.RedisNS(key)
}

// RedisHost ...
func RedisHost() string {
	return instance.RedisHost()
}

// RedisPort ...
func RedisPort() string {
	return instance.RedisPort()
}

// ConfSet ...
type ConfSet interface {
	Name() string
	RedisNS(string) string
	RedisHost() string
	RedisPort() string
}

func init() {
	instance = &prodConf{}
}

// SetEnv ...
func SetEnv(key string) {
	switch key {
	case "test":
		instance = &testConf{}
	}
}
