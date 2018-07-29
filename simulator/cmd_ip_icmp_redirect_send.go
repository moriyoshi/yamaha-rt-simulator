package simulator

import "context"

var CmdIpIcmpRedirectSend = &CommandSpec{
	[]Token{TIp, TIcmp, TRedirect, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpRedirectSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TRedirect, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

