package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformInfrastructure(t *testing.T) {
	// Lire les variables d'environnement
	resourceGroupName := os.Getenv("RESOURCE_GROUP_NAME")
	acrName := os.Getenv("ACR_NAME")
	location := os.Getenv("LOCATION")

	// Définir les options Terraform
	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"acr_name":            acrName,
			"location":            location,
		},
	}

	// Détruire l'infrastructure à la fin du test
	defer terraform.Destroy(t, terraformOptions)

	// Initialiser et appliquer Terraform
	terraform.InitAndApply(t, terraformOptions)

	// Récupérer la sortie Terraform
	outputACRName := terraform.Output(t, terraformOptions, "acr_login_server") // correspond à ton output Terraform
	assert.Equal(t, acrName+".azurecr.io", outputACRName)
}
