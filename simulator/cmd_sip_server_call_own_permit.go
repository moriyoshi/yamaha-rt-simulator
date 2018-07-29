package simulator

import "context"

var CmdSipServerCallOwnPermit = &CommandSpec{
	[]Token{TSip, TServer, TCall, TOwn, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerCallOwnPermit = &CommandSpec{
	[]Token{TNo, TSip, TServer, TCall, TOwn, TPermit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

