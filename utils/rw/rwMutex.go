package rw

import "sync"

type Unlockable interface {
	Unlock()
}

type Mutex struct {
	mLockable *sync.RWMutex
}

func NewMutex() Mutex {
	return Mutex{mLockable: &sync.RWMutex{}}
}

func (this *Mutex) RLock() Unlockable {
	this.mLockable.RLock()
	return &rLock{mMutex: this}
}

func (this *Mutex) WLock() Unlockable {
	this.mLockable.Lock()
	return &wLock{mMutex: this}
}

type rLock struct {
	mMutex *Mutex
}

func (this *rLock) Unlock() {
	this.mMutex.mLockable.RUnlock()
}

type wLock struct {
	mMutex *Mutex
}

func (this *wLock) Unlock() {
	this.mMutex.mLockable.Unlock()
}
