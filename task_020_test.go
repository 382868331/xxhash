package xxhash
import "testing"
func TestTaskXXHash020FourByteTailFence(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("1234")); return d.Sum64()==Sum64([]byte("1234")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
