package reflectiongo

import (
	"reflect"
	"sort"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}
	type PersonPtr struct {
		Name    string
		Profile *Profile
	}
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{Name: "asis"},
			[]string{"asis"},
		},
		{
			"struct with two string field",
			struct {
				Name  string
				Place string
			}{Name: "asis", Place: "Mumbai"},
			[]string{"asis", "Mumbai"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{Name: "asis", Age: 23},
			[]string{"asis"},
		},
		{
			"struct with nested fields",
			Person{
				"asis",
				Profile{23, "Mumbai"},
			},
			[]string{"asis", "Mumbai"},
		},
		{
			"struct with nested fields and pointers",
			&Person{
				"asis",
				Profile{23, "Mumbai"},
			},
			[]string{"asis", "Mumbai"},
		},
		{
			"struct with nested pointers",
			PersonPtr{
				"asis",
				&Profile{23, "Bombay"},
			},
			[]string{"asis", "Bombay"},
		},
		{
			"struct with slice",
			[]Profile{
				{33, "Bombay"},
				{33, "Pune"},
			},
			[]string{"Bombay", "Pune"},
		},
		{
			"struct with array",
			[2]Profile{
				{33, "Bombay"},
				{33, "Pune"},
			},
			[]string{"Bombay", "Pune"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(test.ExpectedCalls, got) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})

	}
	t.Run("maps", func(t *testing.T) {
		var got []string
		expectedCalls := []string{"Bombay", "Pune"}
		input := map[int]Profile{
			1: {33, "Bombay"},
			2: {33, "Pune"},
		}
		walk(input, func(input string) {
			got = append(got, input)
		})
		if len(got) != 2 {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
		if !reflect.DeepEqual(sort.StringSlice(expectedCalls), sort.StringSlice(got)) {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
	})

	t.Run("channel", func(t *testing.T) {
		var got []string
		expectedCalls := []string{"Delhi", "Pune"}
		ch := make(chan Profile)

		go func() {
			ch <- Profile{22, "Delhi"}
			ch <- Profile{23, "Pune"}
			close(ch)
		}()

		walk(ch, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expectedCalls) {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
	})
	t.Run("function", func(t *testing.T) {
		var got []string
		expectedCalls := []string{"Delhi", "Pune"}

		fn := func() (Profile, Profile) {
			return Profile{22, "Delhi"}, Profile{23, "Pune"}
		}

		walk(fn, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expectedCalls) {
			t.Errorf("got %v, want %v", got, expectedCalls)
		}
	})
}
