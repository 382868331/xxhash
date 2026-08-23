package xxhash
import "testing"
func TestTaskXXHash009WriteTotal(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abc")); return d.Sum64()==Sum64([]byte("abc")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash009WriteTotalAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("length-accounting")); return d.Sum64()==Sum64([]byte("length-accounting")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
