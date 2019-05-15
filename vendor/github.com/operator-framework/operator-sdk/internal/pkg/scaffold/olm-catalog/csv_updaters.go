// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package catalog

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> v0.0.4

	"github.com/ghodss/yaml"
	olmapiv1alpha1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
	olminstall "github.com/operator-framework/operator-lifecycle-manager/pkg/controller/install"
	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// CSVUpdater is an interface for any data that can be in a CSV, which will be
// set to the corresponding field on Apply().
type CSVUpdater interface {
	// Apply applies a data update to a CSV argument.
	Apply(*olmapiv1alpha1.ClusterServiceVersion) error
}

type updaterStore struct {
<<<<<<< HEAD
	installStrategy *CSVInstallStrategyUpdate
	crdUpdate       *CSVCustomResourceDefinitionsUpdate
=======
	installStrategy *InstallStrategyUpdate
	crds            *CustomResourceDefinitionsUpdate
	almExamples     *ALMExamplesUpdate
>>>>>>> v0.0.4
}

func NewUpdaterStore() *updaterStore {
	return &updaterStore{
<<<<<<< HEAD
		installStrategy: &CSVInstallStrategyUpdate{
			&olminstall.StrategyDetailsDeployment{},
		},
		crdUpdate: &CSVCustomResourceDefinitionsUpdate{
			&olmapiv1alpha1.CustomResourceDefinitions{},
		},
=======
		installStrategy: &InstallStrategyUpdate{
			&olminstall.StrategyDetailsDeployment{},
		},
		crds: &CustomResourceDefinitionsUpdate{
			&olmapiv1alpha1.CustomResourceDefinitions{},
			make(map[string]struct{}),
		},
		almExamples: &ALMExamplesUpdate{},
>>>>>>> v0.0.4
	}
}

// Apply iteratively calls each stored CSVUpdater's Apply() method.
func (s *updaterStore) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) error {
<<<<<<< HEAD
	for _, updater := range []CSVUpdater{s.installStrategy, s.crdUpdate} {
=======
	updaters := []CSVUpdater{s.installStrategy, s.crds, s.almExamples}
	for _, updater := range updaters {
>>>>>>> v0.0.4
		if err := updater.Apply(csv); err != nil {
			return err
		}
	}
	return nil
}

func getKindfromYAML(yamlData []byte) (string, error) {
<<<<<<< HEAD
	// Get Kind for inital categorization.
=======
>>>>>>> v0.0.4
	var temp struct {
		Kind string
	}
	if err := yaml.Unmarshal(yamlData, &temp); err != nil {
		return "", err
	}
	return temp.Kind, nil
}

<<<<<<< HEAD
func (s *updaterStore) AddToUpdater(yamlSpec []byte) error {
	kind, err := getKindfromYAML(yamlSpec)
	if err != nil {
		return err
	}

	switch kind {
	case "Role":
		return s.AddRole(yamlSpec)
	case "ClusterRole":
		return s.AddClusterRole(yamlSpec)
	case "Deployment":
		return s.AddDeploymentSpec(yamlSpec)
	case "CustomResourceDefinition":
		// All CRD's present will be 'owned'.
		return s.AddOwnedCRD(yamlSpec)
	}
	return nil
}

type CSVInstallStrategyUpdate struct {
=======
func (s *updaterStore) AddToUpdater(yamlSpec []byte, kind string) (found bool, err error) {
	found = true
	switch kind {
	case "Role":
		err = s.AddRole(yamlSpec)
	case "ClusterRole":
		err = s.AddClusterRole(yamlSpec)
	case "Deployment":
		err = s.AddDeploymentSpec(yamlSpec)
	case "CustomResourceDefinition":
		// All CRD's present will be 'owned'.
		err = s.AddOwnedCRD(yamlSpec)
	default:
		found = false
	}
	return found, err
}

type InstallStrategyUpdate struct {
>>>>>>> v0.0.4
	*olminstall.StrategyDetailsDeployment
}

func (store *updaterStore) AddRole(yamlDoc []byte) error {
	role := &rbacv1.Role{}
	if err := yaml.Unmarshal(yamlDoc, role); err != nil {
		return err
	}
	perm := olminstall.StrategyDeploymentPermissions{
		ServiceAccountName: role.ObjectMeta.Name,
		Rules:              role.Rules,
	}
	store.installStrategy.Permissions = append(store.installStrategy.Permissions, perm)

	return nil
}

func (store *updaterStore) AddClusterRole(yamlDoc []byte) error {
	clusterRole := &rbacv1.ClusterRole{}
	if err := yaml.Unmarshal(yamlDoc, clusterRole); err != nil {
		return err
	}
	perm := olminstall.StrategyDeploymentPermissions{
		ServiceAccountName: clusterRole.ObjectMeta.Name,
		Rules:              clusterRole.Rules,
	}
	store.installStrategy.ClusterPermissions = append(store.installStrategy.ClusterPermissions, perm)

	return nil
}

func (store *updaterStore) AddDeploymentSpec(yamlDoc []byte) error {
	dep := &appsv1.Deployment{}
	if err := yaml.Unmarshal(yamlDoc, dep); err != nil {
		return err
	}
	depSpec := olminstall.StrategyDeploymentSpec{
		Name: dep.ObjectMeta.Name,
		Spec: dep.Spec,
	}
	store.installStrategy.DeploymentSpecs = append(store.installStrategy.DeploymentSpecs, depSpec)

	return nil
}

<<<<<<< HEAD
func (u *CSVInstallStrategyUpdate) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) (err error) {
=======
func (u *InstallStrategyUpdate) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) (err error) {
>>>>>>> v0.0.4
	// Get install strategy from csv. Default to a deployment strategy if none found.
	var strat olminstall.Strategy
	if csv.Spec.InstallStrategy.StrategyName == "" {
		csv.Spec.InstallStrategy.StrategyName = olminstall.InstallStrategyNameDeployment
		strat = &olminstall.StrategyDetailsDeployment{}
	} else {
		var resolver *olminstall.StrategyResolver
		strat, err = resolver.UnmarshalStrategy(csv.Spec.InstallStrategy)
		if err != nil {
			return err
		}
	}

	switch s := strat.(type) {
	case *olminstall.StrategyDetailsDeployment:
		// Update permissions and deployments.
		u.updatePermissions(s)
		u.updateClusterPermissions(s)
		u.updateDeploymentSpecs(s)
	default:
		return fmt.Errorf("install strategy (%v) of unknown type", strat)
	}

	// Re-serialize permissions into csv strategy.
	updatedStrat, err := json.Marshal(strat)
	if err != nil {
		return err
	}
	csv.Spec.InstallStrategy.StrategySpecRaw = updatedStrat

	return nil
}

<<<<<<< HEAD
func (u *CSVInstallStrategyUpdate) updatePermissions(strat *olminstall.StrategyDetailsDeployment) {
=======
func (u *InstallStrategyUpdate) updatePermissions(strat *olminstall.StrategyDetailsDeployment) {
>>>>>>> v0.0.4
	if len(u.Permissions) != 0 {
		strat.Permissions = u.Permissions
	}
}

<<<<<<< HEAD
func (u *CSVInstallStrategyUpdate) updateClusterPermissions(strat *olminstall.StrategyDetailsDeployment) {
=======
func (u *InstallStrategyUpdate) updateClusterPermissions(strat *olminstall.StrategyDetailsDeployment) {
>>>>>>> v0.0.4
	if len(u.ClusterPermissions) != 0 {
		strat.ClusterPermissions = u.ClusterPermissions
	}
}

<<<<<<< HEAD
func (u *CSVInstallStrategyUpdate) updateDeploymentSpecs(strat *olminstall.StrategyDetailsDeployment) {
=======
func (u *InstallStrategyUpdate) updateDeploymentSpecs(strat *olminstall.StrategyDetailsDeployment) {
>>>>>>> v0.0.4
	if len(u.DeploymentSpecs) != 0 {
		strat.DeploymentSpecs = u.DeploymentSpecs
	}
}

<<<<<<< HEAD
type CSVCustomResourceDefinitionsUpdate struct {
	*olmapiv1alpha1.CustomResourceDefinitions
=======
type CustomResourceDefinitionsUpdate struct {
	*olmapiv1alpha1.CustomResourceDefinitions
	crKinds map[string]struct{}
>>>>>>> v0.0.4
}

func (store *updaterStore) AddOwnedCRD(yamlDoc []byte) error {
	crd := &apiextv1beta1.CustomResourceDefinition{}
	if err := yaml.Unmarshal(yamlDoc, crd); err != nil {
		return err
	}
<<<<<<< HEAD
	store.crdUpdate.Owned = append(store.crdUpdate.Owned, olmapiv1alpha1.CRDDescription{
=======
	store.crds.Owned = append(store.crds.Owned, olmapiv1alpha1.CRDDescription{
>>>>>>> v0.0.4
		Name:    crd.ObjectMeta.Name,
		Version: crd.Spec.Version,
		Kind:    crd.Spec.Names.Kind,
	})
<<<<<<< HEAD
=======
	store.crds.crKinds[crd.Spec.Names.Kind] = struct{}{}
>>>>>>> v0.0.4
	return nil
}

// Apply updates csv's "owned" CRDDescriptions. "required" CRDDescriptions are
// left as-is, since they are user-defined values.
<<<<<<< HEAD
func (u *CSVCustomResourceDefinitionsUpdate) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) error {
=======
func (u *CustomResourceDefinitionsUpdate) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) error {
>>>>>>> v0.0.4
	set := make(map[string]olmapiv1alpha1.CRDDescription)
	for _, csvDesc := range csv.Spec.CustomResourceDefinitions.Owned {
		set[csvDesc.Name] = csvDesc
	}
	du := u.DeepCopy()
	for i, uDesc := range u.Owned {
		if csvDesc, ok := set[uDesc.Name]; ok {
			csvDesc.Name = uDesc.Name
			csvDesc.Version = uDesc.Version
			csvDesc.Kind = uDesc.Kind
			du.Owned[i] = csvDesc
		}
	}
	csv.Spec.CustomResourceDefinitions.Owned = du.Owned
	return nil
}
<<<<<<< HEAD
=======

type ALMExamplesUpdate struct {
	crs []string
}

func (store *updaterStore) AddCR(yamlDoc []byte) error {
	if len(yamlDoc) == 0 {
		return nil
	}
	crBytes, err := yaml.YAMLToJSON(yamlDoc)
	if err != nil {
		return err
	}
	store.almExamples.crs = append(store.almExamples.crs, string(crBytes))
	return nil
}

func (u *ALMExamplesUpdate) Apply(csv *olmapiv1alpha1.ClusterServiceVersion) error {
	if len(u.crs) == 0 {
		return nil
	}
	if csv.GetAnnotations() == nil {
		csv.SetAnnotations(make(map[string]string))
	}
	sb := &strings.Builder{}
	sb.WriteString(`[`)
	for i, example := range u.crs {
		sb.WriteString(example)
		if i < len(u.crs)-1 {
			sb.WriteString(`,`)
		}
	}
	sb.WriteString(`]`)

	csv.GetAnnotations()["alm-examples"] = sb.String()
	return nil
}
>>>>>>> v0.0.4
