package main

type SubAgent struct {
	// ContextName selects from SNMPV3 ContextName or SNMPV1/V2c community for switch from SubAgent...
	//             set to nil means all requests will gets here(of default)

	CommunityIDs []string

	// OIDs for Read/Write actions
	OIDs []*PDUValueControlItem

	// UserErrorMarkPacket decides if shll treat user returned error as generr
	UserErrorMarkPacket bool

	Logger ILogger

	master *MasterAgent
}
