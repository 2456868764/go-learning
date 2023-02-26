package unsafe

import (
	"fmt"
	"github.com/2456868764/go-learning/advance/unsafe/types"
	"testing"
	"unsafe"
)

func TestOutputFieldLayout(t *testing.T) {
	fmt.Printf("size: %d \n", unsafe.Sizeof(types.BlogV1{}))
	OutputFieldLayout(types.BlogV1{})

	fmt.Printf("size: %d \n", unsafe.Sizeof(types.BlogV2{}))
	OutputFieldLayout(types.BlogV2{})

	fmt.Printf("size: %d \n", unsafe.Sizeof(types.BlogV3{}))
	OutputFieldLayout(types.BlogV3{})
}
