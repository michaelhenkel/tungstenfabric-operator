package controller

import (
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/vrouter"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, vrouter.Add)
}