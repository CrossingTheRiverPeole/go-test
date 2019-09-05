package iota_test

import (
	"fmt"
	"testing"
)

type ByteSize float64
const (
	_ =iota
	KB  ByteSize = 1 << ( iota * 10)
	MB
	GB
	TB
)

func TestIota(t *testing.T)  {
	fmt.Println(KB, MB, GB, TB)

}
