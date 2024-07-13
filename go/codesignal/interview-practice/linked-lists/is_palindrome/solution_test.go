package ispalindrome

import (
	"linkedlist"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	l := &linkedlist.ListNode{Value: "a"}
	l.Add("b")
	l.Add("c")
	l.Add("b")
	l.Add("a")

	if !solution(l) {
		t.Fail()
	}
}

func TestIsPalindrome2(t *testing.T) {
	l := &linkedlist.ListNode{Value: "a"}
	l.Add("b")
	l.Add("c")
	l.Add("c")
	l.Add("b")
	l.Add("a")

	if !solution(l) {
		t.Fail()
	}
}

func TestIsPalindrome3(t *testing.T) {
	l := &linkedlist.ListNode{Value: "a"}

	if !solution(l) {
		t.Fail()
	}
}

func TestIsPalindromeNil(t *testing.T) {
	if !solution(nil) {
		t.Fail()
	}
}

func TestIsNotPalindrome(t *testing.T) {
	l := &linkedlist.ListNode{Value: "a"}
	l.Add("b")
	l.Add("c")
	l.Add("c")
	l.Add("b")
	l.Add("a")
	l.Add("x")

	if solution(l) {
		t.Fail()
	}
}
