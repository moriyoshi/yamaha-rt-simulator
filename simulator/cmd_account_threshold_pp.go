package simulator

import "context"

var CmdAccountThresholdPp = &CommandSpec{
	[]Token{TAccount, TThreshold, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
