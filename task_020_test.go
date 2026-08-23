package xxhash
import "testing"
func TestTaskXXHash020FourByteTailFence(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("1234")); return d.Sum64()==Sum64([]byte("1234")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash020FourByteTailFenceAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("wxyz")); return d.Sum64()==Sum64([]byte("wxyz")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
