package xxhash
import "testing"
func TestTaskXXHash007DigestSize(t *testing.T){got:=New().Size();if got!=8{t.Fatalf("got %v want %v",got,8)}}
