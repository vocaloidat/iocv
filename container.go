package iocv

// iocv.Controller().Registry
var Controller Container = &MapContainer{
	name:   "controller",
	storge: make(map[string]Object),
}

// iocv.Api().Registry
var Api Container = &MapContainer{
	name:   "api",
	storge: make(map[string]Object),
}
