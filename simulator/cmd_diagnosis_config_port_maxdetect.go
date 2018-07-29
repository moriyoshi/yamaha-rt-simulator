package simulator

import "context"

var CmdDiagnosisConfigPortMaxDetect = &CommandSpec{
	[]Token{TDiagnosis, TConfig, TPort, TMaxDetect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

