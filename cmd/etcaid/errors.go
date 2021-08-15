package main

import "errors"

// errUninitialized signifies that etcaid hasn't been initialized yet.
var errUninitialized = errors.New("etcaid isn't initialized yet. Use `etcaid init` to initialize.")
