package simulator

import "context"

var CmdShowPkiCertificateSummary = &CommandSpec{
	[]Token{TShow, TPki, TCertificate, TSummary}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

