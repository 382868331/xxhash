package xxhash
import "testing"
func TestTaskXXHash010ExactBlockFence(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("12345678901234567890123456789012")); return d.Sum64()==Sum64([]byte("12345678901234567890123456789012")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash010ExactBlockFenceAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyzABCDEF")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyzABCDEF")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
