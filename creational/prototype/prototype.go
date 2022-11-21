package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeMgr struct {
	Prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeMgr {
	return &PrototypeMgr{
		Prototypes: make(map[string]Cloneable),
	}
}

func (p *PrototypeMgr) Get(name string) Cloneable {
	return p.Prototypes[name]
}

func (p *PrototypeMgr) Set(name string, prototype Cloneable) {
	p.Prototypes[name] = prototype
}
