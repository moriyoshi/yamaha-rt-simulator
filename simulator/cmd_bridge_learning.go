package simulator

import "context"

var CmdBridgeLearningSwitch = &CommandSpec{
	[]Token{TBridge, TLearning, TSomething, TOnOff}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBridgeLearningSwitch = &CommandSpec{
	[]Token{TNo, TBridge, TLearning, TSomething, TOnOff}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdBridgeLearningStatic = &CommandSpec{
	[]Token{TBridge, TLearning, TSomething, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBridgeLearningStatic = &CommandSpec{
	[]Token{TNo, TBridge, TLearning, TSomething, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
