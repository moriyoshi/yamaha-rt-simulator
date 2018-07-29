package simulator

import "context"

var CmdIpInterfaceIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TIp, TL2Interface, TIntrusion, TDetection, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TIp, TPp, TIntrusion, TDetection, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TIp, TTunnel, TIntrusion, TDetection, TThreshold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIntrusion, TDetection, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIntrusion, TDetection, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIntrusionDetectionThreshold = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIntrusion, TDetection, TThreshold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

