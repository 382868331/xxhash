package xxhash
import "testing"
func TestTaskXXHash003SeedLaneThree(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("third-lane-abcdefghijklmnopqrstuvwxyz")); return d.Sum64()==Sum64([]byte("third-lane-abcdefghijklmnopqrstuvwxyz")) }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
