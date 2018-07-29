package simulator

import "context"

var CmdLanBackupRecoveryTime = &CommandSpec{
	[]Token{TLan, TBackup, TRecovery, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanBackupRecoveryTime = &CommandSpec{
	[]Token{TNo, TLan, TBackup, TRecovery, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

