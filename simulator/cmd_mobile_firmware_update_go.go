package simulator

import "context"

var CmdMobileFirmwareUpdateGo = &CommandSpec{
	[]Token{TMobile, TFirmware, TUpdate, TGo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

