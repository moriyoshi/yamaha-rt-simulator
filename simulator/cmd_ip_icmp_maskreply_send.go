package simulator

import "context"

var CmdIpIcmpMaskReplySend = &CommandSpec{
	[]Token{TIp, TIcmp, TMaskReply, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpMaskReplySend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TMaskReply, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

