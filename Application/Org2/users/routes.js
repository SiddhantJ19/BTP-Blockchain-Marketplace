const router = require('express').Router();
const {registerUser, enrollAdmin, connectGateway} = require('./controllers')

router.post('/admin/enroll',enrollAdmin)
router.post('/register', registerUser)
router.post('/connect', connectGateway)


module.exports = router