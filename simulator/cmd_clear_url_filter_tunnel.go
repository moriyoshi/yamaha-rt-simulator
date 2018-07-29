package simulator

import "context"

var CmdClearUrlFilterTunnel = &CommandSpec{
	[]Token{TClear, TUrl, TFilter, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
