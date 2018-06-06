package main

import (
	"log"
	"sync"

	"github.com/google/btree"
	cache "github.com/patrickmn/go-cache"
)

var MicroserviceInstanceCache = cache.New(0, 0)

func main() {
	key := "RESTServer"
	microServiceInstances := []*MicroServiceInstance{
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.1", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.2", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.2", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.3", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.4", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.5", "project": "nino"}},
		{Metadata: map[string]string{"version": "0.0.5", "project": "nino"}},
	}

	refreshCache(key, microServiceInstances)

	tags := map[string]string{"version": "0.0.2"}
	v, ok := findMicroServiceInstances(key, tags)
	if ok {
		log.Printf("GetKey %s from cache %v", key, v)
	}
}

func refreshCache(key string, microServiceInstances []*MicroServiceInstance) {
	MicroserviceInstanceCache.Set(key, microServiceInstances, 0)
	for _, m := range microServiceInstances {
		Indexer.Put(m, "version")
	}
}

func findMicroServiceInstances(key string, tags map[string]string) ([]*MicroServiceInstance, bool) {
	v, ok := MicroserviceInstanceCache.Get(key)git sta
	if !ok {
		return nil, false
	}
	log.Printf("GetKey %s from cache %v", key, v)
	// microServiceInstances, ok := v.([]*MicroServiceInstance)
	// if !ok {
	// 	return nil, false
	// }
	log.Printf("GetIndex %v", Indexer.Get(tags["version"]))
	return Indexer.Get(tags["version"]), true
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
	tree *btree.BTree
}

type index interface {
	Put(m *MicroServiceInstance, label string)
	Get(version string) []*MicroServiceInstance
}

var Indexer = newTreeIndex()

func newTreeIndex() index { return &treeIndex{tree: btree.New(32)} }

type keyIndex struct {
	label string
	key   string

	microIns []*MicroServiceInstance
}

func (a *keyIndex) Less(b btree.Item) bool       { return a.key < b.(*keyIndex).key }
func (k *keyIndex) get() []*MicroServiceInstance { return k.microIns }
func (k *keyIndex) put(m *MicroServiceInstance) {
	if k.microIns == nil {
		k.microIns = make([]*MicroServiceInstance, 0)
	}
	k.microIns = append(k.microIns, m)
}

func (ti *treeIndex) Put(m *MicroServiceInstance, label string) {
	log.Printf("key(%v) label(%v)", m.Metadata[label], label)
	keyi := &keyIndex{
		key:   m.Metadata[label],
		label: label}

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
