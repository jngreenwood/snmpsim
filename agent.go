package main

import "github.com/gosnmp/gosnmp"

type PrimaryAgent struct {
	SecurityConfig SecurityConfig

	SubAgents []*SubAgent

	// Logger ILogger

	priv struct {
		communityToSubAgent map[string]*SubAgent
		defaultSubAgent     *SubAgent
	}
}

type SecurityConfig struct {
	NoSecurity bool

	// AuthoritativeEngineID is SNMPV3 AuthoritativeEngineID
	AuthoritativeEngineID SNMPEngineID
	// AuthoritativeEngineBoots is SNMPV3 AuthoritativeEngineBoots
	AuthoritativeEngineBoots uint32
	// OnGetAuthoritativeEngineTime will be called to get SNMPV3 AuthoritativeEngineTime
	//      if sets to nil, the sys boottime will be used
	OnGetAuthoritativeEngineTime FuncGetAuthoritativeEngineTime

	Users []gosnmp.UsmSecurityParameters
}

type SNMPEngineID struct {
	// See https://tools.ietf.org/html/rfc3411#section-5
	// 			SnmpEngineID ::= TEXTUAL-CONVENTION
	//      SYNTAX       OCTET STRING (SIZE(5..32))
	EngineIDData string
}
