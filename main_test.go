package main

import (
	"strings"
	"testing"
)

func TestAddPhrase(t *testing.T) {
	root := NewNode(-1)
	addPhrase(root, "password is", 1)
	addPhrase(root, "aws credentials", 2)

	// there should be 2 children, one per phrase
	if len(root.Children) != 2 {
		t.Fail()
	}
}

func TestFindPhrase(t *testing.T) {
	root := NewNode(-1)
	addPhrase(root, "password is", 1)
	addPhrase(root, "aws credentials", 2)
	textRaw := `line 1
	line 2
	password isn't
	password is
	line 3
	aws credentials`
	n := strings.Replace(textRaw, "\n", "", -1)
	st := strings.Replace(n, "\t", " ", -1)
	found := FindPhrases(root, st)

	if len(found) != 2 {
		t.Fail()
	}
}
