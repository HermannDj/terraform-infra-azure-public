package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructure(t *testing.T) {
	// Récupération des variables d'environnement
	resourceGroupName := os.Getenv("RESOURCE_GROUP_NAME")
	acrName := os.Getenv("ACR_NAME")
	location := os.Getenv("LOCATION")

	if resourceGroupName == "" || acrName == "" || location == "" {
		t.Fatal("Environment variables RESOURCE_GROUP_NAME, ACR_NAME, and LOCATION must be set")
	}

	terraformOptions := &terraform.Options{
		// Dossier où se trouvent les fichiers Terraform
		TerraformDir: "../terraform",

		// Variables à passer à Terraform
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"acr_name":            acrName,
			"location":            location,
		},

		// Auto-approve pour les tests
		NoColor: true,
	}

	// Détruire les ressources à la fin du test
	defer terraform.Destroy(t, terraformOptions)

	// Init et Apply
	terraform.InitAndApply(t, terraformOptions)

	// Vérifier la sortie Terraform
	outputACRName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputACRName)
}
