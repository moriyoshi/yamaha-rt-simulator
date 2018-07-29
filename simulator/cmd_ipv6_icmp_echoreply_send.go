package simulator

import "context"

var CmdIpv6IcmpEchoReplySend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TEchoReply, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpEchoReplySendSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TEchoReply, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

