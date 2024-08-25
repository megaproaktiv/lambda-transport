package lambda_test

import (
	"testing"

	mylambda "github.com/tecracer/lambda-transport/lambda"
	"gotest.tools/v3/assert"
)

func TestConfiguration(t *testing.T) {
	t.Log("TestConfigurations")
	testConfig := "testdata/config-example.yml"
	cfg, err := mylambda.ReadConfig(testConfig)
	assert.NilError(t, err, "Reading config should give no error")
	actual := cfg.Cfg["dev"].Source.LambdaName
	expected := "demo"
	assert.Equal(t, actual, expected, "Source lambda name should be demo")
}
