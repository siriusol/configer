package main

type Getter interface {
	Get(string) ([]byte, error)
}
