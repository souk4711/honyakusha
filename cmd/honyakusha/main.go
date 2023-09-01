package main

import "github.com/souk4711/honyakusha/internal/cmd"

var (
	GitCommit = ""
	BuiltTime = ""
)

func main() {
	cmd.Execute(GitCommit, BuiltTime)
}
