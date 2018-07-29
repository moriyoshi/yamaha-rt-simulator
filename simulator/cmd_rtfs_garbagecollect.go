package simulator

import "context"

var CmdRtfsGarbageCollect = &CommandSpec{
	[]Token{TRtfs, TGarbageCollect}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

