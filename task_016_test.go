package xxhash
import("encoding/binary";"testing")
func TestTaskXXHash016SumByteOrder(t *testing.T){got:=func() bool { d:=New(); d.Write([]byte("abc")); b:=d.Sum(nil); return len(b)==8 && binary.BigEndian.Uint64(b)==d.Sum64() }();if got!=true{t.Fatalf("got %v want %v",got,true)}}
func TestTaskXXHash016SumByteOrderAdjacent(t *testing.T) {
	got := func() bool { d:=New(); d.Write([]byte("different")); b:=d.Sum([]byte{9}); return len(b)==9 && binary.BigEndian.Uint64(b[1:])==d.Sum64() }()
	want := true
	if got != want {
		t.Fatalf("adjacent got %v want %v", got, want)
	}
}
