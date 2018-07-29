package simulator

import "context"

var CmdIpInterfaceDhcpRetry = &CommandSpec{
	[]Token{TIp, TL2Interface, TDhcp, TRetry}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceDhcpRetryRetry = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TDhcp, TRetry}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

