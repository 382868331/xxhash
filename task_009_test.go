package xxhash
import "testing"
func TestTaskXXHash009WriteTotal(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abc")); return d.Sum64()==Sum64([]byte("abc")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
