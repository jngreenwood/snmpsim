package main

import "github.com/gosnmp/gosnmp"

type PDUValueControlItem struct {
	// OID controls which OID does this PDUValue works
	OID string
	// Type defines which type this OID is.
	Type gosnmp.Asn1BER

	// NonWalkable marks this oid as not walkable. It **WILL NOT** returned in walk items. but do retuend
	//             in direct get.
	//             All write only item will be NonWalkable
	NonWalkable bool

	/////////// Callbacks

	// OnCheckPermission will be called on access this OID. set to nil to allow all access.
	//     return PermissionAllowanceAllowed for allow this access.
	//            (otherwrise) PermissionAllowanceDenied for disable access.
	// OnCheckPermission FuncPDUControlCheckPermission

	// OnGet will be called on any GET / walk option. set to nil for mark this as a write-only item
	// OnGet FuncPDUControlGet
	// // OnSet will be called on any Set option. set to nil for mark as a read-only item.
	// OnSet FuncPDUControlSet
	// // OnTrap will be called on TRAP.
	// OnTrap FuncPDUControlTrap
	//////////// For human document

	//Document for this PDU Item. ignored by the program.
	Document string
}
