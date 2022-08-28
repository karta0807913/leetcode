package main

import (
	"bytes"
)

func decodeMessage(key string, message string) string {
	m := make(map[byte]bool)
	r := make(map[byte]byte)
	b := bytes.NewBuffer(nil)
	c := byte('a')
	for i := range key {
		k := key[i]
		if !m[k] && k != ' ' {
			m[k] = true
			r[k] = c
			c += 1
		}
	}
	for i := range message {
		if message[i] == ' ' {
			b.WriteByte(' ')
		} else {
			b.WriteByte(r[message[i]])
		}
	}
	return b.String()
}
