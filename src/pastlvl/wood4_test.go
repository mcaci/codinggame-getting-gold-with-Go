package main

import "testing"

func TestDistance(t *testing.T) {
	e := EntityInfo{x: 1, y: 2}
	f := EntityInfo{x: 12, y: 3}
	d := Distance(e, f)
	t.Log(d)
	t.Fail()
}
