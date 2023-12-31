package unigo

import (
	"errors"
	"fmt"
)

type DisjointedSetsDataStructure[T comparable] interface {
	Union(item, item2 T) error
	Find(item T) (T, error)
	Connected(item, item2 T) (bool, error)
	MakeSet(item T) (T, error)
}

type Unigo[T comparable] struct {
	id map[T]T
	sz map[T]int
}

/*
*	MakeSet initializes the map if necessary; Checks to set if set with
* 	this key has already been initialized; Initializes set with
*	the provided key.
 */
func (iuf *Unigo[T]) MakeSet(item T) error {
	if !iuf.mapIntialized() {
		iuf.initializeMap()
	}

	_, exists := iuf.id[item]
	if exists {
		return fmt.Errorf("set %v already exists", item)
	}

	iuf.id[item] = item
	iuf.sz[item] = 1

	return nil
}

/*
*	Find initializes map if necessary; Uses the supplied key to check
*	if that key exists in the map; Find the set identifier for the
*	supplied key and returns it.
 */
func (iuf *Unigo[T]) Find(item T) (T, error) {
	var setId T
	var compBuffer []T

	if !iuf.mapIntialized() {
		iuf.initializeMap()
	}

	setId, exists := iuf.id[item]

	if !exists {
		return setId, fmt.Errorf("value %v does not exist in any set", item)
	}

	for iuf.id[setId] != setId {
		compBuffer = append(compBuffer, setId)
		setId = iuf.id[setId]
	}

	for _, val := range compBuffer {
		iuf.id[val] = setId
	}

	return setId, nil
}

/*
*	Find initializes map if necessary; Finds the set identifier for both
*	supplied keys; Returns errors if any key does not belong to a set;
*	Changes the parent of item2 to be the parent of item.
 */
func (iuf *Unigo[T]) Union(item, item2 T) (T, error) {
	if !iuf.mapIntialized() {
		iuf.initializeMap()
	}

	setId, findErr := iuf.Find(item)
	if findErr != nil {
		return setId, findErr
	}

	setId2, findErr2 := iuf.Find(item2)
	if findErr2 != nil {
		return setId2, findErr2
	}

	if setId == setId2 {
		return setId, nil
	}

	if iuf.sz[setId] < iuf.sz[setId2] {
		iuf.sz[setId2] += iuf.sz[setId]
		iuf.id[setId] = setId2
	} else {
		iuf.sz[setId] += iuf.sz[setId2]
		iuf.id[setId2] = setId
	}

	return setId, nil
}

/*
*	Connected initializes map if necessary; Finds the set identifier for
*	both supplied keys; Returns errors if key does not belong to a set;
*	Returns a boolean informing if both keys belong to same set;
 */
func (iuf *Unigo[T]) Connected(item, item2 T) (bool, error) {
	if !iuf.mapIntialized() {
		iuf.initializeMap()
	}

	setId, findErr := iuf.Find(item)
	if findErr != nil {
		return false, findErr
	}

	setId2, findErr2 := iuf.Find(item2)
	if findErr2 != nil {
		return false, findErr2
	}

	return setId == setId2, nil
}

/*
*	Checks if map was initialized.
 */
func (iuf Unigo[T]) mapIntialized() bool {
	return iuf.id != nil && iuf.sz != nil
}

/*
*	Initializes map.
 */
func (iuf *Unigo[T]) initializeMap() error {
	if iuf.mapIntialized() {
		return errors.New("map is already initialized")
	}

	iuf.id = make(map[T]T)
	iuf.sz = make(map[T]int)

	return nil
}
