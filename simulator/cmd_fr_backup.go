package simulator

import "context"

var CmdFrBackup = &CommandSpec{
	[]Token{TFr, TBackup}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrBackup = &CommandSpec{
	[]Token{TNo, TFr, TBackup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

