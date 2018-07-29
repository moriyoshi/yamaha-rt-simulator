package simulator

import "context"

var CmdIpIcmpEchoReplySend = &CommandSpec{
	[]Token{TIp, TIcmp, TEchoReply, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpEchoReplySend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TEchoReply, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

