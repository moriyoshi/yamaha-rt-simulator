package simulator

import "context"

var CmdShowStatusPacketBuffer = &CommandSpec{
	[]Token{TShow, TStatus, TPacketBuffer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

