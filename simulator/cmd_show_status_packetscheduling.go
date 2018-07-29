package simulator

import "context"

var CmdShowStatusPacketScheduling = &CommandSpec{
	[]Token{TShow, TStatus, TPacketScheduling}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

