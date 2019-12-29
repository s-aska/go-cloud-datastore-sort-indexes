package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-yaml/yaml"
)

// Index Cloud Datastore Index
type Index struct {
	Indexes []struct {
		Kind       string `yaml:"kind"`
		Properties []struct {
			Name      string `yaml:"name"`
			Direction string `yaml:"direction,omitempty"`
		} `yaml:"properties"`
	} `yaml:"indexes"`
}

func main() {
	err := sortIndexes()
	if err != nil {
		panic(err)
	}
}

func sortIndexes() (err error) {
	var in *os.File
	var out *os.File

	var input string
	var output string
	flag.StringVar(&input, "i", "", "input index.yaml path")
	flag.StringVar(&output, "o", "", "output index.yaml path")
	flag.Parse()

	if input != "" {
		in, err = os.Open(input)
		if err != nil {
			return
		}
	} else {
		in = os.Stdin
	}

	if output != "" {
		out, err = os.Create(output)
		if err != nil {
			return
		}
	} else {
		out = os.Stdout
	}

	index := &Index{}
	err = yaml.NewDecoder(in).Decode(index)
	if err != nil {
		return
	}
	sort.Slice(index.Indexes, func(i, j int) bool {
		return strings.Compare(fmt.Sprintf("%+v", index.Indexes[i]), fmt.Sprintf("%+v", index.Indexes[j])) < 0
	})
	err = yaml.NewEncoder(out).Encode(index)
	if err != nil {
		return
	}
	return out.Close()
}
