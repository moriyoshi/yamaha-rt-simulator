package simulator

import "context"

var CmdExternalMemoryStatisticsFilenamePrefix = &CommandSpec{
	[]Token{TExternalMemory, TStatistics, TFilename, TPrefix}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryStatisticsFilenamePrefix = &CommandSpec{
	[]Token{TNo, TExternalMemory, TStatistics, TFilename, TPrefix}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

