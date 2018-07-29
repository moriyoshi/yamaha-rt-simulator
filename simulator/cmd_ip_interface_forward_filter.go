package simulator

import "context"

var CmdIpInterfaceForwardFilter = &CommandSpec{
	[]Token{TIp, TL2Interface, TForward, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpForwardFilter = &CommandSpec{
	[]Token{TIp, TPp, TForward, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelForwardFilter = &CommandSpec{
	[]Token{TIp, TTunnel, TForward, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpLocalForwardFilter = &CommandSpec{
	[]Token{TIp, TLocal, TForward, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceForwardFilter = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TForward, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpForwardFilter = &CommandSpec{
	[]Token{TNo, TIp, TPp, TForward, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelForwardFilter = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TForward, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpLocalForwardFilter = &CommandSpec{
	[]Token{TNo, TIp, TLocal, TForward, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

