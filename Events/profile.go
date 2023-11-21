package main

// Config represents the configuration for a role.
type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}

// Create a Profile map
var profile = map[string]Config{

	"organiser": {
		CertPath:     "../Fabric-Network/organizations/peerOrganizations/organiser.ticket.com/users/User1@organiser.ticket.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric-Network/organizations/peerOrganizations/organiser.ticket.com/users/User1@organiser.ticket.com/msp/keystore/",
		TLSCertPath:  "../Fabric-Network/organizations/peerOrganizations/organiser.ticket.com/peers/peer0.organiser.ticket.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.organiser.ticket.com",
		MSPID:        "organiserMSP",
	},

	"reseller": {
		CertPath:     "../Fabric-Network/organizations/peerOrganizations/reseller.ticket.com/users/User1@reseller.ticket.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric-Network/organizations/peerOrganizations/reseller.ticket.com/users/User1@reseller.ticket.com/msp/keystore/",
		TLSCertPath:  "../Fabric-Network/organizations/peerOrganizations/reseller.ticket.com/peers/peer0.reseller.ticket.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.reseller.ticket.com",
		MSPID:        "resellerMSP",
	},

	"attendee": {
		CertPath:     "../Fabric-Network/organizations/peerOrganizations/attendee.ticket.com/users/User1@attendee.ticket.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Fabric-Network/organizations/peerOrganizations/attendee.ticket.com/users/User1@attendee.ticket.com/msp/keystore/",
		TLSCertPath:  "../Fabric-Network/organizations/peerOrganizations/attendee.ticket.com/peers/peer0.attendee.ticket.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.attendee.ticket.com",
		MSPID:        "attendeeMSP",
	},
}