package main

import (
	"fmt"
	"github.com/blang/semver/v4"
	"github.com/hashicorp/go-version"
)

func main() {
	v1, _ := version.NewVersion("1.23.7")
	v2, _ := version.NewVersion("1.24.0-alpha.0")
	v3, _ := version.NewVersion("1.25.0")
	if v1.LessThan(v2) {
		fmt.Println("v1 less than v2")
	}
	if v2.LessThan(v3) {
		fmt.Println("v2 less than v3")
	}

	sv1, _ := semver.Make("1.0.0-beta")
	sv2, _ := semver.Make("2.0.0-beta")
	result := sv1.Compare(sv2)
	fmt.Printf("result=%d", result)

}
