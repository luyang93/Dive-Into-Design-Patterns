package main

import "fmt"

// Object represents the object that will be reused
type Object struct{}

// ObjectPool represents the object pool that stores the reusable objects
type ObjectPool struct {
	objects []*Object
}

// NewObjectPool creates a new object pool
func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		objects: make([]*Object, 0),
	}
}

// AcquireObject acquires an object from the object pool
func (op *ObjectPool) AcquireObject() *Object {
	if len(op.objects) > 0 {
		// If there are available objects in the pool, return one of them
		obj := op.objects[0]
		fmt.Println("Reusing existing object")
		op.objects = op.objects[1:]
		return obj
	}
	// If no objects are available, create a new one and return it
	obj := &Object{}
	fmt.Println("Creating new object")
	return obj
}

// ReleaseObject releases an object back to the object pool
func (op *ObjectPool) ReleaseObject(obj *Object) {
	// Reset the object's state if necessary
	// ...
	// Add the object back to the pool
	op.objects = append(op.objects, obj)
	fmt.Println("Releasing object")
}

func main() {
	pool := NewObjectPool()

	// Acquire objects from the pool
	obj1 := pool.AcquireObject()
	obj2 := pool.AcquireObject()

	// Release objects back to the pool
	pool.ReleaseObject(obj1)
	pool.ReleaseObject(obj2)

	// Acquire the released objects again
	obj3 := pool.AcquireObject()
	obj4 := pool.AcquireObject()
	obj5 := pool.AcquireObject()

	// Release objects back to the pool
	pool.ReleaseObject(obj3)
	pool.ReleaseObject(obj4)
	pool.ReleaseObject(obj5)
}
