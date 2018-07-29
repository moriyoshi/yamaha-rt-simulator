package simulator

import "context"

var CmdShowIpConnection = &CommandSpec{
	[]Token{TShow, TIp, TConnection}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpConnectionPp = &CommandSpec{
	[]Token{TShow, TIp, TConnection, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpConnectionTunnel = &CommandSpec{
	[]Token{TShow, TIp, TConnection, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
