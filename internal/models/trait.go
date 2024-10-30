package models

// Trait includes a name, a dominant boolean flag, and a gene label
type Trait struct {
	Name      string
	Dominant  bool
	GeneLabel string
}

// Plant includes a name and a list of traits
type Plant struct {
	Name   string
	Traits []Trait
}

// ErrorTraitResponse represents a trait error response
type ErrorTraitResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// TraitRequest represents a trait request
type TraitRequest struct {
	MotherPlant Plant `json:"motherPlant"`
	FatherPlant Plant `json:"fatherPlant"`
	PlantCount  int   `json:"plantCount"`
}

// TraitCombination represents a trait combination
type TraitCombination struct {
	Description    string  `json:"description"`
	Genotype       string  `json:"genotype"`
	NumberOfPlants int     `json:"numberOfPlants"`
	Percentage     float32 `json:"percentage"`
}

// TraitResponse represents a trait response
type TraitResponse struct {
	PlantCount      int                `json:"plantCount"`
	MotherPlantName string             `json:"motherPlantName"`
	FatherPlantName string             `json:"fatherPlantName"`
	Combinations    []TraitCombination `json:"combinations"`
}
