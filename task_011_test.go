package xxhash
import "testing"
func TestTaskXXHash011BufferedLength(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abc")); return d.Sum64()==Sum64([]byte("abc")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash011BufferedLengthAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("buffered-length")); return d.Sum64()==Sum64([]byte("buffered-length")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
