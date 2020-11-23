const router = require('express').Router();
const {getOnSaleDevices} = require('./controllers')

router.post('/devices/onsale',getOnSaleDevices) // get all devices on sale



module.exports = router