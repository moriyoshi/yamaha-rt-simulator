package simulator

import "context"

var CmdShowIpTrafficList = &CommandSpec{
	[]Token{TShow, TIp, TTraffic, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpTrafficListPp = &CommandSpec{
	[]Token{TShow, TIp, TTraffic, TList, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpTrafficListTunnel = &CommandSpec{
	[]Token{TShow, TIp, TTraffic, TList, TTunnel}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

