// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

// facadeVersions lists the best version of facades that we know about. This
// will be used to pick out a default version for communication, given the list
// of known versions that the API server tells us it is capable of supporting.
// This map should be updated whenever the API server exposes a new version (so
// that the client will use it whenever it is available).
// New facades should start at 1.
// Facades that existed before versioning start at 0.
var facadeVersions = map[string]int{
	"Action":           2,
	"Agent":            2,
	"Annotations":      2,
	"Backups":          1,
	"Bundle":           1,
	"Client":           1,
	"DiscoverSpaces":   2,
	"EntityWatcher":    2,
	"HighAvailability": 2,
	"HostKeyReporter":  1,
	"LifeFlag":         1,
	"LogForwarding":    1,
	"Logger":           1,
	"NotifyWatcher":    1,
	"Payloads":         1,
	"Pinger":           1,
	"ProxyUpdater":     1,
	"Reboot":           2,
	"Resources":        1,
	"Resumer":          2,
	"Singular":         1,
	"Spaces":           2,
	"StringsWatcher":   1,
	"Subnets":          2,
	"Undertaker":       1,
	"UnitAssigner":     1,
	"Uniter":           4,
}

// bestVersion tries to find the newest version in the version list that we can
// use.
func bestVersion(desiredVersion int, versions []int) int {
	best := 0
	for _, version := range versions {
		if version <= desiredVersion && version > best {
			best = version
		}
	}
	return best
}
