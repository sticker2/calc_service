package calculation

import "testing"

func TestCalc(t *testing.T) {
    tests := []struct {
        expr     string
        expected float64
    }{
        {"2+2", 4},
        {"2+2*2", 6},
        {"(2+2)*2", 8},
        {"10/2", 5},
    }

    for _, test := range tests {
        result, err := Calc(test.expr)
        if err != nil {
            t.Errorf("unexpected error for %s: %v", test.expr, err)
        }
        if result != test.expected {
            t.Errorf("for %s, expected %f, got %f", test.expr, test.expected, result)
        }
    }
}
