/* eslint-disable strict */
const { Gateway, Wallets } = require('fabric-network');
const FabricCAServices = require('fabric-ca-client');
const path = require('path');
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require('./base/CAUtil');
const { buildCCPOrg2, buildWallet } = require('./base/AppUtil');

const express = require('express');
const dotenv = require('dotenv');
const config = require('./config/base');
const cors = require('cors')

const app = express();
// const uiroutes = require('./ui/routes')
const userRoutes = require('./users/routes');
const deviceRoutes = require('./devices/routes');
const marketRoutes = require('./marketplace/routes');

const channelName = config.channelName;
const chaincodeName = config.chaincodeName;
const mspOrg1 = 'Org1MSP';

const walletPath = path.join(__dirname, 'wallet');
config.walletPath = walletPath

config.ccp = buildCCPOrg2();
config.caClient = buildCAClient(FabricCAServices, config.ccp, 'ca.org2.example.com');

buildWallet(Wallets, walletPath).then( wallet => {
    config.wallet = wallet;
    console.log('Wallet Ready');
});


app.use(cors())
app.use(express.json());


const port = 4000;

app.use('/users', userRoutes);

app.use('/devices', deviceRoutes);
app.use('/market', marketRoutes);


// app.use('/ui', uiroutes)

app.listen(port, () => console.log(`Org1 Server running on ${port}`));
