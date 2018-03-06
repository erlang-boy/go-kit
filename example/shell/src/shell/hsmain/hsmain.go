package hsmain

import (
	"fmt"
	"github.com/juju/loggo"
	"shell/config"
	daemon "shell/masterDaemon"
	"runtime"
)

var logger = loggo.GetLogger("hsmain.run")

func Run(shellName string) {
	showVersionInfo(shellName)
	showBuildInfo()

	// load file
	conf := config.Load()

	loggo.ConfigureLoggers(conf.LogLevel)

	logger.Infof("%+v", conf)

	// boot daemon
	err := daemon.NewDaemon(shellName, conf).Run()

	// todo cmd error handle
	if err != nil {
		logger.Errorf("%+v\n", err)
	}
}

func showVersionInfo(shellName string) {
	fmt.Print(versionString(shellName))
}

func showBuildInfo() {
	fmt.Print(releaseString())
	fmt.Print(commitString())
}

func versionString(shellName string) string {
	return fmt.Sprintf("%s-%s\n", shellName, HiveShellVersion)
}

func releaseString() string {
	return fmt.Sprintf("%s/%s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
}

func commitString() string {
	return fmt.Sprintf("build date: %s\nlast commit: %s\n ", BuildDate, GitCommit)
}
