package testsample

import (
	"sync"
	"sync/atomic"
	"testing"
)

type manager struct {
	sync.RWMutex
	agents int
}

func BenchmarkManagerLock(b *testing.B) {
	m := new(manager)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			m.agents = 100
			m.Unlock()
		}
	})
}

func BenchmarkManagerRLock(b *testing.B) {
	m := manager{agents: 100}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.RLock()
			_ = m.agents
			m.RUnlock()
		}
	})
}

func BenchmarkManagerAtomicValueStore(b *testing.B) {
	var managerVal atomic.Value
	m := manager{agents: 100}
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			managerVal.Store(m)
		}
	})
}

func BenchmarkManagerAtomicValueLoad(b *testing.B) {
	var managerVal atomic.Value
	managerVal.Store(&manager{agents: 100})
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = managerVal.Load().(*manager)
		}
	})
}
