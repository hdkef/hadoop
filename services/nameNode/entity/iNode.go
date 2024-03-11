package entity

type INode struct {
	id     string
	blocks []*BlockTarget
	hash   string
}
