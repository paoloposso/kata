package variablename

import "testing"

func Test1(t *testing.T) {
	if solution("1var_name_12") {
		t.Fail()
	}
	if solution("var_1name_12") {
		t.Fail()
	}
	if !solution("var__a15_name_12") {
		t.Fail()
	}
	if !solution("var_name_12") {
		t.Fail()
	}
}
