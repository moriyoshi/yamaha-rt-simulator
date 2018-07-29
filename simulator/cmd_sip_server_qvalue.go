package simulator

import "context"

var CmdSipServerQvalue = &CommandSpec{
	[]Token{TSip, TServer, TQvalue}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerQvalue = &CommandSpec{
	[]Token{TNo, TSip, TServer, TQvalue}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

