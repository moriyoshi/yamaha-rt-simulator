package simulator

import "context"

var CmdMobileDisplayCallerId = &CommandSpec{
	[]Token{TMobile, TDisplay, TCaller, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileDisplayCallerId = &CommandSpec{
	[]Token{TNo, TMobile, TDisplay, TCaller, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

