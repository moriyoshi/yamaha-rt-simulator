package simulator

import "context"

var CmdJateNumber = &CommandSpec{
	[]Token{TJate, TNumber}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoJateNumber = &CommandSpec{
	[]Token{TNo, TJate, TNumber}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

