package simulator

import "context"

var CmdLanBackup = &CommandSpec{
	[]Token{TLan, TBackup}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdLanBackupInterfacePp = &CommandSpec{
	[]Token{TLan, TBackup, TL2Interface, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdLanBackupInterfaceTunnel = &CommandSpec{
	[]Token{TLan, TBackup, TL2Interface, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanBackup = &CommandSpec{
	[]Token{TNo, TLan, TBackup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
