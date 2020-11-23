const router = require('express').Router();
const {registerDevice, updateDevice} = require('./controllers')

router.post('/register',registerDevice) // register a new device
router.post('/update',updateDevice) // update device details

/*
router.post('/update',updateDeviceData) // update existing device
router.post('/delete', deleteDeviceData) // delete a device
router.post('/sell', deleteDeviceData) // agree to sell
*/


module.exports = router