package simulator

import "context"

var CmdIpInterfaceSecureFilter = &CommandSpec{
	[]Token{TIp, TL2Interface, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpSecureFilter = &CommandSpec{
	[]Token{TIp, TPp, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelSecureFilter = &CommandSpec{
	[]Token{TIp, TTunnel, TSecure, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpInterfaceSecureFilterName = &CommandSpec{
	[]Token{TIp, TL2Interface, TSecure, TFilter, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpSecureFilterName = &CommandSpec{
	[]Token{TIp, TPp, TSecure, TFilter, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelSecureFilterName = &CommandSpec{
	[]Token{TIp, TTunnel, TSecure, TFilter, TName}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceSecureFilter = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpSecureFilter = &CommandSpec{
	[]Token{TNo, TIp, TPp, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelSecureFilter = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TSecure, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceSecureFilterName = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TSecure, TFilter, TName}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpSecureFilterName = &CommandSpec{
	[]Token{TNo, TIp, TPp, TSecure, TFilter, TName}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelSecureFilterName = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TSecure, TFilter, TName}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

