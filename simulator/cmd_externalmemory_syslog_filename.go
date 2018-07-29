package simulator

import "context"

var CmdExternalMemorySyslogFilename = &CommandSpec{
	[]Token{TExternalMemory, TSyslog, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemorySyslogFilename = &CommandSpec{
	[]Token{TNo, TExternalMemory, TSyslog, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

