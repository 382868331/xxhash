package xxhash
import "testing"
func TestTaskXXHash013CompletedBlockReset(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abc")); d.Write([]byte("defghijklmnopqrstuvwxyz0123456789")); return d.Sum64()==Sum64([]byte("abcdefghijklmnopqrstuvwxyz0123456789")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
