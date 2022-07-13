package models

type Request struct {
	request interface{}
}

func (r *Request) Set(i interface{}) {
	r.request = i
}

func (r Request) Get() interface{} {
	return r.request
}
