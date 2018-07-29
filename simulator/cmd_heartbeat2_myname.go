package simulator

import "context"

var CmdHeartbeat2Myname = &CommandSpec{
	[]Token{THeartbeat2, TMyname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeat2Myname = &CommandSpec{
	[]Token{TNo, THeartbeat2, TMyname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

