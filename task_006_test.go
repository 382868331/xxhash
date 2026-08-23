package xxhash
import "testing"
func TestTaskXXHash006ResetBufferedCount(t *testing.T){got:=func() bool { d:=New(); return d.Sum64()==Sum64(nil) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
