package controller

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/tungstenfabricconfig"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
	if err := tungstenfabricconfig.Add(m); err != nil {
		return err
	}
	/*
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	*/
	return nil
}
