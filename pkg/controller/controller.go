package controller

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
<<<<<<< HEAD
=======
	"github.com/michaelhenkel/tungstenfabric-operator/pkg/controller/tungstenfabricmanager"
>>>>>>> v0.0.4
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
<<<<<<< HEAD
=======
	if err := tungstenfabricmanager.Add(m); err != nil {
		return err
	}
	/*
>>>>>>> v0.0.4
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
<<<<<<< HEAD
=======
	*/
>>>>>>> v0.0.4
	return nil
}
