package main
import "testing"
import "reflect"

func TestSplitWrongSep(t *testing.T) {
    got := []string{"a/b/c"}
    want := []string{"a/b/c"}
    if !reflect.DeepEqual(want, got) {
        t.Fatalf("expected: %v, got: %v", want, got)
    }
}