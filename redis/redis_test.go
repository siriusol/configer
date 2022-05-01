package redis

import "testing"

func TestGetKeys(t *testing.T) {
	c := New(&Option{})
	keys, err := c.GetKeys()
	if err != nil {
		panic(err)
	}
	t.Log(keys)
}
