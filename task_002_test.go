package xxhash
import "testing"
func TestTaskXXHash002SeedLaneTwo(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCD")); return d.Sum64()==Sum64([]byte("0123456789abcdefghijklmnopqrstuvwxyzABCD")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
