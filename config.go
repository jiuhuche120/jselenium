package jselenium

var Config *RunConfig
type RunConfig struct {
	DriverType string
	Port       int
	Rerun      uint64
	MaxFail    uint64
	Driver     string
	ReportPath string
}

func DefaultConfig(driverType, driver string) RunConfig {
	return RunConfig{
		DriverType: driverType,
		Port:       7222,
		Rerun:      1,
		MaxFail:    5,
		Driver:     driver,
		ReportPath: "./reports",
	}
}

func LoadConfig(driverType string, port int, rerun, maxFail uint64, driver, reportPath string) RunConfig {
	return RunConfig{
		DriverType: driverType,
		Port:       port,
		Rerun:      rerun,
		MaxFail:    maxFail,
		Driver:     driver,
		ReportPath: reportPath,
	}
}
