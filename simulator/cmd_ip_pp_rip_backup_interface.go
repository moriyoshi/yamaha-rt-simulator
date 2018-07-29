package simulator

import "context"

var CmdIpPpRipBackupInterface = &CommandSpec{
	[]Token{TIp, TPp, TRip, TBackup, TInterface}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpPpRipBackupInterface = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TBackup, TInterface}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
