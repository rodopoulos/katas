package main

type Cache interface {
	Add(value int)
	Get(value int) error
}
