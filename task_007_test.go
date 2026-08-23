package xxhash
import "testing"
func TestTaskXXHash007DigestSize(t *testing.T){got:=New().Size();if got!=8{t.Fatalf("got %v want %v",got,8)}}
func TestTaskXXHash007DigestSizeAdjacent(t *testing.T) {
	got := NewWithSeed(7).Size()
	want := 8
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
