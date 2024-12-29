package dfs

/*
#cgo LDFLAGS: -ldaos -ldaos_fs
#include <daos_.h>
#include <daos_fs.h>
*/
import "C"
import "unsafe"

// DFS represents a DFS object
type DFS struct {
	dfs *C.dfs_t
}


// Connect establishes a connection to a DFS (Distributed File System) instance.
//
// Parameters:
//   - pool: The pool name to connect to.
//   - sys: The system name to connect to.
//   - cont: The container name to connect to.
//   - flags: Connection flags.
//   - attr: Pointer to DFS attributes.
//
// Returns:
//   - error: An error if the connection fails, otherwise nil.
func (d *DFS) Connect(pool, sys, cont string, flags int, attr *C.dfs_attr_t) error {
	cPool := C.CString(pool)
	defer C.free(unsafe.Pointer(cPool))
	cSys := C.CString(sys)
	defer C.free(unsafe.Pointer(cSys))
	cCont := C.CString(cont)
	defer C.free(unsafe.Pointer(cCont))

	rc := C.dfs_connect(cPool, cSys, cCont, C.int(flags), attr, &d.dfs)
	if rc != 0 {
		return fmt.Errorf("failed to connect DFS: %d", rc)
	}
	return nil
}


// Disconnect disconnects the DFS instance. 
// It calls the underlying C function to perform the disconnection.
// If the disconnection fails, it returns an error with the failure code.
func (d *DFS) Disconnect() error {
	rc := C.dfs_disconnect(d.dfs)
	if rc != 0 {
		return fmt.Errorf("failed to disconnect DFS: %d", rc)
	}
	return nil
}


// Open opens a DFS object with the specified parameters.
//
// Parameters:
// - parent: A pointer to the parent DFS object.
// - name: The name of the DFS object to open.
// - mode: The mode in which to open the DFS object.
// - flags: Flags to control the behavior of the open operation.
// - cid: The object class ID for the DFS object.
// - chunkSize: The chunk size for the DFS object.
// - value: A string value associated with the DFS object.
//
// Returns:
// - A pointer to the opened DFS object.
// - An error if the operation fails.
func (d *DFS) Open(parent *C.dfs_obj_t, name string, mode int, flags int, cid C.daos_oclass_id_t, chunkSize C.daos_size_t, value string) (*C.dfs_obj_t, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	var obj *C.dfs_obj_t
	rc := C.dfs_open(d.dfs, parent, cName, C.mode_t(mode), C.int(flags), cid, chunkSize, cValue, &obj)
	if rc != 0 {
		return nil, fmt.Errorf("failed to open DFS object: %d", rc)
	}
	return obj, nil
}


// Read reads data from a DFS object into a scatter-gather list.
//
// Parameters:
//   obj - Pointer to the DFS object to read from.
//   sgl - Pointer to the scatter-gather list where the read data will be stored.
//   off - Offset within the DFS object to start reading from.
//   readSize - Pointer to the size of the data to be read.
//   ev - Pointer to the event structure for asynchronous operations.
//
// Returns:
//   error - An error if the read operation fails, otherwise nil.
func (d *DFS) Read(obj *C.dfs_obj_t, sgl *C.d_sg_list_t, off C.daos_off_t, readSize *C.daos_size_t, ev *C.daos_event_t) error {
	rc := C.dfs_read(d.dfs, obj, sgl, off, readSize, ev)
	if rc != 0 {
		return fmt.Errorf("failed to read from DFS object: %d", rc)
	}
	return nil
}


// Release releases a DFS object.
//
// Parameters:
//   obj - A pointer to the DFS object to be released.
//
// Returns:
//   An error if the release operation fails, otherwise nil.
func (d *DFS) Release(obj *C.dfs_obj_t) error {
	rc := C.dfs_release(obj)
	if rc != 0 {
		return fmt.Errorf("failed to release DFS object: %d", rc)
	}
	return nil
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