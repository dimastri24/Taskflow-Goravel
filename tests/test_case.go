package tests

import (
	"github.com/goravel/framework/testing"

	"taskflow/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
