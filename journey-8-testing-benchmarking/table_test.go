package main

import "testing"

func Perkalian(a, b int) int {
    return a * b
}

func TestPerkalian(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {2, 3, 6},
        {0, 5, 0},
        {4, -2, -8},
    }

    for _, tt := range tests {
        hasil := Perkalian(tt.a, tt.b)
        if hasil != tt.expected {
            t.Errorf("Perkalian(%d, %d) = %d; ingin %d", tt.a, tt.b, hasil, tt.expected)
        }
    }
}
