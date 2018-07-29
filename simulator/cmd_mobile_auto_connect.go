package simulator

import "context"

var CmdMobileAutoConnect = &CommandSpec{
	[]Token{TMobile, TAuto, TConnect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAutoConnect = &CommandSpec{
	[]Token{TNo, TMobile, TAuto, TConnect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

