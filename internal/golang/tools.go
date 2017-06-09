package golang

import "github.com/maprost/gox/internal/shell"

func GoDep() error {
	// TODO: check if vendor or GoDep folder are available -> if yes try godep update ./... else godep save ./...

	_, err := shell.Command("godep", "save", "./...")
	return err
}
