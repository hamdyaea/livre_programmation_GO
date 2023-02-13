package main

import "testing"

var tests = []struct {
    name string
    a float32
    b float32
    want float32
    isErr bool
}{
    {"valid", 10.0, 2.0, 5.0, false},
    {"invalid", 10.0, 0.0, 0.0, true},

}

func TestDivise(t *testing.T){
    for _, tt := range tests {
        want, err := divise(tt.a, tt.b)
        if (err != nil) != tt.isErr {
            t.Errorf("%s: error %v, want error %v", tt.name, err, tt.isErr)
        }

        if want != tt.want {
            t.Errorf("%s: is %v, want %v", tt.name, want, tt.want)
        }
    }
}
