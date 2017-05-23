package gxcfg

import "strings"

type DatabaseAccess int

const (
	DatabaseAccessPort = DatabaseAccess(1)
	DatabaseAccessLink = DatabaseAccess(2)
)

const (
	DatabaseAccessPortString = "port"
	DatabaseAccessLinkString = "link"
)

func (d DatabaseAccess) String() string {
	return databaseAccessIDStringMap[d]
}

var databaseAccessStringIDMap = map[string]DatabaseAccess{
	strings.ToLower(DatabaseAccessPortString): DatabaseAccessPort,
	strings.ToLower(DatabaseAccessLinkString): DatabaseAccessLink,
}

func DatabaseAccessStringToID(productType string) (DatabaseAccess, bool) {
	res, ok := databaseAccessStringIDMap[strings.ToLower(productType)]
	return res, ok
}

var databaseAccessIDStringMap = map[DatabaseAccess]string{
	DatabaseAccessPort: DatabaseAccessPortString,
	DatabaseAccessLink: DatabaseAccessLinkString,
}
