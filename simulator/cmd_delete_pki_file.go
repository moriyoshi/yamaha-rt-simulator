package simulator

import "context"

var CmdDeletePkiFile = &CommandSpec{
	[]Token{TDelete, TPki, TFile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

