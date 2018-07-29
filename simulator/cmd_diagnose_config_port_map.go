package simulator

import "context"

var CmdDiagnoseConfigPortMap = &CommandSpec{
	[]Token{TDiagnose, TConfig, TPort, TMap}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

