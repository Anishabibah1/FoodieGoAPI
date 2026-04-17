const express = require('express');
const router = express.Router();
const verifyToken = require('../middleware/auth');
const { getAllRestaurants, getRestaurantById, searchRestaurants } = require('../controllers/restaurantController');

router.get('/', verifyToken, getAllRestaurants);
router.get('/search', verifyToken, searchRestaurants);
router.get('/:id', verifyToken, getRestaurantById);

module.exports = router;