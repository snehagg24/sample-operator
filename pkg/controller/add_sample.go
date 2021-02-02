package controller

import (
	"github.com/snehakpersistent/sample-operator/pkg/controller/sample"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, sample.Add)
}
