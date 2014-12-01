package iup

/*
Package iup provides a wrapper for the IUP library, a GUI library written in C.

The API is as close to the original C API as possible. For example boolean values
are usually passed as integers where 0 means false and 1 means true. Attributes
for controls are passed as strings.

Setting callbacks in Go is done with SetCallback and the parameter must be a function
returning int. The parameters are Go's equivalents to the C types. HOWEVER the first
callback type is usually the Ihandle* that created the event but this is omitted
in this wrapper. The reason is that most of the time the handle is not needed and
if you do need it you could easily pass it as a closure.

Callbacks are treated specially. Since you cannot simply pass a go function to C
as a callback, there is a workaround to make IUP's SetCallback work in Go. There
is a fixed number of C functions, a pool of functions. Each time you call SetCallback
in an IUP control, one function from the pool is assigned as the C callback. It
does nothing more than call the Go function that you passed to SetCallback. This
mechanism gives maximum efficiency but has the drawback that you only have a finite
number of callbacks that can be assigned. If you need more, you need to re-generate
the "callbacks.go" file with a higher number of callbacks.
You will know when you exceeded the callback limit once you call SetCallback and
the program stops with a panic informing you about it. You can then re-generate
the callback wrappers but the current number should be big enough for small to medium
sized GUIs.
*/
