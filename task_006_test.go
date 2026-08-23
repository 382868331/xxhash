package xxhash
import "testing"
func TestTaskXXHash006ResetBufferedCount(t *testing.T){got:=func() bool { d:=New(); return d.Sum64()==Sum64(nil) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash006ResetBufferedCountAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("old")); d.Reset(); d.Write([]byte("x")); return d.Sum64()==Sum64([]byte("x")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
