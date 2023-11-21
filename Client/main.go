package main

import "fmt"

func main() {

	// use this functions to evaluate and submit txns
	// try calling these functions

	result := submitTxnFn(
		"organiser",
		"ticketchannel",
		"Final-Project",
		"TicketContract",
		"invoke",
		make(map[string][]byte),
		"CreateTicket",
		"Ticket-09",
		"Garba",
		"Event-05",
		"CollegeClub",
		"25/10/2023",
		"750.00",
	)

	// privateData := map[string][]byte{
	// 	"event":       []byte("Concert"),
	// 	"eventid":      []byte("Event-06"),
	// 	"resellerName": []byte("Sun"),
	// }

	// result := submitTxnFn("reseller", "ticketchannel", "Final-Project", "OrderContract", "private", privateData, "CreateOrder", "ORD-03")
	// result := submitTxnFn("reseller", "ticketchannel", "Final-Project", "OrderContract", "query", event(map[string][]byte), "ReadOrder", "ORD-03")
	// result := submitTxnFn("organiser", "ticketchannel", "Final-Project", "TicketContract", "query", event(map[string][]byte), "GetAllTickets")
	// result := submitTxnFn("organiser", "ticketchannel", "Final-Project", "OrderContract", "query", event(map[string][]byte), "GetAllOrders")
	// result := submitTxnFn("organiser", "ticketchannel", "Final-Project", "TicketContract", "query", event(map[string][]byte), "GetMatchingOrders", "Ticket-06")
	// result := submitTxnFn("organiser", "ticketchannel", "Final-Project", "TicketContract", "invoke", event(map[string][]byte), "MatchOrder", "Ticket-06", "ORD-01")
	// result := submitTxnFn("attendee", "ticketchannel", "Final-Project", "TicketContract", "invoke", event(map[string][]byte), "RegisterTicket", "Ticket-06", "Dani", "+91 9999999999")
	// fmt.Println(result)
}
