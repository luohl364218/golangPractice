package source

import (
	"fmt"
	"errors"
	"math/rand"
	"hash/crc32"
)

//默认初始化方法，供别的方法调用
func init()  {
	RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {
}


func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	var defaultKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defaultKey = key[0]
	}

	if len(insts) == 0 {
		err = errors.New("instance is empty")
		return
	}

	lens := len(insts)
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defaultKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
