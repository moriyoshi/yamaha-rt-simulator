package simulator

import "context"

var CmdSipServerPilotAddress = &CommandSpec{
	[]Token{TSip, TServer, TPilot, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerPilotAddress = &CommandSpec{
	[]Token{TNo, TSip, TServer, TPilot, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

