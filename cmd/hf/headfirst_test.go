package main

import "testing"

func TestTwoElements(t *testing.T) {
	phrases := []string{"my parents", "a rodeo clown"}
	assert := "my parents and a rodeo clown"
	result := JoinWithCommas(phrases)

	if assert != result {
		t.Error("func joinWithCommas returns noncorrect result: ", result, " expected: ", assert)
	}
}

func TestThreeElements(t *testing.T) {
	phrases := []string{"my parents", "a rodeo clown", "a prize bull"}
	assert := "my parents, a rodeo clown, and a prize bull"
	result := JoinWithCommas(phrases)

	if assert != result {
		t.Error("func joinWithCommas returns noncorrect result: ", result, " expected: ", assert)
	}
	//t.Error("no test here either")
}
