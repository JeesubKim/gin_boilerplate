package main

import (
	"fmt"
	"testing"
)

func TestGetAllPost(t *testing.T) {
	postList := getAllPost()
	fmt.Println(postList)
	//How to test for this?
	//need to be checked that how to mock database
	// t.Fail()
}
