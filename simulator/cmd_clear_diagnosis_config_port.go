package simulator

import "context"

var CmdClearDiagnosisConfigPort = &CommandSpec{
	[]Token{TClear, TDiagnosis, TConfig, TPort}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

