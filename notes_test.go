package main

import (
    "testing"
)

type testpair struct {
    values []int
    addition int
}

var addTests = []testpair{
    { []int{1, 2, 3, 4}, 10},
    { []int{2, 3, 4, 5}, 14},
}

func Testadd(t *testing.T) {
    for _, pair := range addTests {
        v := add(pair.values...)
        if v != pair.addition {
            t.Error(
                "For", pair.values,
                "expected", pair.addition,
                "got", v,
            )
        }
    }
}
