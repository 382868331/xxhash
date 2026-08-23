package xxhash
import "testing"
func TestTaskXXHash019EightByteTailFence(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("12345678")); return d.Sum64()==Sum64([]byte("12345678")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
