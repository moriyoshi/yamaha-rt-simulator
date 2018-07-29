package simulator

import "context"

var CmdIpsecTransport = &CommandSpec{
	[]Token{TIpsec, TTransport}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecTransport = &CommandSpec{
	[]Token{TNo, TIpsec, TTransport}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

