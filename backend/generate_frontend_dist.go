//go:generate go run generate_frontend_dist.go
//go:build ignore

package main

import (
	"encoding/base64"
	"os"
	"path/filepath"
)

var (
	outputFilename = "domain/generated/frontend_dist.go"
	packageName    = "generated"
	variableName   = "FrontendDist"
	dirname, _     = filepath.Abs(os.Getenv("FRONTEND_DIST_DIR"))
)

func main() {
	if dirname == "" {
		panic("not set env FRONTEND_DIST_DIR")
	}

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	outputFile.WriteString("package " + packageName + "\n\n")
	outputFile.WriteString("var " + variableName + " = map[string]string {\n")

	err = filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		fileData, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dirname, path)
		if err != nil {
			return err
		}

		outputFile.WriteString("\"" + relPath + "\"" + ": " + "`" + base64.StdEncoding.EncodeToString(fileData) + "`" + ",\n")

		return nil
	})

	if err != nil {
		panic(err)
	}

	outputFile.WriteString("}\n")

}
