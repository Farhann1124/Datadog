package pulumi_env

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DataDog/datadog-agent/test/new-e2e/runner"
	"github.com/DataDog/datadog-agent/test/new-e2e/utils/infra"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"os"
)

func SetupConfig() runner.ConfigMap {
	res := runner.ConfigMap{}
	config := os.Getenv("PULUMI_CONFIG")
	if config != "" {
		var result map[string]any
		err := json.Unmarshal([]byte(config), &result)
		if err != nil {
			configs := result["config"].(map[string]string)
			for key, value := range configs {
				res[key] = auto.ConfigValue{Value: value}
			}
		}
	}

	return res
}

func TearDown() []error {
	fmt.Fprintf(os.Stderr, "Cleaning up stacks")
	errs := infra.GetStackManager().Cleanup(context.Background())
	if errs != nil {
		return errs
	}
	return nil
}
