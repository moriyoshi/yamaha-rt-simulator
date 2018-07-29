package simulator

import "context"

var CmdSnmpTrapSendLinkdown = &CommandSpec{
	[]Token{TSnmp, TTrap, TSend, TLinkdown}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdSnmpTrapSendLinkdownPp = &CommandSpec{
	[]Token{TSnmp, TTrap, TSend, TLinkdown, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdSnmpTrapSendLinkdownTunnel = &CommandSpec{
	[]Token{TSnmp, TTrap, TSend, TLinkdown, TTunnel}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapSendLinkdown = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TSend, TLinkdown}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapSendLinkdownPp = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TSend, TLinkdown, TPp}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapSendLinkdownTunnel = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TSend, TLinkdown, TTunnel}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

