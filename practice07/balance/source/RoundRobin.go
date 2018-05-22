package source

import "errors"

//默认初始化方法，供别的方法调用
func init()  {
	RegisterBalancer("roundRobin", &RoundRobin{})
}

type RoundRobin struct {
	curIndex int
}
//轮询 简单示例
func (p *RoundRobin) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("instance is empty")
		return
	}
	len := len(insts)
	if p.curIndex >= len {
		p.curIndex = 0
	}

	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % len
	return
}
