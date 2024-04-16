package slowmap

import (
	"log"
	"runtime"
	"testing"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%d KB\n", m.Alloc/1024)
}

func TestMap(t *testing.T) {
	nmap := NewSlowMap()
	t.Log("NewSlowMap")
	printAlloc()
	for i := range 1_000_000 {
		nmap.Set(i, [128]byte{})
	}
	t.Log("Set")
	printAlloc()

	for i := range 999999 {
		nmap.Del(i)
	}
	t.Log("Del")
	//t.Log(nmap.Get(999999))
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(nmap)
}
