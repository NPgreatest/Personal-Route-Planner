import express from 'express';
const router = express.Router();
const app = express();
app.use(express.json());
import { App, Model, User, File, KnowledgeBase } from '../../db.js';


router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello", (req, res) => {
    console.log(req.body);
    res.send('ok');
});


router.post('/', async (req, res) => {
    const newUser = new User(req.body);
    try {
        const savedUser = await newUser.save();
        res.json(savedUser);
    } catch (err) {
        res.status(500).json(err);
    }
});


router.get('/', async (req, res) => {
    try {
        const users = await User.find();
        res.json(users);
    } catch (err) {
        res.status(500).json(err);
    }
});


router.put('/id/:id', async (req, res) => {
    try {
        const updatedUser = await User.findByIdAndUpdate(req.params.id, req.body, { new: true });
        res.json(updatedUser);
    } catch (err) {
        res.status(500).json(err);
    }
});


router.delete('/id/:id', async (req, res) => {
    try {
        await User.findByIdAndDelete(req.params.id);
        res.json('The user entry has been deleted');
    } catch (err) {
        res.status(500).json(err);
    }
});

export default router;

