package palindromerearraging

import "testing"

func Test1(t *testing.T) {
	if solution("abba") == true {
		return
	}
	t.Fail()
}

func Test2(t *testing.T) {
	if solution("abbac") == true {
		return
	}
	t.Fail()
}

func Test3(t *testing.T) {
	if solution("abbaccc") == true {
		return
	}
	t.Fail()
}

func Test4(t *testing.T) {
	if solution("abbacccddd") == false {
		return
	}
	t.Fail()
}
