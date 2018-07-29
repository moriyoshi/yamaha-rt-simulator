package simulator

import "context"

var CmdShowConfigTunnel = &CommandSpec{
	[]Token{TShow, TConfig, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdLessConfigTunnel = &CommandSpec{
	[]Token{TLess, TConfig, TTunnel}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
