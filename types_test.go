package adoptium

import (
	"fmt"
	"log"
	"testing"
)

func TestTypesArchitectures(t *testing.T) {
	r, err := TypesArchitectures()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}

func TestTypesOperatingSystems(t *testing.T) {
	r, err := TypesOperatingSystems()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
}

func ExampleTypesArchitectures() {
	r, err := TypesArchitectures()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// Output: [x64 x32 ppc64 ppc64le s390x aarch64 arm sparcv9 riscv64]
}

func ExampleTypesOperatingSystems() {
	r, err := TypesOperatingSystems()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// Output: [linux windows mac solaris aix alpine-linux]
}
