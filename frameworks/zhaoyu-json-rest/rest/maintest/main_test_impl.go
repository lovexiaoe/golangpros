package main

import (
	"mygolang/zhaoyu-json-rest/rest/trie"
)

func main() {
	trie := trie.New()
	//	trie.AddRoute("GET", "/r/:id/property.*format", "property_format")
	//	trie.AddRoute("GET", "/user/#username/property", "user_property")
	trie.AddRoute("GET", "/user/", "property_format")
	trie.AddRoute("GET", "/a/", "user_property")

	trie.PrintDebug()

	trie.Compress()

	trie.PrintDebug()
}
