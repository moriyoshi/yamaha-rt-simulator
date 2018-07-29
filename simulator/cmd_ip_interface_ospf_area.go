package simulator

import "context"

var CmdIpInterfaceOspfArea = &CommandSpec{
	[]Token{TIp, TL2Interface, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpOspfArea = &CommandSpec{
	[]Token{TIp, TPp, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelOspfArea = &CommandSpec{
	[]Token{TIp, TTunnel, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceOspfArea = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpOspfArea = &CommandSpec{
	[]Token{TNo, TIp, TPp, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelOspfArea = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TOspf, TArea}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

