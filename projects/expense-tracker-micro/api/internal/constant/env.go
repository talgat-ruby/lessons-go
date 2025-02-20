package constant

// Environment is the enumeration for the various runtime environments
type Environment string

const (
	EnvironmentLocal Environment = "LOCAL"
	EnvironmentTest  Environment = "TEST"
	EnvironmentDev   Environment = "DEV"
	EnvironmentProd  Environment = "PROD"
)
