package xxhash
import "testing"
func TestTaskXXHash008DigestBlockSize(t *testing.T){got:=New().BlockSize();if got!=32{t.Fatalf("got %v want %v",got,32)}}
