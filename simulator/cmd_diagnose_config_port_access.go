package simulator

import "context"

var CmdDiagnoseConfigPortAccess = &CommandSpec{
	[]Token{TDiagnose, TConfig, TPort, TAccess}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

