package simulator

import "context"

var CmdEmbeddedFile = &CommandSpec{
	[]Token{TEmbedded, TFile}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoEmbeddedFile = &CommandSpec{
	[]Token{TNo, TEmbedded, TFile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

