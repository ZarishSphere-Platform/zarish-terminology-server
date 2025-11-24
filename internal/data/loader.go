package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/zarishsphere-platform/zarish-terminology-server/internal/database"
	"github.com/zarishsphere-platform/zarish-terminology-server/internal/models"
	"gorm.io/datatypes"
)

func LoadTerminologyData(dataDir string) {
	// Load CodeSystems
	patterns := []string{
		filepath.Join(dataDir, "codesystems", "*.json"),
		filepath.Join(dataDir, "terminology", "bangladesh", "*.json"),
		filepath.Join(dataDir, "terminology", "bd-core", "package", "*.json"),
	}

	var files []string
	for _, p := range patterns {
		matches, err := filepath.Glob(p)
		if err != nil {
			log.Printf("Failed to glob pattern %s: %v", p, err)
			continue
		}
		files = append(files, matches...)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Printf("Failed to read file %s: %v", file, err)
			continue
		}

		var resource map[string]interface{}
		if err := json.Unmarshal(content, &resource); err != nil {
			log.Printf("Failed to parse JSON %s: %v", file, err)
			continue
		}

		resourceType, _ := resource["resourceType"].(string)
		id, _ := resource["id"].(string)
		url, _ := resource["url"].(string)
		name, _ := resource["name"].(string)
		title, _ := resource["title"].(string)
		status, _ := resource["status"].(string)

		if resourceType == "CodeSystem" {
			cs := models.CodeSystem{
				ID:        id,
				URL:       url,
				Name:      name,
				Title:     title,
				Status:    status,
				Content:   datatypes.JSON(content),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := database.DB.Save(&cs).Error; err != nil {
				log.Printf("Failed to save CodeSystem %s: %v", id, err)
			} else {
				log.Printf("Loaded CodeSystem: %s", id)
			}
		} else if resourceType == "ValueSet" {
			vs := models.ValueSet{
				ID:        id,
				URL:       url,
				Name:      name,
				Title:     title,
				Status:    status,
				Content:   datatypes.JSON(content),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := database.DB.Save(&vs).Error; err != nil {
				log.Printf("Failed to save ValueSet %s: %v", id, err)
			} else {
				log.Printf("Loaded ValueSet: %s", id)
			}
		} else if resourceType == "StructureDefinition" {
			kind, _ := resource["kind"].(string)
			type_, _ := resource["type"].(string)
			
			sd := models.StructureDefinition{
				ID:        id,
				URL:       url,
				Name:      name,
				Title:     title,
				Status:    status,
				Kind:      kind,
				Type:      type_,
				Content:   datatypes.JSON(content),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			if err := database.DB.Save(&sd).Error; err != nil {
				log.Printf("Failed to save StructureDefinition %s: %v", id, err)
			} else {
				log.Printf("Loaded StructureDefinition: %s", id)
			}
		}
	}
}
