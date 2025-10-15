package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)


func TestTerraformInfrastructure(t *testing.T) {
	resourceGroupName := os.Getenv("RESOURCE_GROUP_NAME")
	acrName := os.Getenv("acr_name")
	location := os.Getenv("location")

	terraformOptions := &terraform.Options{
		TerraformDir: "../terraform",
		Vars: map[string]interface{}{
			"resource_group_name": resourceGroupName,
			"acr_name":            acrName,
			"location":            location,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	t
	erraform.InitAndApply(t, terraformOptions)s

	outputACRName := terraform.Output(t, terraformOptions, "acr_name")
	assert.Equal(t, acrName, outputACRName)

}