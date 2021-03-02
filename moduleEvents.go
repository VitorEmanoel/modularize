package modularize

var OnEnableEvent = "module/OnEnable"
type OnEnableExecutor func ()

var OnDisableEvent = "module/OnDisable"
type OnDisableExecutor func ()
