//go:build integration
// +build integration

package jcosmos

import "testing"

// emulator must be started with the following details

const (
	key    = ""
	host   = ""
	dbName = ""
)

var (
	colls = [5]string{"col1", "col2", "col3", "col4", "col5"}
	users = [5]string{"usr1", "user2", "user3", "user4", "user5"}
)
