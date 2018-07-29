package simulator

import "context"

var CmdDiagnosisConfigPortHistoryNum = &CommandSpec{
	[]Token{TDiagnosis, TConfig, TPort, THistoryNum}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

