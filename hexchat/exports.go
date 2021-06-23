package hexchat

// #include "hexchat-plugin.h"
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

var (
	name          string
	description   string
	version       string
	pluginCreated bool
	onLoad        func() int

	pluginHandle *C.hexchat_plugin
)

func RegisterPlugin(plugin_name, plugin_description, plugin_version string, onload func() int) {
	if pluginCreated {
		panic("Attempt to create plugin twice")
	}
	name = plugin_name
	description = plugin_description
	version = plugin_version
	pluginCreated = true
	onLoad = onload
}

//export hexchat_plugin_init
func hexchat_plugin_init(handle *C.hexchat_plugin, pluginName **C.char, pluginDesc **C.char, pluginVersion **C.char, args C.char) int {
	if !pluginCreated {
		panic("No plugin has been created!")
	}
	*pluginName = C.CString(name)
	*pluginDesc = C.CString(description)
	*pluginVersion = C.CString(version)

	pluginHandle = handle
	return onLoad()
}

type HookFunc func(word []string, word_eol []string, userdata string) int

var (
	callbackMutex      = sync.Mutex{}
	allNormalCallbacks = func() map[int32]HookFunc {
		return map[int32]HookFunc{}
	}()
)
var currentIdx int32 = 0

func getStrSlice(cData **C.char) []string {
	// The max length these can be is 32 according to docs
	result := (*[32]*C.char)(unsafe.Pointer(cData))
	// This copies but I have no choice really
	out := make([]string, 32)
	for i := 0; i < 32; i++ {
		out[i] = C.GoString(result[i])
	}

	return out
}

func storeGoCallback(cb HookFunc) int32 {
	callbackMutex.Lock()
	defer callbackMutex.Unlock()
	idx := currentIdx
	allNormalCallbacks[idx] = cb
	currentIdx++
	return idx
}

//export __call_go_callback
func __call_go_callback(index C.int, word **C.char, wordeol **C.char) int {
	fmt.Printf("i: %d, w: %v, w_e: %v", index, word, wordeol)
	// TODO: mutex
	return allNormalCallbacks[int32(index)](getStrSlice(word), getStrSlice(wordeol), "")
}
