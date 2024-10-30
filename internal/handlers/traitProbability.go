package handlers

import (
	"breeding-rest/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"sort"
	"strings"
)

func calculateDenominator(numTraits int) int {
	return int(math.Pow(2, float64(2*numTraits)))
}

func calculateProbability(numTraits, numRecessive int) int {
	dominantTraits := numTraits - numRecessive
	return int(math.Pow(3, float64(dominantTraits)))
}

func calculateF2Probabilities(plant1, plant2 models.Plant, totalPlants int) ([]models.TraitCombination, error) {
	numTraits := len(plant1.Traits)
	if numTraits != len(plant2.Traits) {
		return nil, errors.New("plants must have same number of traits")
	}

	denominator := calculateDenominator(numTraits)
	numCombinations := int(math.Pow(2, float64(numTraits)))
	combinations := make([]models.TraitCombination, 0)

	for i := 0; i < numCombinations; i++ {
		traits := make([]bool, numTraits)
		description := make([]string, 0)
		geneNotation := make([]string, 0)
		numRecessive := 0

		for j := 0; j < numTraits; j++ {
			traits[j] = (i & (1 << j)) == 0
			if !traits[j] {
				numRecessive++
			}

			var traitName string
			if traits[j] {
				if plant1.Traits[j].Dominant {
					traitName = plant1.Traits[j].Name
				} else {
					traitName = plant2.Traits[j].Name
				}
				geneNotation = append(geneNotation, plant1.Traits[j].GeneLabel+"_")
			} else {
				if plant1.Traits[j].Dominant {
					traitName = plant2.Traits[j].Name
				} else {
					traitName = plant1.Traits[j].Name
				}
				geneNotation = append(geneNotation, strings.ToLower(plant1.Traits[j].GeneLabel+plant1.Traits[j].GeneLabel))
			}
			description = append(description, traitName)
		}

		probability := calculateProbability(numTraits, numRecessive)
		expectedPlants := int(math.Round(float64(probability) * float64(totalPlants) / float64(denominator)))
		percentage := float32(probability) / float32(denominator) * 100

		combinations = append(combinations, models.TraitCombination{
			Description:    strings.Join(description, ", "),
			Genotype:       strings.Join(geneNotation, " "),
			NumberOfPlants: expectedPlants,
			Percentage:     percentage,
		})
	}

	// Sort by percentage descending
	sort.Slice(combinations, func(i, j int) bool {
		return combinations[i].Percentage > combinations[j].Percentage
	})

	return combinations, nil
}

// @Summary Calculate Trait Probability
// @Description Calculate the probability of a trait based on the genotype of two plants
// @Tags trait-probability
// @Accept json
// @Produce json
// @Param trait body models.TraitRequest true "Trait probability calculation request"
// @Success 200 {object} models.TraitResponse
// @Failure 400 {object} models.ErrorTraitResponse
// @Router /trait-probability [post]
func CalculateTraitProbability(c *gin.Context) {
	var request models.TraitRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorTraitResponse{
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	// Validate input
	if len(request.MotherPlant.Traits) != len(request.FatherPlant.Traits) {
		c.JSON(http.StatusBadRequest, models.ErrorTraitResponse{
			Message: "Invalid request",
			Error:   "Mother and father plants must have the same number of traits",
		})
		return
	}

	if request.PlantCount <= 0 {
		c.JSON(http.StatusBadRequest, models.ErrorTraitResponse{
			Message: "Invalid request",
			Error:   "Plant count must be greater than 0",
		})
		return
	}

	combinations, err := calculateF2Probabilities(
		request.MotherPlant,
		request.FatherPlant,
		request.PlantCount,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorTraitResponse{
			Message: "Calculation error",
			Error:   err.Error(),
		})
		return
	}

	if combinations == nil {
		c.JSON(http.StatusBadRequest, models.ErrorTraitResponse{
			Message: "Calculation error",
			Error:   "Failed to calculate probabilities",
		})
		return
	}

	response := models.TraitResponse{
		PlantCount:      request.PlantCount,
		MotherPlantName: request.MotherPlant.Name,
		FatherPlantName: request.FatherPlant.Name,
		Combinations:    combinations,
	}

	c.JSON(http.StatusOK, response)
}
