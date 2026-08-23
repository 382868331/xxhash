package xxhash
import "testing"
func TestTaskXXHash017LongInputFence(t *testing.T){got:=func() bool { b:=[]byte("12345678901234567890123456789012"); d:=New();d.Write(b);return d.Sum64()==Sum64(b) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash017LongInputFenceAdjacent(t *testing.T) {
	got := func() bool { b:=[]byte("abcdefghijklmnopqrstuvwxyzABCDEF"); d:=New();d.Write(b);return d.Sum64()==Sum64(b) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
