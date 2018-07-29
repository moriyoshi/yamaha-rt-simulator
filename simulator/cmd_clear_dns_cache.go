package simulator

import "context"

var CmdClearDnsCache = &CommandSpec{
	[]Token{TClear, TDns, TCache}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

