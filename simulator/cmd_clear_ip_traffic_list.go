package simulator

import "context"

var CmdClearIpTrafficList = &CommandSpec{
	[]Token{TClear, TIp, TTraffic, TList, &EitherToken{TPp, TTunnel}}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
