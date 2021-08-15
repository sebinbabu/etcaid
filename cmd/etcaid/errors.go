package main

import "errors"

var errUninitialized = errors.New("etcaid isn't initialized yet. Use 'etcaid init' to initialize.")
