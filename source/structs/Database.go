package structs

import "pacman-backup/console"

type Database struct {
	Packages []Package `json:"packages"`
}

func NewDatabase() Database {

	var database Database

	database.Packages = make([]Package, 0)

	return database

}

func (database *Database) AddPackage(pkg Package) {

	console.Inspect(pkg)

}
