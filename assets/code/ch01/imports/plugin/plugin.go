// Package plugin exists only to show a blank import: its init function runs
// even though the importer never calls anything in it.
package plugin

import "fmt"

func init() {
	fmt.Println("plugin: registered itself via init()")
}
