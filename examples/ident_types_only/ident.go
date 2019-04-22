package ident

type IdentPrim int

type Ident struct {
	String string

	Byte byte

	Int8    int8
	Uint8   uint8
	Int16   int16
	Uint16  uint16
	Int     int
	Uint    uint
	Int32   int32
	Uint32  uint32
	Int64   int64
	Uint64  uint64
	Uintptr uintptr

	Rune rune

	Float32 float32
	Float64 float64

	Complex64  complex64
	Complex128 complex128

	Bool bool

	IPrim IdentPrim

	Chan chan int

	PointerInt *int

	PointerPointerBool **bool

	SendChan    chan<- int
	ReceiveChan <-chan int

	SliceInt  []int
	ArrayInt2 [0xA]int

	MapIntChanString map[int]chan string
}
