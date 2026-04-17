const db = require('../config/db');

const getAllRestaurants = async (req, res) => {
  try {
    const [rows] = await db.query('SELECT * FROM restaurants');
    res.json({ status: 'success', data: rows });
  } catch (err) {
    res.status(500).json({ status: 'error', message: err.message });
  }
};

const getRestaurantById = async (req, res) => {
  try {
    const [rows] = await db.query('SELECT * FROM restaurants WHERE id = ?', [req.params.id]);
    if (rows.length === 0) {
      return res.status(404).json({ status: 'error', message: 'Restaurant not found.' });
    }
    res.json({ status: 'success', data: rows[0] });
  } catch (err) {
    res.status(500).json({ status: 'error', message: err.message });
  }
};

const searchRestaurants = async (req, res) => {
  try {
    const { q } = req.query;
    if (!q) {
      return res.status(400).json({ status: 'error', message: 'Query parameter q is required.' });
    }
    const [rows] = await db.query(
      'SELECT * FROM restaurants WHERE name LIKE ? OR category LIKE ?',
      [`%${q}%`, `%${q}%`]
    );
    res.json({ status: 'success', data: rows });
  } catch (err) {
    res.status(500).json({ status: 'error', message: err.message });
  }
};

module.exports = { getAllRestaurants, getRestaurantById, searchRestaurants };