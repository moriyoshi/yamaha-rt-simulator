package simulator

import "context"

var CmdFrPpDequeueType = &CommandSpec{
	[]Token{TFr, TPp, TDequeue, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrPpDequeueType = &CommandSpec{
	[]Token{TNo, TFr, TPp, TDequeue, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

