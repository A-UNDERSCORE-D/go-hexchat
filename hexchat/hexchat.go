package hexchat

/*
#include "hexchat-plugin.h"
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
int __call_go_callback(int id, char *word[], char *word_eol[]);
int __callGoWordWordEOLIdx(int id, char *word[], char *word_eol[]) {
	return __call_go_callback(id, word, word_eol);
}

// This is the actual wrapper we use to call functions. Unfortunately this means that userdata is broken unless
// I do something clever in go land
int __wrapper(char *word[], char *word_eol[], void *user_data) {
	return __callGoWordWordEOLIdx((int) (intptr_t) user_data, word, word_eol);
}

void __PrintWrapper(hexchat_plugin *plugin_handle, char *msg) {	plugin_handle->hexchat_print(plugin_handle, msg); }
hexchat_hook* __hook_command_wrapper(hexchat_plugin *ph, const char *name, int pri, int callback_id) {
	// void on x86_64 is 64 bits, we store our callback IDs in an int32. Thus we get a warning because the sizes differ
	// To solve this, we can cast to uintptr FIRST. which is what we do here. This should be cross platform and
	// work fine on x86. But also here be dragons. I promise nothing.
	void* callback_id_ptr = (void *) (uintptr_t) callback_id;
	ph->hexchat_hook_command(ph, name, pri, __wrapper, "", callback_id_ptr);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func CheckPluginExists() {
	if !pluginCreated {
		panic("Plugin has not been created")
	}
}

func Print(msg string) {
	CheckPluginExists()
	cs := C.CString(msg)
	defer C.free(unsafe.Pointer(cs))
	C.__PrintWrapper(pluginHandle, cs)
}

func HookCommand(cmdName string, priority int, help string, callback HookFunc) {
	CheckPluginExists()
	fmt.Println(cmdName, priority, help, callback)
	fmt.Println(pluginHandle)
	csCmdName := C.CString(cmdName)
	callbackIDX := storeGoCallback(callback)
	callbackCIdx := C.int(callbackIDX)
	fmt.Printf("Go: %d, C: %#v\n", callbackIDX, callbackCIdx)
	C.__hook_command_wrapper(pluginHandle, csCmdName, C.int(priority), callbackCIdx)
}
