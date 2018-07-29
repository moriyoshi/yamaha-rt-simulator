package simulator

import "context"

var CmdRotateExternalMemorySyslog = &CommandSpec{
	[]Token{TRotate, TExternalMemory, TSyslog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

