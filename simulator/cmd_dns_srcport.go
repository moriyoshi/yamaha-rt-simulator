package simulator

import "context"

var CmdDnsSrcport = &CommandSpec{
	[]Token{TDns, TSrcport}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsSrcport = &CommandSpec{
	[]Token{TNo, TDns, TSrcport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

