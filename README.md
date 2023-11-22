# Event Ticketing System

# Introduction

The Event Ticketing System is a blockchain-based application built on Hyperledger Fabric, providing a secure and transparent platform for managing and facilitating event tickets. The system involves three key participants: organizers, resellers, and attendees. Using smart contracts, the system supports functionalities such as creating tickets, transferring and selling tickets, managing orders, and more.

# Installation

Clone the repository: git clone https://github.com/Samridhi102/Final-Project.git cd Final-Project/ ./startNetwork.sh

# Getting Started

# Network Setup

Navigate to the Fabric-Network folder.

Execute docker-compose -f docker/docker-compose-ca.yaml up -d to start the Certificate Authorities.

Run the registerEnroll.sh script to register and enroll users for each organization.

Build the infrastructure by executing docker-compose -f docker/docker-compose-3org.yaml up -d.

Generate the genesis block with configtxgen -profile ThreeOrgsChannel -outputBlock ./channel-artifacts/ticketchannel.block -channelID ticketchannel.

Create the application channel using osnadmin channel join ... commands for each peer.

Join peers to the channel using peer channel join -b ./channel-artifacts/ticketchannel.block.

Anchor peer updates ensure proper communication between peers within organizations.

Commands are in commands.txt file

# Smart Contracts

The ticket-contract.go file contains the smart contract code defining the behavior of the ticketing system. It includes functions for creating, transferring, selling, and deleting tickets, as well as managing orders.

# Chaincode Lifecycle

Build and install the chaincode on each peer.

Approve the chaincode for each organization.

Check the commit readiness of the chaincode.

Commit the chaincode to the channel.

Invoke chaincode functions using peer chaincode invoke commands.






Step by step commands are written in commands.txt file for manually starting the network and commands(running_throught_script).txt for starting the network from script file.


