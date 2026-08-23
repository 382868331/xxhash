package xxhash
import "testing"
func TestTaskXXHash015RemainingLength(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash015RemainingLengthAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("tail-abcdefghijklmnopqrstuvwxyz0123456789")); return d.Sum64()==Sum64([]byte("tail-abcdefghijklmnopqrstuvwxyz0123456789")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
