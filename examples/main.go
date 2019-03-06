package main

import (

)

func main() {
	id := "100111"
	a := New(id, "ZhangSan", 100)
	a.Query(id)
	a.Update(id, 500)
}

func init() {
	New = func(id, name string, value int) Account {
		a := &AccountImpl{id, name, value}
		p := &Proxy{a}
		return p
	}
}