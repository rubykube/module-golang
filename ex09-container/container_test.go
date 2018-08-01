package main

import (
	"io"
	"strings"
	"testing"

	"github.com/rendon/testcli"
)

//If the test doesn't compile on your machine try to run 'go get github.com/rendon/testcli' first

func TestBuild(t *testing.T) {
	c := testcli.Command("go", "build", "-o", "container")
	c.Run()
	if !c.Success() {
		t.Fatalf("[ERROR]: The container was expected to build!")
	}
}

func TestRun(t *testing.T) {
	c := testcli.Command("sudo", "./container", "run", "/bin/bash")
	c.Run()
	if !c.Success() {
		t.Fatalf("[ERROR]: I couldn't run your container!")
	}
}

func TestNoArgs(t *testing.T) {
	c := testcli.Command("sudo ./container")
	d := testcli.Command("sudo ./container run")
	c.Run()
	if c.Success() {
		t.Fatalf("[ERROR]: Your container shouldn't be running without 'run' command")
	}
	d.Run()
	if d.Success() {
		t.Fatalf("[ERROR]: Container should return an error if you don't specify what to run from it")
	}
}

func TestProcs(t *testing.T) {
	var r io.Reader
	r = strings.NewReader("ls /proc")
	c := testcli.Command("sudo", "./container", "run", "/bin/bash")
	c.SetStdin(r)
	c.Run()
	if !c.Success() {
		t.Fatalf("[ERROR]: Something went wrong when doing ls /proc inside your container")
	}
	d := testcli.Command("ls", "/proc")
	d.Run()
	if !d.Success() {
		t.Fatalf("[ERROR]: Something went wrong while doing ls /proc outside the container")
	}
	if strings.Compare(c.Stdout(), d.Stdout()) == 0 {
		t.Fatalf("[ERROR]: Your container shares processes with host machine!")
	}
}
