package collect


type AgentAct interface{
	Collect()
	GetReportStr()string
}
