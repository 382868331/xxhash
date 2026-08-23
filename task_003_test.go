package xxhash
import "testing"
func TestTaskXXHash003SeedLaneThree(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("third-lane-abcdefghijklmnopqrstuvwxyz")); return d.Sum64()==Sum64([]byte("third-lane-abcdefghijklmnopqrstuvwxyz")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash003SeedLaneThreeAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz-third-lane")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz-third-lane")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
