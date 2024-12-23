package dfs

/*
#cgo LDFLAGS: -ldaos -ldaos_fs
#include <daos_fs.h>
*/
import "C"
import "unsafe"

// DFS represents a DFS object
type DFS struct {
	dfs *C.dfs_t
}

// Mount mounts a DFS namespace
func Mount(pool, cont string) (*DFS, error) {
	var dfs DFS
	cPool := C.CString(pool)
	defer C.free(unsafe.Pointer(cPool))
	cCont := C.CString(cont)
	defer C.free(unsafe.Pointer(cCont))

	rc := C.dfs_mount(cPool, cCont, &dfs.dfs)
	if rc != 0 {
		return nil, fmt.Errorf("failed to mount DFS: %d", rc)
	}
	return &dfs, nil
}

// Unmount unmounts a DFS namespace
func (d *DFS) Unmount() error {
	rc := C.dfs_umount(d.dfs)
	if rc != 0 {
		return fmt.Errorf("failed to unmount DFS: %d", rc)
	}
	return nil
}

// HelloWorld prints "Hello, World!" to the console
func HelloWorld() {
	fmt.Println("Hello, World!")
}