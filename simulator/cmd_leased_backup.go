package simulator

import "context"

var CmdLeasedBackup = &CommandSpec{
	[]Token{TLeased, TBackup}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLeasedBackup = &CommandSpec{
	[]Token{TNo, TLeased, TBackup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

