package simulator

import "context"

var CmdShowLanMap = &CommandSpec{
	[]Token{TShow, TLanMap}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

