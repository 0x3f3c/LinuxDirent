package main

import (
	"fmt"
	"unsafe"
)

type LinuxDirent struct {
	DIno    uint64
	DOff    uint64
	DReclen uint16
	DName   [256]byte
}

const MAGIC_PREFIX = "dopmaine_secret"
const PF_INVISIBLE =
const MODULE_NAME = "dopmaine"

const (
	SIGINVIS = iota + 31
	SIGSUPER
	SIGMODINVIS
)

func IS_ENABLED(option bool) bool {
	return option
}

func kallsymsLookupName(symbolName string) uintptr {
	fmt.Println("Looking up symbol:", symbolName)
	return uintptr(unsafe.Pointer(nil))
}

func main() {
	var dirent LinuxDirent
	dirent.DIno = 12345
	dirent.DReclen = uint16(unsafe.Sizeof(dirent))

	fmt.Printf("LinuxDirent: %+v\n", dirent)

	kpSymbolName := "kallsyms_lookup_name"
	address := kallsymsLookupName(kpSymbolName)
	fmt.Printf("Address of %s: %v\n", kpSymbolName, address)

	optionEnabled := IS_ENABLED(true)
	fmt.Println("Option enabled:", optionEnabled)
}
