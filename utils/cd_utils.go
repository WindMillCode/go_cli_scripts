package utils

import (
	"os"
	"path/filepath"
)

func CDToLocation(location string) {
	if err := os.Chdir(location); err != nil {
		panic(err)
	}
}
func CDToWorkspaceRoot() {
	CDToLocation(filepath.Join("..", "..", "..", ".."))
}
func CDToAngularApp() {
	CDToLocation(filepath.Join("apps", "frontend", "AngularApp"))
}

func CDToFirebaseApp() {
	CDToLocation(filepath.Join("apps", "cloud", "FirebaseApp"))
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
