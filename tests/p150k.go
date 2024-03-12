package tests

//easyjson:json
type Rsp struct {
	A []byte `json:"a"`
}

//easyjson:json
type S2 struct {
	A int  `json:"a,string"`
	B bool `json:"b,string"`
}
