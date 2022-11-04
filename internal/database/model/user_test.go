package model

import "testing"

func TestModel(t *testing.T) {
	user := User{UserName: "aaa"}
	test(&user)
}
