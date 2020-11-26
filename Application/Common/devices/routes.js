const router = require('express').Router();
const {registerDevice, updateDevice, agreeToSell, getDeviceDetails, deleteDevice, getDeviceAllData, getDeviceLatestData, newData} = require('./controllers')

router.post('/register',registerDevice) // register a new device
router.post('/update',updateDevice) // update device details
router.post('/delete',deleteDevice) // update device details

router.post('/data/latest',getDeviceLatestData) // update device details
router.post('/data/all',getDeviceAllData) // update device details
router.post('/data/add',newData) // update device details

router.post('/agreetosell',agreeToSell) //
router.post('/',getDeviceDetails) // get device details

/*
router.post('/update',updateDeviceData) // update existing device
router.post('/delete', deleteDeviceData) // delete a device
router.post('/sell', deleteDeviceData) // agree to sell
*/


module.exports = router