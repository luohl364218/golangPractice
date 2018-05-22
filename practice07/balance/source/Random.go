package source

import (
	"errors"
	"math/rand"
)

//默认初始化方法，供别的方法调用
func init()  {
	RegisterBalancer("random", &RandomBalance{})
}

//随机调度算法 简单实现
type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("instance is empty")
		return
	}
	length := len(insts)
	id := rand.Intn(length)
	inst = insts[id]
	return
}

