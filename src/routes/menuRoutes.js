const express = require('express');
const router = express.Router();
const verifyToken = require('../middleware/auth');

router.get('/:restaurant_id/menus', verifyToken, (req, res) => {
  res.json({ status: 'success', message: 'Menu endpoint ready' });
});

module.exports = router;