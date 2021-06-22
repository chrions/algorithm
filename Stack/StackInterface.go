package Stack

//20,155,232,844,224,682,496.
type Statck interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush()
}
