package main

import (
	"log"

	"github.com/f110/release-radar/pkg/producer"
)

func main() {
	p := producer.NewGithubRelease("memcached", "memcached")
	log.Print(p.Produce())
}
