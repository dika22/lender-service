package constant

type CtxKey string

const(
	NewRelicTransactionCtx CtxKey = "newRelicTransaction"

	Proposed  = 1
	Approved  = 2
	Invested  = 3
	Disbursed = 4

	StateProposed  = "Proposed"
	StateApproved  = "Approved"
	StateInvested  = "Invested"
	StateDisbursed = "Disbursed"

	QueueHigh   = "lender:queue:high"
	QueueMedium = "lender:queue:medium"
	QueueLow    =  "lender:queue:low"

)