package simulator

import "context"

var CmdShowIpv6Address = &CommandSpec{
	[]Token{TShow, TIpv6, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpv6AddressPp = &CommandSpec{
	[]Token{TShow, TIpv6, TAddress, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpv6AddressTunnel = &CommandSpec{
	[]Token{TShow, TIpv6, TAddress, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

