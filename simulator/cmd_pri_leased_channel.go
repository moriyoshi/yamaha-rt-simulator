package simulator

import "context"

var CmdPriLeasedChannel = &CommandSpec{
	[]Token{TPri, TLeased, TChannel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPriLeasedChannel = &CommandSpec{
	[]Token{TNo, TPri, TLeased, TChannel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

