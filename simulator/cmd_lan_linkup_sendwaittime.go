package simulator

import "context"

var CmdLanLinkupSendWaitTime = &CommandSpec{
	[]Token{TLan, TLinkup, TSendWaitTime}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanLinkupSendWaitTime = &CommandSpec{
	[]Token{TNo, TLan, TLinkup, TSendWaitTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

