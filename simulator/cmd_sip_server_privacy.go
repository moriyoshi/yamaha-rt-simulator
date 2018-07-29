package simulator

import "context"

var CmdSipServerPrivacy = &CommandSpec{
	[]Token{TSip, TServer, TPrivacy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipServerPrivacy = &CommandSpec{
	[]Token{TNo, TSip, TServer, TPrivacy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

