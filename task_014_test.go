package xxhash
import "testing"
func TestTaskXXHash014RemainingCopyOffset(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz0123456789x")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
