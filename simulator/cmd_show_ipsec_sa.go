package simulator

import "context"

var CmdShowIpsecSa = &CommandSpec{
	[]Token{TShow, TIpsec, TSa}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowIpsecSaGateway = &CommandSpec{
	[]Token{TShow, TIpsec, TSa, TGateway}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

