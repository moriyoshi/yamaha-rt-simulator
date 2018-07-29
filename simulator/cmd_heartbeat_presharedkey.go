package simulator

import "context"

var CmdHeartbeatPreSharedKey = &CommandSpec{
	[]Token{THeartbeat, TPreSharedKey}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHeartbeatPreSharedKey = &CommandSpec{
	[]Token{TNo, THeartbeat, TPreSharedKey}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

