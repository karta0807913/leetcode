package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func simplifyPath(path string) string {
	stack := make([]int, 0)
	bpath := make([]byte, 0, len(path))
	path += "/"
	for idx := range path {
		if path[idx] == '/' {
			prev := 0
			if len(stack) != 0 {
				n := len(stack) - 1
				prev = stack[n]
				stack = stack[:n]
			}
			if prev == len(bpath) || prev+1 == len(bpath) || bytes.Equal(bpath[prev:], []byte("/.")) {
				bpath = bpath[:prev]
			} else if bytes.Equal(bpath[prev:], []byte("/..")) {
				prev = 0
				if len(stack) != 0 {
					n := len(stack) - 1
					prev = stack[n]
					stack = stack[:n]
				}
				bpath = bpath[:prev]
			} else {
				stack = append(stack, prev)
			}
			stack = append(stack, len(bpath))
		}
		bpath = append(bpath, path[idx])
	}
	if len(bpath) != 1 {
		bpath = bpath[:len(bpath)-1]
	}
	return string(bpath)
}

func TestPath(t *testing.T) {
	assert.Equal(t, "/home", simplifyPath("/home/"))
	assert.Equal(t, "/home", simplifyPath("//home/"))
	assert.Equal(t, "/home/user", simplifyPath("//home//user"))
	assert.Equal(t, "/home/user", simplifyPath("//home//user//////"))
	assert.Equal(t, "/", simplifyPath("/."))
	assert.Equal(t, "/", simplifyPath("/.."))
	assert.Equal(t, "/home", simplifyPath("/home/user/.."))
	assert.Equal(t, "/a/b", simplifyPath("/a//b/c/./d/e//../..//../"))
	assert.Equal(t, "/home", simplifyPath("/home"))
}
