package source

import "strconv"

//负载均衡主机实例
type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) *Instance {
	return &Instance{
		host:host,
		port:port,
	}
}

func (p *Instance) GetHost() string {
	return p.host
}

func (p *Instance) GetPort() int {
	return p.port
}

func (p *Instance) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}