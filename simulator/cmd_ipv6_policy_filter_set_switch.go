package simulator

import "context"

var CmdIpv6PolicyFilterSetSwitch = &CommandSpec{
	[]Token{TIpv6, TPolicy, TFilter, TSet, TSwitch}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
