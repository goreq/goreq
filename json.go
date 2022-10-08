package goreq

type JsonEncoder interface {
	Encode(interface{}) error
}

type JsonDecoder interface {
	Decode(interface{}) error
}
