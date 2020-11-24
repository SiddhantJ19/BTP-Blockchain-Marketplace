const router = require('express').Router();
const {getOnSaleDevices, submitInterestToken, getInterestTokensForDevice} = require('./controllers')

router.post('/devices/onsale',getOnSaleDevices) // get all devices on sale
router.post('/devices/interesttokens/all',getInterestTokensForDevice) // submit a new interest token
router.post('/devices/interesttokens/submit',submitInterestToken) // submit a new interest token



module.exports = router