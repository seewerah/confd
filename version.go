package main

var Version = "0.17.99"

// We want to replace this variable at build time with "-ldflags -X main.GitSHA=xxx", where const is not supported.
var GitSHA = ""
