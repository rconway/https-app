import express from 'express';
const router = express.Router();
export default router;

/* GET time. */
router.get('/', function(req, res, next) {
  let time = {
    time: new Date().toUTCString()
  }
  res.send(time);
});
