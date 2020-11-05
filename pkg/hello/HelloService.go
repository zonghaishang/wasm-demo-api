package hello

import (
	"fmt"
	"os"
	"unsafe"
)

//export proxy_on_plugin_hello
func ProxyOnPluginHello(context unsafe.Pointer, ptr int32, len int32, bufPtr int32, size int32) {
	// allocate memory

	// we invoke host service
	// rawhost.ProxyHello(ctx, ptr, len)

	fmt.Fprintf(os.Stdout, "ptr: %v, len: %v, bufPtr: %v, size: %v", ptr, len, bufPtr, size)
}
