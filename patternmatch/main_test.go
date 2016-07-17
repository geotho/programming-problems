package main

import "testing"

type testCase struct {
	pattern   string
	className string
}

func TestMatchPatterns(t *testing.T) {
	type tc struct {
		pattern    string
		classnames []string
		filtered   []string
	}
	successfuls := []tc{
		{
			"H",
			[]string{"Hello", "HelloWorld"},
			[]string{"Hello", "HelloWorld"},
		},
		{
			"HW",
			[]string{"Hello", "HelloWorld"},
			[]string{"HelloWorld"},
		},
		{
			"Hello",
			[]string{"Hello", "HelloWorld"},
			[]string{"Hello", "HelloWorld"},
		},
		{
			"Wor",
			[]string{"Hello", "HelloWorld"},
			[]string{},
		},
		{
			"HelloWorld",
			[]string{"Hello", "HelloWorld"},
			[]string{"HelloWorld"},
		},
		{
			"Foo",
			[]string{"Hello", "HelloWorld", "FooBar"},
			[]string{"FooBar"},
		},
		{
			"FooB",
			[]string{"Hello", "HelloWorld", "FooBar"},
			[]string{"FooBar"},
		},
		{
			"",
			[]string{"Hello", "HelloWorld", "FooBar"},
			[]string{"Hello", "HelloWorld", "FooBar"},
		},
	}

	for _, test := range successfuls {
		if m := MatchPatterns(test.pattern, test.classnames); !equal(m, test.filtered) {
			t.Errorf("pattern=%#v actual=%#v expected=%#v",
				test.pattern, m, test.filtered)
		}
	}
}

func TestMatchSucceeds(t *testing.T) {
	successfuls := []testCase{
		{
			"H",
			"Hello",
		},
		{
			"HW",
			"HelloWorld",
		},
		{
			"HeW",
			"HelloWorld",
		},
		{
			"HelloWorld",
			"HelloWorld",
		},
		{
			"",
			"HelloWorld",
		},
	}

	for _, test := range successfuls {
		if m := matchTest(test.pattern, test.className, true); !m {
			t.Errorf("Expected %v but got %v. pattern=%v, className=%v",
				true, false, test.pattern, test.className)
		}
	}
}

func TestMatchFails(t *testing.T) {
	successfuls := []testCase{
		{
			"ello",
			"Hello",
		},
		{
			"World",
			"HelloWorld",
		},
		{
			"W",
			"HelloWorld",
		},
		{
			"Ho",
			"HiHo",
		},
		{
			"Foo",
			"",
		},
	}

	for _, test := range successfuls {
		if m := matchTest(test.pattern, test.className, false); !m {
			t.Errorf("Expected %v but got %v. pattern=%v, className=%v",
				true, false, test.pattern, test.className)
		}
	}
}

func TestSplit(t *testing.T) {
	type splitCase struct {
		s      string
		splits []string
	}

	tests := []splitCase{
		{
			"HelloWorld",
			[]string{"Hello", "World"},
		},
		{
			"ABC",
			[]string{"A", "B", "C"},
		},
		{
			"",
			[]string{""},
		},
	}
	for _, test := range tests {
		if x := split(test.s); !equal(x, test.splits) {
			t.Errorf("%s: Got %#v but expected %#v", test.s, x, test.splits)
		}
	}
}

func matchTest(pattern, className string, expectSuccess bool) bool {
	return match(pattern, className) == expectSuccess
}

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 || len(s2) == 0 {
		return true
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}

	return true
}
