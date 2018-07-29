package simulator

import "context"

var CmdLanMapSnapshotUse = &CommandSpec{
	[]Token{TLanMap, TSnapshot, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanMapSnapshotUse = &CommandSpec{
	[]Token{TNo, TLanMap, TSnapshot, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

