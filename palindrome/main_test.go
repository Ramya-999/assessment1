package main

import "testing"

func TestReverse(t *testing.T){

  if Reverse("radar") != "radar"{
	  t.Error("Expecting palindrome")
  }
  if Reverse("morning") != "gninrom"{
	t.Error("Expecting NOt palindrome")
  }
}

