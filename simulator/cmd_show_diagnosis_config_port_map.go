package simulator

import "context"

var CmdShowDiagnosisConfigPortMap = &CommandSpec{
	[]Token{TShow, TDiagnosis, TConfig, TPort, TMap}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

