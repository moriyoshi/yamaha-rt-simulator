package simulator

import "context"

var CmdShowIpv6Connection = &CommandSpec{
	[]Token{TShow, TIpv6, TConnection}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpv6ConnectionPp = &CommandSpec{
	[]Token{TShow, TIpv6, TConnection, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpv6ConnectionTunnel = &CommandSpec{
	[]Token{TShow, TIpv6, TConnection, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
