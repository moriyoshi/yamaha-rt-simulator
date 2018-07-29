package simulator

import "context"

var CmdDnsStatic = &CommandSpec{
	[]Token{TDns, TStatic}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
