package client

type UserSpaceSt struct {
	TotalSpace     string
	UsedSpace      string
	LockedSpace    string
	RemainingSpace string
	State          string
	Start          uint32
	Deadline       uint32
}

type ChallengeInfo struct {
	Random []byte
	Start  uint32
}

type ChallengeSnapshot struct {
	NetSnapshot   NetSnapshot
	MinerSnapshot []MinerSnapshot
}

type NetSnapshot struct {
	Start               uint32
	Total_reward        string
	Total_idle_space    string
	Total_service_space string
	Random              []byte
}

type MinerSnapshot struct {
	Miner         string
	Idle_space    string
	Service_space string
}

type TeeWorkerSt struct {
	Controller_account string
	Peer_id            []byte
	Node_key           []byte
	Stash_account      string
}
