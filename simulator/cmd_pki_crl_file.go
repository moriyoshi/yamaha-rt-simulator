package simulator

import "context"

var CmdPkiCrlFile = &CommandSpec{
	[]Token{TPki, TCrl, TFile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPkiCrlFile = &CommandSpec{
	[]Token{TNo, TPki, TCrl, TFile}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
