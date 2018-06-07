package main

import (
	"log"
	"sync"

	"github.com/google/btree"
	cache "github.com/patrickmn/go-cache"
)

var MicroserviceInstanceCache = cache.New(0, 0)
var Indexers = make(map[string]map[string]index, 0)

func main() {
	key := "RESTServer"
	microServiceInstances := []*MicroServiceInstance{
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.2", "project": "eesly"}},
		{Metadata: map[string]string{"version": "0.0.2", "project": "villa"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "villa"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.4", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.5", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.5", "project": "eesly"}},
	}

	tags := map[string]string{"version": "0.0.3", "project": "villa"}
	initCache(key, microServiceInstances, tags)

	v, ok := findMicroServiceInstances(key, tags)
	if ok {
		log.Printf("GetKey %s from cache %v", key, v)
	}
}

func initCache(key string, microServiceInstances []*MicroServiceInstance, tags map[string]string) {
	MicroserviceInstanceCache.Set(key, microServiceInstances, 0)
	if _, ok := Indexers[key]; !ok {
		Indexers[key] = make(map[string]index, 0)
	}
	for ind := range tags {
		Indexers[key][ind] = newTreeIndex(ind)
	}
	for _, m := range microServiceInstances {
		for ind := range Indexers[key] {
			Indexers[key][ind].Put(m)
		}
	}
}

func findMicroServiceInstances(key string, tags map[string]string) ([]*MicroServiceInstance, bool) {
	v, ok := MicroserviceInstanceCache.Get(key)
	if !ok {
		return nil, false
	}
	log.Printf("GetKey %s from cache %v", key, v)
	// microServiceInstances, ok := v.([]*MicroServiceInstance)
	// if !ok {
	// 	return nil, false
	// }
	ret := make([][]*MicroServiceInstance, len(tags))
	i := 0
	indexers := Indexers[key]
	for ind, value := range tags {
		ret[i] = indexers[ind].Get(value)
		i++
		log.Printf("GetIndex %v for %v=%v", indexers[ind].Get(value), ind, value)

	}
	return util.MultiInterSection(ret), true
}

type MicroServiceInstance struct {
	Version         string
	InstanceID      string
	HostName        string
	ServiceID       string
	DefaultProtocol string
	DefaultEndpoint string
	Status          string
	EndpointsMap    map[string]string
	Metadata        map[string]string
}

type treeIndex struct {
	sync.RWMutex
	tree  *btree.BTree
	label string
}

type index interface {
	Put(m *MicroServiceInstance)
	Get(version string) []*MicroServiceInstance
}

func newTreeIndex(label string) index { return &treeIndex{label: label, tree: btree.New(32)} }

func (ti *treeIndex) Put(m *MicroServiceInstance) {
	log.Printf("key(%v) label(%v)", m.Metadata[ti.label], ti.label)
	keyi := &keyIndex{
		key:   m.Metadata[ti.label],
		label: ti.label}

	ti.RLock()
	defer ti.RUnlock()
	item := ti.tree.Get(keyi)
	if item == nil {
		keyi.put(m)
		ti.tree.ReplaceOrInsert(keyi)
		return
	}

	okeyi := item.(*keyIndex)
	okeyi.put(m)
}

func (ti *treeIndex) Get(version string) []*MicroServiceInstance {
	keyi := &keyIndex{key: version}

	ti.RLock()
	defer ti.RUnlock()
	item := ti.tree.Get(keyi)
	if item == nil {
		return nil
	}
	keyi = item.(*keyIndex)
	return keyi.get()
}

type keyIndex struct {
	label string
	key   string

	microIns []*MicroServiceInstance
}

func (k *keyIndex) Less(b btree.Item) bool       { return k.key < b.(*keyIndex).key }
func (k *keyIndex) get() []*MicroServiceInstance { return k.microIns }
func (k *keyIndex) put(m *MicroServiceInstance) {
	if k.microIns == nil {
		k.microIns = make([]*MicroServiceInstance, 0)
	}
	k.microIns = append(k.microIns, m)
}
