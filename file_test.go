package main

import (
	"testing"
)

func TestFileRead(*testing.T) {
	File f = File.Open("~/serverpedido.log", os.O_RDONLY, 0)
	
}