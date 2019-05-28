package test

import (
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleExample(t *testing.T) {
    t.Parallel()
	exampleDir := test_structure.CopyTerraformFolderToTemp(t, "../", "")

	// Get the Project id 
	projectId := gcp.GetGoogleProjectIDFromEnvVar(t)

    // Set expected results
    expectedComputeName := "simple-example-compute"
	expectedZone := "europe-west4-a"
	expectedLabels := map[string]string{"simple": "example"}


	terraformOptions := &terraform.Options{
		TerraformDir: exampleDir,
		Vars: map[string]interface{}{
			"gcp_project_id": projectId,
		},
	}

	// Run `terraform destry` after the test 
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)

	// Get outputs from terraform apply

	// get ressources and their definitions
	instance := gcp.FetchInstance(t, projectId, expectedComputeName)
    resultLabels := instance.GetLabels(t)
    resultZone := instance.GetZone(t)


	// Perform the tests
    assert.Equal(t, expectedZone, resultZone)

	for expectedLabel, expectedLabelValue := range expectedLabels {
		resultLabelValue, resultLabelExists := resultLabels[expectedLabel]
		if !resultLabelExists {
			t.Errorf("Label '%s' is expecting to exist", expectedLabel)
		}
	    assert.Equal(t, expectedLabelValue, resultLabelValue)
	}
}
