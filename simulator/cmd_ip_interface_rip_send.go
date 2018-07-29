package simulator

import "context"

var CmdIpInterfaceRipSend = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipSend = &CommandSpec{
	[]Token{TIp, TPp, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipSend = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipSend = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipSend = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipSend = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

