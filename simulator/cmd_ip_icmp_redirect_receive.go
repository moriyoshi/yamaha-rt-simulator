package simulator

import "context"

var CmdIpIcmpRedirectReceive = &CommandSpec{
	[]Token{TIp, TIcmp, TRedirect, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpRedirectReceive = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TRedirect, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

