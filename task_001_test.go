package xxhash
import "testing"
func TestTaskXXHash001SeedLaneOne(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz0123456789")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz0123456789")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash001SeedLaneOneAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("0123456789abcdefghijklmnopqrstuvwxyz")); return d.Sum64()==Sum64([]byte("0123456789abcdefghijklmnopqrstuvwxyz")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
