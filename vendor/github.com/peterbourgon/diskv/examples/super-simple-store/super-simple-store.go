package main

import (
	"fmt"

<<<<<<< HEAD
	"github.com/peterbourgon/diskv"
=======
	"github.com/peterbourgon/diskv/v3"
>>>>>>> v0.0.4
)

func main() {
	d := diskv.New(diskv.Options{
		BasePath:     "my-diskv-data-directory",
<<<<<<< HEAD
		Transform:    func(s string) []string { return []string{} },
=======
>>>>>>> v0.0.4
		CacheSizeMax: 1024 * 1024, // 1MB
	})

	key := "alpha"
	if err := d.Write(key, []byte{'1', '2', '3'}); err != nil {
		panic(err)
	}

	value, err := d.Read(key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", value)

	if err := d.Erase(key); err != nil {
		panic(err)
	}
}
