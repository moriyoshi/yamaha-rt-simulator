package simulator

import "context"

var CmdPkiCertificateFile = &CommandSpec{
	[]Token{TPki, TCertificate, TFile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPkiCertificateFile = &CommandSpec{
	[]Token{TNo, TPki, TCertificate, TFile}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
