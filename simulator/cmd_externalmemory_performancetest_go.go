package simulator

import "context"

var CmdExternalMemoryPerformanceTestGo = &CommandSpec{
	[]Token{TExternalMemory, TPerformanceTest, TGo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

