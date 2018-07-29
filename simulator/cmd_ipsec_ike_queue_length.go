package simulator

import "context"

var CmdIpsecIkeQueueLength = &CommandSpec{
	[]Token{TIpsec, TIke, TQueue, TLength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeQueueLength = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TQueue, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

