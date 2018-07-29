package simulator

import "context"

var CmdBridgeLearningBridgeInterfaceTimer = &CommandSpec{
	[]Token{TBridge, TLearning, TSomething /* bridge_interface */, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBridgeLearningBridgeInterfaceTimer = &CommandSpec{
	[]Token{TNo, TBridge, TLearning, TSomething /* bridge_interface */, TTimer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

