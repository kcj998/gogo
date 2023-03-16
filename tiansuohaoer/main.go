package main

import "fmt"

func main() {
	gay := I233{
		Sex:  1,
		Type: "NotGay",
		Child: &I233{
			Sex:  2,
			Type: "NotGay",
			Child: &I233{
				Sex:  3,
				Type: "NotGay",
				Child: &I233{
					Sex:   4,
					Type:  "Gay",
					Child: nil,
				},
			},
		},
	}
	gay = Isxianbeigay(gay)
	fmt.Printf("tiansuohaoer Sex %d is Gay", gay.Sex)
}

func Isxianbeigay(sex I233) I233 {
	if sex.IsGay() {
		return Isxianbeigay(*sex.Child)
	} else {
		return sex
	}
}

type I233 struct {
	Sex   int
	Type  string
	Child *I233
}
type I233classify interface {
	IsGay() bool
}

func (it *I233) IsGay() bool {
	if it.Type == "NotGay" {
		return true
	}
	return false
}
