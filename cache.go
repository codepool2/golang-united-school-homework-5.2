package cache

import (
	"time"
)

type Cache struct {
	data map[string]Value
}

type Value struct {
	data      string
	hasExpiry bool
	validTime time.Time
}

func NewCache() Cache {
	return Cache{make(map[string]Value)}
}

func (cache Cache) Get(key string) (string, bool) {

	val, exist := cache.data[key]

	if exist {

		if isExpired(val) {
			return "", false
		} else {
			return val.data, true
		}

	}

	return "", exist
}

func getData(value Value) (string, bool) {

	if isExpired(value) {
		return "", true
	}

	return value.data, false

}

func isExpired(value Value) bool {

	return value.hasExpiry && time.Now().After(value.validTime)

}

func (cache Cache) Put(key, value string) {

	va := Value{
		data:      value,
		hasExpiry: false}
	cache.data[key] = va
}

func (cache Cache) Keys() []string {

	output := make([]string, 0)
	for key, value := range cache.data {
		if isExpired(value) {
			delete(cache.data, key)
		} else {
			output = append(output, key)
		}
	}
	return output
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {

	va := Value{
		data:      value,
		hasExpiry: true,
		validTime: deadline}
	cache.data[key] = va
}
