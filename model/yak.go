package model

type Herd struct {
	Name string  `json:"name"`
	Age  float32 `json:"age"`
	Sex  string  `json:"sex"`
}

type Yak struct {
	Herd []Herd `json:"herd"`
}

type Output struct {
	MilkTotal float32 `json:"milkTotal"`
	WoolTotal int64   `json:"woolTotal"`
	Yak
}
