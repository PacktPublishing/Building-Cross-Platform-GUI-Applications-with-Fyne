package main

/*
#include <stdio.h>
#include <stdlib.h>
void print_hello(const char *name) {
    printf("Hello %s!\n", name);
}
*/
import "C"
import "unsafe"

func main() {
	cName := C.CString("World")
	C.print_hello(cName)

	C.free(unsafe.Pointer(cName))
}
