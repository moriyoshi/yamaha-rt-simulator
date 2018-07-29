package simulator

import "context"

var CmdRipTimer = &CommandSpec{
	[]Token{TRip, TTimer}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRipTimer = &CommandSpec{
	[]Token{TNo, TRip, TTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

