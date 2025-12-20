package pkg

import "fmt"

func RedisKey(domain, entity string, identifier any) string {
	return fmt.Sprintf("%s:%s:%v", domain, entity, identifier)
}
