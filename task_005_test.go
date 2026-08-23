package xxhash
import "testing"
func TestTaskXXHash005ResetTotal(t *testing.T){got:=func() bool { d:=New(); return d.Sum64()==Sum64(nil) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash005ResetTotalAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("old")); d.Reset(); return d.Sum64()==Sum64(nil) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
