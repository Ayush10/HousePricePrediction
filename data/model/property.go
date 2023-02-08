package model

import (
	"fmt"
)

// Property represents the real estate property
type Property struct {
	Name        string
	Location    string
	Price       string
	Bedrooms    string
	Bathrooms   string
	Size        string
	Type        string
	YearBuilt   string
	Description string
}

// String representation of Property struct
func (p Property) String() string {
	return fmt.Sprintf("Name: %s\nLocation: %s\nPrice: %s\nBedrooms: %s\nBathrooms: %s\nSize: %s\nType: %s\nYear Built: %s\nDescription: %s\n\n", p.Name, p.Location, p.Price, p.Bedrooms, p.Bathrooms, p.Size, p.Type, p.YearBuilt, p.Description)
}
