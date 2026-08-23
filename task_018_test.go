package xxhash
import "testing"
func TestTaskXXHash018LengthAvalanche(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("length-avalanche")); return d.Sum64()==Sum64([]byte("length-avalanche")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
