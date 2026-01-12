package cache

import "fmt"

type RedisKey struct{}

func (r RedisKey) UserKey(id int) string {
	return fmt.Sprintf("user:%d", id)
}
