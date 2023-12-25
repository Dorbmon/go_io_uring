package goiouring

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	SYS_IO_URING_SETUP    = 425
	SYS_IO_URING_ENTER    = 426
	SYS_IO_URING_REGISTER = 427
)

type IOSqringOffsets struct {
	head         uint32
	tail         uint32
	ring_mask    uint32
	ring_entries uint32
	flags        uint32
	dropped      uint32
	array        uint32
	resv1        uint32
	resv2        uint64
}

type IOCqringOffsets struct {
	head         uint32
	tail         uint32
	ring_mask    uint32
	ring_entries uint32
	overflow     uint32
	cqes         uint32
	resv         [2]uint64
}

type IOUringParams struct {
	SqEntries    uint32
	CqEntries    uint32
	Flags        uint32
	SqThreadCPU  uint32
	SqThreadIdle uint32
	Features     uint32
	WqFD         uint32
	Resv         [3]uint32
	SqOff        IOSqringOffsets
	CqOff        IOCqringOffsets
}

func IOUringSetup(entries uint32, params *IOUringParams) (fd uintptr, err error) {
	fd, _, err = syscall.Syscall(SYS_IO_URING_SETUP, uintptr(entries), uintptr(unsafe.Pointer(params)), 0)
	return
}

func IOUringEnter(fd uint, toSubmit uint32, minComplete uint32, flags uint32, sig *unix.Sigset_t, sigsz uint) (num uintptr, err error) {
	num, _, err = syscall.Syscall6(SYS_IO_URING_ENTER, uintptr(fd), uintptr(toSubmit), uintptr(minComplete), uintptr(flags), uintptr(unsafe.Pointer(sig)), uintptr(sigsz))
	return
}
func IOUringRegister(fd uint, opcode uint32, arg uintptr, nrArgs uint32) (res uintptr, err error) {
	res, _, err = syscall.Syscall6(SYS_IO_URING_REGISTER, uintptr(fd), uintptr(opcode), arg, uintptr(nrArgs), 0, 0)
	return
}
