package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in MTIP
	DefaultTransferFee float64 = 0.1
	// DefaultTransferMixin is the default mixin
	DefaultTransferMixin = 7

	logWalletdCurrentSessionFilename     = "turtle-service-session.log"
	logWalletdAllSessionsFilename        = "turtle-service.log"
	logTurtleCoindCurrentSessionFilename = "TurtleCoind-session.log"
	logTurtleCoindAllSessionsFilename    = "TurtleCoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "walletd"
	turtlecoindCommandName               = "monkeytipsd"
)
