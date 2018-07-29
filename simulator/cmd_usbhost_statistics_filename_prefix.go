package simulator

import "context"

var CmdUsbhostStatisticsFilenamePrefix = &CommandSpec{
	[]Token{TUsbhost, TStatistics, TFilename, TPrefix}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostStatisticsFilenamePrefix = &CommandSpec{
	[]Token{TNo, TUsbhost, TStatistics, TFilename, TPrefix}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

