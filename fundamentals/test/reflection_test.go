/*
	golang challenge: write a function walk(x interface{}, fn func(string))
	which takes a struct x and calls fn for all strings fields found inside.
	difficulty level: recursively.
	To do this we will need to use reflection.
	Reflection in computing is the ability of a program to examine its own structure,
	particularly through types; it's a form of metaprogramming.
	It's also a great source of confusion.
*/

package fundamentalstest

import (
	"reflect"
	"testing"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem() // mengambil value dari pointer
	}

	return val
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"struct with one field string",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with two field string",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 30},
			[]string{"Chris"},
		},
		{
			"struct with nested fields",
			Person{
				"Chris",
				Profile{30, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"struct with pointer to things",
			&Person{
				"Chris",
				Profile{30, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices struct",
			[]Profile{
				{30, "London"},
				{20, "Simanabun"},
			},
			[]string{"London", "Simanabun"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Error("got", got, "want", test.Expected)
			}
		})
	}

	// map kita pisah karena kadang key nya tidak sesuai order
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Name": "Dhany",
			"City": "London",
		}

		got := []string{}
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Dhany")
		assertContains(t, got, "London")
	})

	t.Run("with chan", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{30, "London"}
			aChannel <- Profile{30, "Simanabun"}
			close(aChannel)
		}()

		got := []string{}
		want := []string{"London", "Simanabun"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Error("want", want, "got", got)
		}
	})

	t.Run("with func", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{30, "London"}, Profile{20, "Simanabun"}
		}

		got := []string{}
		want := []string{"London", "Simanabun"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Error("want", want, "got", got)
		}
	})
}

// fungsi untuk mengecek satu2 kedalam
func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Error("expected", needle, "to containt, but it did't")
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
