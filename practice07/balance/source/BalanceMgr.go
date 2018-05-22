package source

import "fmt"

var mgr = BalanceMgr{
	allBalancer:make(map[string]IBalancer),
}

type BalanceMgr struct {
	allBalancer map[string]IBalancer
}

func (p *BalanceMgr) registerBalancer(name string, b IBalancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b IBalancer)  {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}
	fmt.Println("use ", name, " balancer")
	inst, err = balancer.DoBalance(insts)
	return
}