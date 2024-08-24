package testutil

import (
	"encoding/xml"
	"github.com/goccy/go-json"
	"os"
	"path/filepath"
)

const TestDataDir = "./testdata"

func SaveTestDataXML(filename string, data interface{}) error {
	return saveTestDataXML(TestDataDir, filename, data)
}

func LoadTestDataXML(filename string, data interface{}) error {
	return loadTestDataXML(TestDataDir, filename, data)
}

func saveTestDataXML(dir, filename string, data interface{}) error {
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the data and write it to the file
	enc := xml.NewEncoder(file)
	defer enc.Close()
	if err := enc.Encode(data); err != nil {
		return err
	}

	return nil
}

func loadTestDataXML(dir, filename string, data interface{}) error {
	file, err := os.Open(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	dec := xml.NewDecoder(file)
	if err := dec.Decode(data); err != nil {
		return err
	}

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func SaveTestDataJSON(filename string, data interface{}) error {
	return saveTestDataJSON(TestDataDir, filename, data)
}

func LoadTestDataJSON(filename string, data interface{}) error {
	return loadTestDataJSON(TestDataDir, filename, data)
}

func saveTestDataJSON(dir, filename string, data interface{}) error {
	// Create the directory if it doesn't exist
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the data and write it to the file
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		return err
	}

	return nil
}

func loadTestDataJSON(dir, filename string, data interface{}) error {
	file, err := os.Open(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(data); err != nil {
		return err
	}

	return nil
}
