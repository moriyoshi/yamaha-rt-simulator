package simulator

import "context"

var CmdShowSshdPublicKey = &CommandSpec{
	[]Token{TShow, TSshd, TPublic, TKey}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

