package xxhash
import "testing"
func TestTaskXXHash002SeedLaneTwo(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCD")); return d.Sum64()==Sum64([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCD")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash002SeedLaneTwoAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("lane-two-boundary-abcdefghijklmnopqrstuvwxyz")); return d.Sum64()==Sum64([]byte("lane-two-boundary-abcdefghijklmnopqrstuvwxyz")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
