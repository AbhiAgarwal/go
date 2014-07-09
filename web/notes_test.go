package main

import "testing"

type testpair struct {
    values []float64
    average float64
}

var tests = []testpair{
    {[]int{1,2, 3, 4}},
}

func Testadd(t *testing.T) {
    for _, pair := range tests {
        v := add(pair.values)
        if v != pair.average {
            t.Error(
                "For", pair.values,
                "expected", pair.average,
                "got", v,
            )
        }
    }
}
