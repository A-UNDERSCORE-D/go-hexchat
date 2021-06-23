package main

import (
	"fmt"

	"awesome-dragon.science/go/go-hexchat/hexchat"
)

func init() {
	fmt.Println("Does init work?")
	hexchat.RegisterPlugin("test", "tests things", "0.0.1", func() int {
		hexchat.HookCommand("adtest", hexchat.PriorityNorm, "", func(word, word_eol []string, userdata string) int {
			hexchat.Print("STUFF!")
			return hexchat.EatNone
		})

		return 1
	})
}

func main() {
	panic("Main called")
}

// //export hexchat_plugin_init
// func hexchat_plugin_init(handle *C.hexchat_plugin, pluginName **C.char, pluginDesc **C.char, pluginVersion **C.char, args C.char) int {
// 	pluginHandle = handle

// 	*pluginName = C.CString(Name)
// 	*pluginDesc = C.CString(Desc)
// 	*pluginVersion = C.CString(Ver)

// 	fmt.Println("Loaded!")
// 	hexchatPrint(pluginHandle, "loaded!")
// 	return 1
// }
