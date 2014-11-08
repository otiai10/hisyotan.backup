package conf

import "fmt"

type prodConf struct{}

// Name ...
func (c *prodConf) Name() string {
	return "hisyotan"
}

// RedisNS ...
func (c *prodConf) RedisNS(key string) string {
	return fmt.Sprintf("%s.%s", c.Name(), key)
}

// RedisHost ...
func (c *prodConf) RedisHost() string {
	return "localhost"
}

// RedisPort ...
func (c *prodConf) RedisPort() string {
	return "6379"
}
