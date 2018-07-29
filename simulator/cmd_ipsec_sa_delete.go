package simulator

import "context"

var CmdIpsecSaDelete = &CommandSpec{
	[]Token{TIpsec, TSa, TDelete}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

