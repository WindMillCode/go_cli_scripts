package utils

import (
	"os"
	"path/filepath"
)

func CDToLocation(location string, opts ...interface{}) {
	createIfNotExist := false

	if len(opts) > 0 {
		// Expecting the first option to be a boolean
		if create, ok := opts[0].(bool); ok {
			createIfNotExist = create
		}
	}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		if createIfNotExist {
			// Create the directory if it does not exist
			if mkErr := os.MkdirAll(location, os.ModePerm); mkErr != nil {
				panic(mkErr) // Panic if unable to create the directory
			}
		} else {
			panic("Destination does not exist")
		}
	}

	// Change directory
	if err := os.Chdir(location); err != nil {
		panic(err) // Panic if unable to change directory
	}
}
func CDToWorkspaceRoot() {
	CDToLocation(filepath.Join("..", "..", ".."))
}
func CDToAngularApp() {
	CDToLocation(filepath.Join("apps", "frontend", "AngularApp"))
}

func CDToFirebaseApp() {
	CDToLocation(filepath.Join("apps", "cloud", "FirebaseApp"))
}

func CDToShopifyApp() {
	CDToLocation(filepath.Join("apps", "cloud", "ShopifyApp"))
}

func CDToFlaskApp() {
	CDToLocation(filepath.Join("apps", "backend", "FlaskApp"))
}

func CDToTestNGApp() {
	CDToLocation(filepath.Join("apps", "testing", "testng"))
}

func CDToFlutterApp() {
	CDToLocation(filepath.Join("apps", "mobile", "FlutterApp"))
}





