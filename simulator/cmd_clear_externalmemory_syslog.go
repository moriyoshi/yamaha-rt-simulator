package simulator

import "context"

var CmdClearExternalMemorySyslog = &CommandSpec{
	[]Token{TClear, TExternalMemory, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

