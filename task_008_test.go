package xxhash
import "testing"
func TestTaskXXHash008DigestBlockSize(t *testing.T){got:=New().BlockSize();if got!=32{t.Fatalf("got %v want %v",got,32)}}
func TestTaskXXHash008DigestBlockSizeAdjacent(t *testing.T) {
	got := NewWithSeed(9).BlockSize()
	want := 32
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
