package simulator

import "context"

var CmdClearPppoePassThroughLearning = &CommandSpec{
	[]Token{TClear, TPppoe, TPassThrough, TLearning}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

