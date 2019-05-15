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

package controllermap

import (
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
<<<<<<< HEAD
	"k8s.io/apimachinery/pkg/types"
=======
>>>>>>> v0.0.4
	"sigs.k8s.io/controller-runtime/pkg/controller"
)

// ControllerMap - map of GVK to ControllerMapContents
type ControllerMap struct {
	mutex    sync.RWMutex
<<<<<<< HEAD
	internal map[schema.GroupVersionKind]*ControllerMapContents
}

// UIDMap - map of UID to namespaced name of owner
type UIDMap struct {
	mutex    sync.RWMutex
	internal map[types.UID]types.NamespacedName
=======
	internal map[schema.GroupVersionKind]*Contents
>>>>>>> v0.0.4
}

// WatchMap - map of GVK to interface. Determines if resource is being watched already
type WatchMap struct {
	mutex    sync.RWMutex
	internal map[schema.GroupVersionKind]interface{}
}

<<<<<<< HEAD
// ControllerMapContents- Contains internal data associated with each controller
type ControllerMapContents struct {
	Controller                  controller.Controller
	WatchDependentResources     bool
	WatchClusterScopedResources bool
	WatchMap                    *WatchMap
	UIDMap                      *UIDMap
=======
// Contents - Contains internal data associated with each controller
type Contents struct {
	Controller                  controller.Controller
	WatchDependentResources     bool
	WatchClusterScopedResources bool
	OwnerWatchMap               *WatchMap
	AnnotationWatchMap          *WatchMap
>>>>>>> v0.0.4
}

// NewControllerMap returns a new object that contains a mapping between GVK
// and ControllerMapContents object
func NewControllerMap() *ControllerMap {
	return &ControllerMap{
<<<<<<< HEAD
		internal: make(map[schema.GroupVersionKind]*ControllerMapContents),
=======
		internal: make(map[schema.GroupVersionKind]*Contents),
>>>>>>> v0.0.4
	}
}

// NewWatchMap - returns a new object that maps GVK to interface to determine
// if resource is being watched
func NewWatchMap() *WatchMap {
	return &WatchMap{
		internal: make(map[schema.GroupVersionKind]interface{}),
	}
}

<<<<<<< HEAD
// NewUIDMap - returns a new object that maps UID to namespaced name of owner
func NewUIDMap() *UIDMap {
	return &UIDMap{
		internal: make(map[types.UID]types.NamespacedName),
	}
}

// Get - Returns a ControllerMapContents given a GVK as the key. `ok`
// determines if the key exists
func (cm *ControllerMap) Get(key schema.GroupVersionKind) (value *ControllerMapContents, ok bool) {
=======
// Get - Returns a ControllerMapContents given a GVK as the key. `ok`
// determines if the key exists
func (cm *ControllerMap) Get(key schema.GroupVersionKind) (value *Contents, ok bool) {
>>>>>>> v0.0.4
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	value, ok = cm.internal[key]
	return value, ok
}

// Delete - Deletes associated GVK to controller mapping from the ControllerMap
func (cm *ControllerMap) Delete(key schema.GroupVersionKind) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.internal, key)
}

// Store - Adds a new GVK to controller mapping
<<<<<<< HEAD
func (cm *ControllerMap) Store(key schema.GroupVersionKind, value *ControllerMapContents) {
=======
func (cm *ControllerMap) Store(key schema.GroupVersionKind, value *Contents) {
>>>>>>> v0.0.4
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.internal[key] = value
}

// Get - Checks if GVK is already watched
func (wm *WatchMap) Get(key schema.GroupVersionKind) (value interface{}, ok bool) {
	wm.mutex.RLock()
	defer wm.mutex.RUnlock()
	value, ok = wm.internal[key]
	return value, ok
}

// Delete - Deletes associated watches for a specific GVK
func (wm *WatchMap) Delete(key schema.GroupVersionKind) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()
	delete(wm.internal, key)
}

// Store - Adds a new GVK to be watched
func (wm *WatchMap) Store(key schema.GroupVersionKind) {
	wm.mutex.Lock()
	defer wm.mutex.Unlock()
	wm.internal[key] = nil
}
<<<<<<< HEAD

// Get - Returns a NamespacedName of the owner given a UID
func (um *UIDMap) Get(key types.UID) (value types.NamespacedName, ok bool) {
	um.mutex.RLock()
	defer um.mutex.RUnlock()
	value, ok = um.internal[key]
	return value, ok
}

// Delete - Deletes associated UID to NamespacedName mapping
func (um *UIDMap) Delete(key types.UID) {
	um.mutex.Lock()
	defer um.mutex.Unlock()
	delete(um.internal, key)
}

// Store - Adds a new UID to NamespacedName mapping
func (um *UIDMap) Store(key types.UID, value types.NamespacedName) {
	um.mutex.Lock()
	defer um.mutex.Unlock()
	um.internal[key] = value
}
=======
>>>>>>> v0.0.4
