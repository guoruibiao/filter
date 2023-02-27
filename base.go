package main

type Handler func(input interface{}, operator string, value interface{}) (bool, error)
