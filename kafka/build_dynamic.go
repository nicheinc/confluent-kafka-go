// +build !static
// +build !static_all

package kafka

// #cgo CFLAGS: -I${SRCDIR}/clib/include
// #cgo LDFLAGS: -L${SRCDIR}/clib/lib -lrdkafka
import "C"
