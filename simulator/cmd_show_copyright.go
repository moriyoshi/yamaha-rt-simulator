package simulator

import "context"

var CmdShowCopyright = &CommandSpec{
	[]Token{TShow, TCopyright}, 2,
	func(_ context.Context, sess *SimulatorSession, _ []TokenInstance) error {
		sess.WriteF("%s\n", sess.Firmware.Copyright)
		return nil
	},
}
