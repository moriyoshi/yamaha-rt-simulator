package simulator

import "context"

var CmdIpInterfaceInboundFilterList = &CommandSpec{
	[]Token{TIp, TL2Interface, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6InterfaceInboundFilterList = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpInboundFilterList = &CommandSpec{
	[]Token{TIp, TPp, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpInboundFilterList = &CommandSpec{
	[]Token{TIpv6, TPp, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelInboundFilterList = &CommandSpec{
	[]Token{TIp, TTunnel, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6TunnelInboundFilterList = &CommandSpec{
	[]Token{TIpv6, TTunnel, TInbound, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceInboundFilterList = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceInboundFilterList = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpInboundFilterListId = &CommandSpec{
	[]Token{TNo, TIp, TPp, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpInboundFilterListId = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelInboundFilterListId = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6TunnelInboundFilterListId = &CommandSpec{
	[]Token{TNo, TIpv6, TTunnel, TInbound, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

