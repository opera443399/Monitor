package main

//Version app version
const Version = "0.6.0-dev"

//GitSHA We want to replace this variable at build time with "-ldflags -X main.GitSHA=xxx", where const is not supported.
var GitSHA = ""
