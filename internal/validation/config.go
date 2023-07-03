package validation

import (
	"fmt"

	"github.com/gookit/validate"
	"github.com/snoopy82481/clusterflux/internal/config"
	"github.com/snoopy82481/clusterflux/internal/logger"
)

func ValidateConfig(config *config.Config) error {
	v := validate.Struct(config)

	// custom validation

	if !v.Validate() {
		for _, err := range v.Errors {
			logger.LogWarn(err.String())
		}

		return fmt.Errorf("Config validation failed")
	}

	return nil
}
