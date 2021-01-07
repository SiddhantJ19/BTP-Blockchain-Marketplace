# BTP-Blockchain-Marketplace
To demonstrate Smart Access control on user/iot data, we have developed a marketplace on Hyperledger Fabric. The smart contracts (chaincode) control the access to all the data on the marketplace. The [application UI](https://github.com/adwait-thattey/btp-ui) and application-server are decoupled. App server communicates with Chaincode deployed on the peer nodes to interact with world state and blockchain ledger. 

We have used the private data store feature of Hyperledger Fabric to store device's private details. 
Complete Details, Flow and Screenshots can be found in the Docs folder.
