package translation

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	s1 := "actually"
	expected := "akshully"
	dictionary := map[string]string {
		s1: expected,
	}
	t1 := Translate(s1, dictionary)
	if t1 != expected {
		t.Errorf("Wrong translation: %s, should be %s", t1, expected)
	}

	s2 := "kitten"
	t2 := Translate(s2, dictionary)
	if t2 != s2 {
		t.Errorf("Wrong translation: %s, should be %s", t2, s2)
	}
}