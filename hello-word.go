package go_utils

import "fmt"

func DoSome(name string) {
	fmt.Println(name + "asdfsadfasdfsafd.4.0")
}



func (m *Mgo) clone() *Mgo {
	clone := &Mgo{
		Collection: m.Collection,
		Pip:        make([]bson.D, 0),
	}
	for ix := range m.Pip {
		clone.Pip = append(clone.Pip, m.Pip[ix])
	}
	return clone
}
