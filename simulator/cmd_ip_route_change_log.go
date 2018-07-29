package simulator

import "context"

var CmdIpRouteChangeLog = &CommandSpec{
	[]Token{TIp, TRoute, TChange, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpRouteChangeLog = &CommandSpec{
	[]Token{TNo, TIp, TRoute, TChange, TLog}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

