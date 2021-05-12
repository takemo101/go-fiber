package contract

// Middleware interface
type Middleware interface {
	Setup()
}

// Route interface
type Route interface {
	Setup()
}

// Command interface
type Command interface {
	Setup()
}
