package domain

import (
	"encoding/json"
	"os"
)

type(
	// Tax
	TaxData struct {
		StateAbbreviation string `json:"stateAbbreviation"`
		StateName string `json:"stateName"`
		TaxRate float64 `json:"taxRate"`
	}
	TaxContainer struct {
		TaxDataList []*TaxData
		IsCreated   bool
	}
)

func(t *TaxContainer) WriteJson() error {
	file, err := t.GetOrCreateFile()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(t.TaxDataList,JSONPrefix, JSONIndent)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func(t *TaxContainer) WasFileCreated() bool {
	return t.IsCreated
}

func(t *TaxContainer) FilePath() string {
	return MainDirectory + TaxFilePath
}

func(t *TaxContainer) createFile() (*os.File, error) {
	jsonFile, err := os.Create(t.FilePath())
	if err != nil {
		return nil,err
	}
	t.IsCreated = true
	return jsonFile,nil
}

func(t *TaxContainer) RemoveFile() error {
	err := os.Remove(t.FilePath())
	if err != nil {
		return err
	}
	t.IsCreated = false
	return nil
}

func(t *TaxContainer) openAndGetFile() (*os.File, error) {
	jsonFile, err := os.OpenFile(t.FilePath(), os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func(t *TaxContainer) GetOrCreateFile() (*os.File, error) {
	if t.IsCreated {
		return t.openAndGetFile()
	} else {
		return t.createFile()
	}
}