package simulator

import "context"

var CmdShowFileList = &CommandSpec{
	[]Token{TShow, TFile, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdLessFileList = &CommandSpec{
	[]Token{TLess, TFile, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

