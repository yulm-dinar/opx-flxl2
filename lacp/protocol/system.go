// system
package lacp

import (
	"net"
)

// 6.4.5 Variables associated with the System
type LacpSystem struct {
	// System Priority
	actor_system_priority uint16
	// MAC address component of the System Id
	actor_system [6]uint8
}

func (s *LacpSystem) LacpSystemActorSystemIdSet(actor_system net.HardwareAddr) {
	s.actor_system = convertNetHwAddressToSysIdKey(actor_system)
}

func (s *LacpSystem) LacpSystemActorSystemPrioritySet(actor_system_priority uint16) {
	s.actor_system_priority = actor_system_priority
}

//6.3.2 System identification
//The globally unique identifier used to identify a System shall be the concatenation of a globally
//administered individual MAC address and the System Priority. The MAC address chosen may be the
//individual MAC address associated with one of the Aggregation Ports of the System. In the case of DRNI
//(Clause 9), all Portal Systems in a Portal have the same System Identifier, which is provided by the
//concatenation of the Portal’s administrated MAC address (7.4.1.1.4) and the Portal’s System Priority
//(7.4.1.1.5).
//
//Where it is necessary to perform numerical comparisons between System Identifiers, each System Identifier
//is considered to be an eight octet unsigned binary number, constructed as follows:
//
// a) The two most significant octets of the System Identifier comprise the System Priority. The System
//    Priority value is taken to be an unsigned binary number; the most significant octet of the System
//    Priority forms the most significant octet of the System Identifier.
//
// b) The third most significant octet of the System Identifier is derived from the initial octet of the MAC
//    address; the least significant bit of the octet is assigned the value of the first bit of the MAC address,
//    the next most significant bit of the octet is assigned the value of the next bit of the MAC address,
//    and so on. The fourth through eighth octets are similarly assigned the second through sixth octets of
//    the MAC address.
func LacpSystemIdGet(s LacpSystem) [8]uint8 {

	var sysId [8]uint8

	mac := s.actor_system

	sysId[7] = uint8(s.actor_system_priority >> 16 & 0xff)
	sysId[6] = uint8(s.actor_system_priority & 0xff)
	sysId[5] = mac[5]
	sysId[4] = mac[4]
	sysId[3] = mac[3]
	sysId[2] = mac[2]
	sysId[1] = mac[1]
	sysId[0] = mac[0]
	return sysId
}
