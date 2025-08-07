package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewAppRepo,
	NewBackupRepo,
	NewCacheRepo,
	NewCertRepo,
	NewCertAccountRepo,
	NewCertDNSRepo,
	NewContainerRepo,
	NewContainerComposeRepo,
	NewContainerImageRepo,
	NewContainerNetworkRepo,
	NewContainerVolumeRepo,
	NewCronRepo,
	NewDatabaseRepo,
	NewDatabaseServerRepo,
	NewDatabaseUserRepo,
	NewMonitorRepo,
	NewSafeRepo,
	NewSettingRepo,
	NewSSHRepo,
	NewTaskRepo,
	NewUserRepo,
	NewUserTokenRepo,
	NewWebsiteRepo,
)
