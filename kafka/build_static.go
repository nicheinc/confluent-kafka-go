// +build static
// +build !static_all

package kafka

// #cgo linux CFLAGS: -I${SRCDIR}/clib/include
// #cgo linux LDFLAGS: -L${SRCDIR}/clib/lib -Wl,-Bstatic -lrdkafka -Wl,-Bdynamic -lpthread -ldl
// #cgo !linux pkg-config: --static rdkafka
// #cgo !linux LDFLAGS: -Wl,-Bstatic -lrdkafka -Wl,-Bdynamic
import "C"
