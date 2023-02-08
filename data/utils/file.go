package utils

import (
	_ "HousePricePrediction/data/model"
	"encoding/csv"
	"fmt"
	"os"
)

// SaveProperties saves the properties to a file
func SaveProperties(filename string, properties []Property) error {
	ext := getFileExtension(filename)
	switch ext {
	case ".xlsx":
		return savePropertiesAsExcel(filename, properties)
	case ".csv":
		return savePropertiesAsCSV(filename, properties)
	default:
		return fmt.Errorf("unsupported file format: %s", ext)
	}
}

func savePropertiesAsExcel(filename string, properties []Property) error {
	f, err := excelize.CreateFile()
	if err != nil {
		return err
	}
	defer f.SaveAs(filename)

	// Add a new sheet for the properties
	sheetName := "Properties"
	f.NewSheet(sheetName)

	// Add header row
	headerRow := []string{"Title", "Location", "Price", "Bedrooms", "Bathrooms", "Size", "Type", "Year Built", "View"}
	f.SetRow(sheetName, 1, &excelize.Row{Cells: headerRow})

	// Add data rows
	for i, property := range properties {
		row := []string{property.Title, property.Location, property.Price, property.Bedrooms, property.Bathrooms, property.Size, property.Type, property.YearBuilt, property.View}
		f.SetRow(sheetName, i+2, &excelize.Row{Cells: row})
	}

	return nil
}

func savePropertiesAsCSV(filename string, properties []Property) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	headerRow := []string{"Title", "Location", "Price", "Bedrooms", "Bathrooms", "Size", "Type", "Year Built", "View"}
	err = writer.Write(headerRow)
	if err != nil {
		return err
	}

	// Write data rows
	for _, property := range properties {
		row := []string{property.Title, property.Location, property.Price, property.Bedrooms, property.Bathrooms, property.Size, property.Type, property.YearBuilt, property.View}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileExtension(filename string) string {
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[i:]
		}
	}
	return ""
}
