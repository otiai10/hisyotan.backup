package conf

import "fmt"

type testConf struct{}

// Name ...
func (c *testConf) Name() string {
	return "hisyotan"
}

// RedisNS ...
func (c *testConf) RedisNS(key string) string {
	return fmt.Sprintf("%s.%s", "hisyotan.test", key)
}

// RedisHost ...
func (c *testConf) RedisHost() string {
	return "localhost"
}

// RedisPort ...
func (c *testConf) RedisPort() string {
	return "6379"
}
