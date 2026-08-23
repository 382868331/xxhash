package xxhash
import "testing"
func TestTaskXXHash014RemainingCopyOffset(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash014RemainingCopyOffsetAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("0123456789abcdefghijklmnopqrstuvwxyz-tail")); return d.Sum64()==Sum64([]byte("0123456789abcdefghijklmnopqrstuvwxyz-tail")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
