package source

//实现一个负载均衡调度算法
type IBalancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}