package redis

import "testing"

func TestGetKeys(t *testing.T) {
	keys, err := GetKeys()
	if err != nil {
		panic(err)
	}
	t.Log(keys)
}
