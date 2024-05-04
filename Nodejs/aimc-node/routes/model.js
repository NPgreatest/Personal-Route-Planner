// 使用ESM导入语法
import express from 'express';
const router = express.Router();
const app = express();
app.use(express.json());
import { App, Model,  User, File, KnowledgeBase } from '../db.js';



router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello",  (req, res) => {
    console.log(req.body);
    res.send('ok');

});

router.post('/', async (req, res) => {

    // const model = await Model.findOne({ name: req.body.name });
    // if (model) {
    //     res.status(400).json({ message: 'Model name already exists' });
    // } else {
        const newModel = new Model(req.body);
        // try {
            const savedModel = await newModel.save();
            res.json({ message: 'The model entry has been created successfully', data: savedModel });
        // } catch (err) {
        //     res.status(500).json({ message: 'An error occurred while creating the model entry', error: err });
        // }
    // }
});


router.get('/', async (req, res) => {
    try {
        const models = await Model.find();
        res.json(models);
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the model entries', error: err });
    }
});

router.get('/name/:name', async (req, res) => {
    try {
        const model = await Model.findOne({ name: req.params.name });
        if (model) {
            res.json(model);
        } else {
            res.status(404).json({ message: 'No model found with the provided name' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the model entry', error: err });
    }
});

router.get('/id/:id', async (req, res) => {
    try {
        const model = await Model.findById( req.params.id );
        if (model) {
            res.json(model);
        } else {
            res.status(404).json({ message: 'No model found ' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the model entry', error: err });
    }
});




export default router;
