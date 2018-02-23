package kafka

// NOTE(griffy): 
// We are simplifying how this package is built in a few ways for our own purposes:
// 1) Pre-built releases of librdkafka are vendored under the `clib` directory which we link against without the usage of pkg-config
// 2) The build tags (`static` and `static_all`) are removed in favor of the following default behavior:
//    - For Linux, librdkafka is statically linked using the included pre-built package file while its (stripped down) dependencies are not (this corresponds to the previous `static` build tag)
//    - For Mac/Windows, librdkafka is dynamically linked using the pre-built package file
// The goal with these changes is that a single `go get` will be all that is needed to get up and running.

// #cgo linux CFLAGS: -I${SRCDIR}/clib/include
// #cgo linux LDFLAGS: -L${SRCDIR}/clib/lib -Wl,-Bstatic -lrdkafka -Wl,-Bdynamic -lpthread -ldl
// #cgo darwin windows CFLAGS: -I${SRCDIR}/clib/include
// #cgo darwin windows LDFLAGS: -L${SRCDIR}/clib/lib -lrdkafka
import "C"
