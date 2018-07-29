package simulator

import "context"

var CmdIpInterfaceIntrusionDetection = &CommandSpec{
	[]Token{TIp, TL2Interface, TIntrusion, TDetection}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIntrusionDetection = &CommandSpec{
	[]Token{TIp, TPp, TIntrusion, TDetection}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdIpTunnelIntrusionDetection = &CommandSpec{
	[]Token{TIp, TTunnel, TIntrusion, TDetection}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpInterfaceIntrusionDetection = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIntrusion, TDetection}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIntrusionDetection = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIntrusion, TDetection}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpTunnelIntrusionDetection = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIntrusion, TDetection}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
