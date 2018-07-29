package simulator

import "context"

var CmdSetSerialBaudrate = &CommandSpec{
	[]Token{TSetSerialBaudrate}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

