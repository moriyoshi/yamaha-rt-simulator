package simulator

import "context"

var CmdSipIpProtocol = &CommandSpec{
	[]Token{TSip, TIp, TProtocol}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipIpProtocol = &CommandSpec{
	[]Token{TNo, TSip, TIp, TProtocol}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

