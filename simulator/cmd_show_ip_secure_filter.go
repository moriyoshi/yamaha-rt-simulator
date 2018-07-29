package simulator

import "context"

var CmdShowIpSecureFilter = &CommandSpec{
	[]Token{TShow, TIp, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpSecureFilterPp = &CommandSpec{
	[]Token{TShow, TIp, TSecure, TFilter, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpSecureFilterTunnel = &CommandSpec{
	[]Token{TShow, TIp, TSecure, TFilter, TTunnel}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

