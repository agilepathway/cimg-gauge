package docker

import (
	"os"
	"regexp"
)

var (
	GaugeVersion = SoftwareVersion{
		command:   "gauge --version",
		trimRegex: `Gauge version: (\d+(.\d+)?(.\d+)?)`,
		dummy:     GaugeDummyVersion}
	ChromeVersion = SoftwareVersion{
		command:   "google-chrome --version",
		trimRegex: `Google Chrome ([0|1-9|\.]*)`,
		dummy:     "Google Chrome 81.0.4044.92"}
	GoVersion = SoftwareVersion{
		command:   "go version",
		trimRegex: `go version go(\d+(.\d+)?(.\d+)?)`,
		dummy:     "go version go1.0.2 darwin/amd64"}
	NodeVersion = SoftwareVersion{
		command:   "node --version",
		trimRegex: `v(\d+(.\d+)?(.\d+)?)`,
		dummy:     "v1.0.3"}
	TaikoVersion = SoftwareVersion{
		command:   "npm ls taiko -global -parseable -long",
		trimRegex: `taiko@(\d+(.\d+)?(.\d+)?)`,
		dummy:     "/home/circleci/.npm-global/lib/node_modules/taiko:taiko@1.0.4:undefined"}
)

const GaugeDummyVersion = `Gauge version: 1.0.1
Commit Hash: 1a2b345

Plugins
-------
html-report (4.0.6)
js (2.3.4)
screenshot (0.0.1)`

type SoftwareVersion struct {
	command   string
	trimRegex string
	dummy     string
}

func (sv SoftwareVersion) FindAndTrim() string {
	verboseVersion := sv.findVerbose(sv.command)
	return sv.trim(verboseVersion)
}

func (sv SoftwareVersion) findVerbose(command string) string {
	if doesEnvVarExist("INTEGRATION_TEST") {
		return sv.dummy
	}
	return Execute(command)
}

func (sv SoftwareVersion) trim(verbose string) string {
	compiledRegex := regexp.MustCompile(sv.trimRegex)
	return compiledRegex.FindStringSubmatch(verbose)[1]
}

func doesEnvVarExist(envVar string) bool {
	_, exists := os.LookupEnv(envVar)
	return exists
}
