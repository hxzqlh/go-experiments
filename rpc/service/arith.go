package service

type Args struct {
	A int
	B int
}

type Arith int

func (t *Arith) Add(a *Args, reply *int) error {
	*reply = a.A + a.B
	return nil
}
