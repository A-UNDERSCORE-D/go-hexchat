package hexchat

// #include "hexchat-plugin.h"
import "C"

// Hook Priorities
const (
	PriorityHighest = C.HEXCHAT_PRI_HIGHEST
	PriorityHigh    = C.HEXCHAT_PRI_HIGH
	PriorityNorm    = C.HEXCHAT_PRI_NORM
	PriorityLow     = C.HEXCHAT_PRI_LOW
	PriorityLowest  = C.HEXCHAT_PRI_LOWEST
)

// Behaviour Modifiers
const (
	EatNone   = C.HEXCHAT_EAT_NONE
	EatPlugin = C.HEXCHAT_EAT_PLUGIN
	EatAll    = C.HEXCHAT_EAT_ALL
)
