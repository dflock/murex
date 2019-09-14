package element

import (
	"testing"

	_ "github.com/lmorg/murex/builtins/core/io"
	_ "github.com/lmorg/murex/builtins/types/json"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

func jsonStr() string {
	s := `
	{
		"StrArray": [ "foobar", "a", "b", "-2", "-1", "0", "1", "2", "00100" ],
		"IntArray": [ -2, -1, 0, 1, 2, 3, 4, 5, 100 ],
		"StrMap":   { "foo": "bar", "bar": "foo", "1": "00100", "00100": "1" },
		"IntMap":   { "-2": -2, "-1": -1, "0": 0, "00100": 100 },
		"Nested":   { "self": 5314, "fork": "fork" }
	}`

	//s = strings.ReplaceAll(s, "5314", s)
	//s = strings.ReplaceAll(s, "5314", s)

	return s
}

func TestElementPositive(t *testing.T) {
	params := [][]string{
		{"/StrArray]]"},
		{"/StrArray/]]"},
		{"/IntArray]]"},
		{"/IntArray/]]"},
		{"/StrMap]]"},
		{"/StrMap/]]"},
		{"/IntMap]]"},
		{"/IntMap/]]"},
		{"/Nested/fork]]"},
		{"/Nested/fork/]]"},

		{",StrArray]]"},
		{",StrArray,]]"},
		{",IntArray]]"},
		{",IntArray,]]"},
		{",StrMap]]"},
		{",StrMap,]]"},
		{",IntMap]]"},
		{",IntMap,]]"},
		{",Nested,fork]]"},
		{",Nested,fork,]]"},

		{"#StrArray]]"},
		{"#StrArray#]]"},
		{"#IntArray]]"},
		{"#IntArray#]]"},
		{"#StrMap]]"},
		{"#StrMap#]]"},
		{"#IntMap]]"},
		{"#IntMap#]]"},
		{"#Nested#fork]]"},
		{"#Nested#fork#]]"},

		{"{StrArray]]"},
		{"{StrArray{]]"},
		{"{IntArray]]"},
		{"{IntArray{]]"},
		{"{StrMap]]"},
		{"{StrMap{]]"},
		{"{IntMap]]"},
		{"{IntMap{]]"},
		{"{Nested{fork]]"},
		{"{Nested{fork{]]"},

		/*{"世StrArray]]"},
		{"世StrArray世]]"},
		{"世IntArray]]"},
		{"世IntArray世]]"},
		{"世StrMap]]"},
		{"世StrMap世]]"},
		{"世IntMap]]"},
		{"世IntMap世]]"},
		{"世Nested世fork]]"},
		{"世Nested世fork世]]"},*/

		{"/StrArray", "]]"},
		{"/StrArray/", "]]"},
		{"/IntArray", "]]"},
		{"/IntArray/", "]]"},
		{"/StrMap", "]]"},
		{"/StrMap/", "]]"},
		{"/IntMap", "]]"},
		{"/IntMap/", "]]"},
		{"/Nested/fork", "]]"},
		{"/Nested/fork/", "]]"},

		{",StrArray", "]]"},
		{",StrArray,", "]]"},
		{",IntArray", "]]"},
		{",IntArray,", "]]"},
		{",StrMap", "]]"},
		{",StrMap,", "]]"},
		{",IntMap", "]]"},
		{",IntMap,", "]]"},
		{",Nested,fork", "]]"},
		{",Nested,fork,", "]]"},

		{"#StrArray", "]]"},
		{"#StrArray#", "]]"},
		{"#IntArray", "]]"},
		{"#IntArray#", "]]"},
		{"#StrMap", "]]"},
		{"#StrMap#", "]]"},
		{"#IntMap", "]]"},
		{"#IntMap#", "]]"},
		{"#Nested#fork", "]]"},
		{"#Nested#fork#", "]]"},

		{"{StrArray", "]]"},
		{"{StrArray{", "]]"},
		{"{IntArray", "]]"},
		{"{IntArray{", "]]"},
		{"{StrMap", "]]"},
		{"{StrMap{", "]]"},
		{"{IntMap", "]]"},
		{"{IntMap{", "]]"},
		{"{Nested{fork", "]]"},
		{"{Nested{fork{", "]]"},

		/*{"世StrArray","]]"},
		{"世StrArray世","]]"},
		{"世IntArray","]]"},
		{"世IntArray世","]]"},
		{"世StrMap","]]"},
		{"世StrMap世","]]"},
		{"世IntMap","]]"},
		{"世IntMap世","]]"},
		{"世Nested世fork","]]"},
		{"世Nested世fork世","]]"},*/
	}

	expected := []string{
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		/*`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,*/

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,

		/*`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`["foobar","a","b","-2","-1","0","1","2","00100"]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`[-2,-1,0,1,2,3,4,5,100]`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"00100":"1","1":"00100","bar":"foo","foo":"bar"}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`{"-1":-1,"-2":-2,"0":0,"00100":100}`,
		`fork`,
		`fork`,*/
	}

	for i := range params {
		test.RunMethodTest(
			t, element, "[[", jsonStr(), types.Json, params[i], expected[i], nil)
	}
}

/*func TestElementNegative(t *testing.T) {
	params := [][]string{
		{"StrArray]]"},
		{"/StrArray/]"},
		{"/IntArray"},
	}

	expected := []string{
		`Key 'trArray' not found`,
		"Missing closing brackets, ` ]]`",
		"Missing closing brackets, ` ]]`",
	}

	for i := range params {
		test.RunMethodTest(
			t, element, "[[", jsonStr(), types.Json, params[i], "", errors.New(expected[i]))
	}
}*/