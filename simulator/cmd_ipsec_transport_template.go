package simulator

import "context"

var CmdIpsecTransportTemplate = &CommandSpec{
	[]Token{TIpsec, TTransport, TTemplate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecTransportTemplate = &CommandSpec{
	[]Token{TNo, TIpsec, TTransport, TTemplate}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

