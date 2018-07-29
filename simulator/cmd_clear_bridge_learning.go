package simulator

import "context"

var CmdClearBridgeLearning = &CommandSpec{
	[]Token{TClear, TBridge, TLearning}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
