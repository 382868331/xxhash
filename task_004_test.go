package xxhash
import "testing"
func TestTaskXXHash004SeedLaneFour(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("fourth-lane-abcdefghijklmnopqrstuvwxyz")); return d.Sum64()==Sum64([]byte("fourth-lane-abcdefghijklmnopqrstuvwxyz")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash004SeedLaneFourAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz-fourth-lane")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz-fourth-lane")) }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
