package log

import (
	"github.com/johnweldon/go-misc/util"
)

type Project struct {
	ID   util.UUID `bson:"_id"`
	Name string
}

func (p Project) String() string {
	return p.Name
}

func NewProject(name string) Project {
	return Project{util.NewUUID(), name}
}
