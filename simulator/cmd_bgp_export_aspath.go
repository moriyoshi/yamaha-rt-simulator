package simulator

import "context"

var CmdBgpExportAspath = &CommandSpec{
	[]Token{TBgp, TExport, TAspath}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
