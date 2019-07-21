package dao

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init()

	code := m.Run()
	os.Exit(code)
}