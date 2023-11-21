package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TicketContract contract for managing CRUD for Ticket
type TicketContract struct {
	contractapi.Contract
}

type PaginatedQueryResult struct {
	Records             []*Ticket `json:"records"`
	FetchedRecordsCount int32  `json:"fetchedRecordsCount"`
	Bookmark            string `json:"bookmark"`
}

type Ticket struct {
	AssetType         string `json:"AssetType"`
	TicketId             string `json:"TicketId"`
	EventId             string `json:"EventId"`
	EventDate string `json:"EventDate"`
	Price			  string `json:Price`
	OwnedBy           string `json:"OwnedBy"`
	Event              string `json:"Event"`
	Status            string `json:"Status"`
}

type Order struct { 
	AssetType  string `json:"assetType"`
	EventId      string `json:"eventid"`
	resellerName string `json:"resellerName"`
	price 	   string `json:price`
	Event       string `json:"event"`
	OrderID    string `json:"orderID"`
}

const collectionName string = "OrderCollection"

type HistoryQueryResult struct {
	Record    *Ticket   `json:"record"`
	TxId      string `json:"txId"`
	Timestamp string `json:"timestamp"`
	IsDelete  bool   `json:"isDelete"`
}


// TicketExists returns true when asset with given ID exists in world state
func (c *TicketContract) TicketExists(ctx contractapi.TransactionContextInterface, ticketID string) (bool, error) {
	data, err := ctx.GetStub().GetState(ticketID)

	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)

	}
	return data != nil, nil
}

// CreateTicket creates a new instance of Ticket
func (c *TicketContract) CreateTicket(ctx contractapi.TransactionContextInterface, ticketID string, event string, eventid string, organiserName string, eventDate string, price string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "organiserMSP" {

		exists, err := c.TicketExists(ctx, ticketID)
		if err != nil {
			return "", fmt.Errorf("could not read from world state. %s", err)
		} else if exists {
			return "", fmt.Errorf("the asset %s already exists", ticketID)
		}

		ticket := Ticket{
			AssetType:         "ticket",
			TicketId:             ticketID,
			EventId:             eventid,
			EventDate:         eventDate,
			Price:			   price,
			Event:              event,
			OwnedBy:           organiserName,
			Status:            "In Factory",
		}

		fmt.Println("Create ticket data =====",ticket)

		bytes, _ := json.Marshal(ticket)

		err = ctx.GetStub().PutState(ticketID, bytes)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("successfully added ticket %v", ticketID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}

}

// ReadTicket retrieves an instance of Ticket from the world state
func (c *TicketContract) ReadTicket(ctx contractapi.TransactionContextInterface, ticketID string) (*Ticket, error) {
	exists, err := c.TicketExists(ctx, ticketID)
	if err != nil {
		return nil, fmt.Errorf("could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("the asset %s does not exist", ticketID)
	}

	bytes, _ := ctx.GetStub().GetState(ticketID)

	ticket := new(Ticket)

	err = json.Unmarshal(bytes, &ticket)

	if err != nil {
		return nil, fmt.Errorf("could not unmarshal world state data to type Ticket")
	}

	return ticket, nil
}

// DeleteTicket removes the instance of Ticket from the world state
func (c *TicketContract) DeleteTicket(ctx contractapi.TransactionContextInterface, ticketID string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "organiserMSP" {

		exists, err := c.TicketExists(ctx, ticketID)
		if err != nil {
			return "", fmt.Errorf("could not read from world state. %s", err)
		} else if !exists {
			return "", fmt.Errorf("the asset %s does not exist", ticketID)
		}

		err = ctx.GetStub().DelState(ticketID)
		if err != nil {
			return "", err
		} else {
			return fmt.Sprintf("ticket with id %v is deleted from the world state.", ticketID), nil
		}

	} else {
		return "", fmt.Errorf("user under following MSP:%v cannot able to perform this action", clientOrgID)
	}
}

func (c *TicketContract) GetTicketsByRange(ctx contractapi.TransactionContextInterface, startKey, endKey string) ([]*Ticket, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close() // check usage

	return ticketResultIteratorFunction(resultsIterator)
}

func (c *TicketContract) GetTicketHistory(ctx contractapi.TransactionContextInterface, ticketID string) ([]*HistoryQueryResult, error) {

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(ticketID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close() // check for details why do we need to close it

	var records []*HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var ticket Ticket
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &ticket)
			if err != nil {
				return nil, err
			}
		} else {
			ticket = Ticket{
				TicketId: ticketID,
			}
		}

		timestamp := response.Timestamp.AsTime()

		formattedTime := timestamp.Format(time.RFC1123)

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &ticket,
			IsDelete:  response.IsDelete,
		}
		records = append(records, &record)
	}

	return records, nil
}

func (c *TicketContract) GetAllTickets(ctx contractapi.TransactionContextInterface) ([]*Ticket, error) {

	queryString := `{"selector":{"AssetType":"ticket"}}`

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	return ticketResultIteratorFunction(resultsIterator)
}

func (c *TicketContract) GetTicketsWithPagination(ctx contractapi.TransactionContextInterface, pageSize int32, bookmark string) (*PaginatedQueryResult, error) {
	queryString := `{"selector":{"AssetType":"ticket"}}`
	resultsIterator, responseMetadata, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	tickets, err := ticketResultIteratorFunction(resultsIterator)
	if err != nil {
		return nil, err
	}

	return &PaginatedQueryResult{
		Records:             tickets,
		FetchedRecordsCount: responseMetadata.FetchedRecordsCount,
		Bookmark:            responseMetadata.Bookmark,
	}, nil
}


func ticketResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Ticket, error) {
	var tickets []*Ticket
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var ticket Ticket
		err = json.Unmarshal(queryResult.Value, &ticket)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, &ticket)
	}

	return tickets, nil
}

func (c *TicketContract) SellTicket(ctx contractapi.TransactionContextInterface, ticketID string, ownerName string, phoneNumber string, price string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "resellerMSP" {

		exists, err := c.TicketExists(ctx, ticketID)
		if err != nil {
			return "", fmt.Errorf("Could not read from world state. %s", err)
		}
		if exists {
			ticket, _ := c.ReadTicket(ctx, ticketID)
			ticket.Status = fmt.Sprintf("Sold to  %v with phone number %v for Rs. %v", ownerName, phoneNumber, price)
			ticket.OwnedBy = ownerName

			bytes, _ := json.Marshal(ticket)
			err = ctx.GetStub().PutState(ticketID, bytes)
			if err != nil {
				return "", err
			} else {
				return fmt.Sprintf("Ticket %v successfully sold to %v", ticketID, ownerName), nil
			}

		} else {
			return "", fmt.Errorf("Ticket %v does not exist!", ticketID)
		}

	} else {
		return "", fmt.Errorf("User under following MSPID: %v cannot able to perform this action", clientOrgID)
	}

}

func (c *TicketContract) TransferTicket(ctx contractapi.TransactionContextInterface, ticketID string, ownerName string, phoneNumber string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "attendeeMSP" {

		exists, err := c.TicketExists(ctx, ticketID)
		if err != nil {
			return "", fmt.Errorf("Could not read from world state. %s", err)
		}
		if exists {
			ticket, _ := c.ReadTicket(ctx, ticketID)
			ticket.Status = fmt.Sprintf("Transferred to  %v with phone number %v", ownerName, phoneNumber)
			ticket.OwnedBy = ownerName

			bytes, _ := json.Marshal(ticket)
			err = ctx.GetStub().PutState(ticketID, bytes)
			if err != nil {
				return "", err
			} else {
				return fmt.Sprintf("Ticket %v successfully transferred to %v", ticketID, ownerName), nil
			}

		} else {
			return "", fmt.Errorf("Ticket %v does not exist!", ticketID)
		}

	} else {
		return "", fmt.Errorf("User under following MSPID: %v cannot able to perform this action", clientOrgID)
	}

}

// OrderExists returns true when asset with given ID exists in private data collection
func (o *TicketContract) OrderExists(ctx contractapi.TransactionContextInterface, orderID string) (bool, error) {

	data, err := ctx.GetStub().GetPrivateDataHash(collectionName, orderID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateOrder creates a new instance of Order
func (o *TicketContract) CreateOrder(ctx contractapi.TransactionContextInterface, orderID string) (string, error) {

	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}
	if (clientOrgID == "resellerMSP" || clientOrgID=="organiserMSP") {
		exists, err := o.OrderExists(ctx, orderID)
		if err != nil {
			return "", fmt.Errorf("could not read from world state. %s", err)
		} else if exists {
			return "", fmt.Errorf("the asset %s already exists", orderID)
		}

		order := new(Order)

		transientData, err := ctx.GetStub().GetTransient()
		if err != nil {
			return "", err
		}

		if len(transientData) == 0 {
			return "", fmt.Errorf("Please provide the private data of event, eventid, resellerName")
		}

		event, exists := transientData["event"]
		if !exists {
			return "", fmt.Errorf("The event was not specified in transient data. Please try again")
		}
		order.Event = string(event)

		// model, exists := transientData["model"]
		// if !exists {
		// 	return "", fmt.Errorf("The model was not specified in transient data. Please try again")
		// }
		// order.Model = string(model)

		price, exists := transientData["price"]
		if !exists {
			return "", fmt.Errorf("The price was not specified in transient data. Please try again")
		}
		order.price = string(price)

		eventid, exists := transientData["eventid"]
		if !exists {
			return "", fmt.Errorf("The eventid was not specified in transient data. Please try again")
		}
		order.EventId = string(eventid)

		resellerName, exists := transientData["resellerName"]
		if !exists {
			return "", fmt.Errorf("the reseller was not specified in transient data. Please try again")
		}
		order.resellerName = string(resellerName)

		order.AssetType = "Order"
		order.OrderID = orderID

		bytes, _ := json.Marshal(order)
		err = ctx.GetStub().PutPrivateData(collectionName, orderID, bytes)
		if(err != nil) {
			return "", fmt.Errorf("could not able to write the data")
		}
		return fmt.Sprintf("Order with id %v added successfully", orderID), nil
	} else {
		return fmt.Sprintf("Order cannot be created by organisation with MSPID %v ", clientOrgID), nil
	}
}

// ReadOrder retrieves an instance of Order from the private data collection
func (o *TicketContract) ReadOrder(ctx contractapi.TransactionContextInterface, orderID string) (*Order, error) {
	exists, err := o.OrderExists(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", orderID)
	}

	bytes, err := ctx.GetStub().GetPrivateData(collectionName, orderID)
	if err != nil {
		return nil, err
	}
	order := new(Order)

	err = json.Unmarshal(bytes, order)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal private data collection data to type Order")
	}

	return order, nil

}

// DeleteOrder deletes an instance of Order from the private data collection
func (o *TicketContract) DeleteOrder(ctx contractapi.TransactionContextInterface, orderID string) error {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}
	if clientOrgID == "resellerMSP" {
		exists, err := o.OrderExists(ctx, orderID)

		if err != nil {
			return fmt.Errorf("could not read from world state. %s", err)
		} else if !exists {
			return fmt.Errorf("the asset %s does not exist", orderID)
		}

		return ctx.GetStub().DelPrivateData(collectionName, orderID)
	} else {
		return fmt.Errorf("organisation with %v cannot delete the order", clientOrgID)
	}
}

func OrderResultIteratorFunction(resultsIterator shim.StateQueryIteratorInterface) ([]*Order, error) {
	var orders []*Order
	for resultsIterator.HasNext() {
		queryResult, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var order Order
		err = json.Unmarshal(queryResult.Value, &order)
		if err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}









