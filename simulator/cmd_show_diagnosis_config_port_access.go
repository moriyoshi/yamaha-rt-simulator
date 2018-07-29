package simulator

import "context"

var CmdShowDiagnosisConfigPortAccess = &CommandSpec{
	[]Token{TShow, TDiagnosis, TConfig, TPort, TAccess}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

