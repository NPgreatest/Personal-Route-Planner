import express from 'express';
const router = express.Router();
const app = express();
app.use(express.json());
import { App, Model,  User, File, KnowledgeBase } from '../db.js';



router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello",  (req, res) => {
    console.log("Hello World!!!");
    res.send('ok');

});

router.post('/', async (req, res) => {

    const file = await File.findOne({ name: req.body.name });
    if (file) {
        res.status(400).json({ message: 'File name already exists' });
    } else {
        const newFile = new File(req.body);
        try {
            const savedFile = await newFile.save();
            res.json({ message: 'The file entry has been created successfully', data: savedFile });
        } catch (err) {
            res.status(500).json({ message: 'An error occurred while creating the file entry', error: err });
        }
    }
});

router.get('/', async (req, res) => {
    try {
        const files = await File.find();
        res.json(files);
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the file entries', error: err });
    }
});



router.get('/kbid/:id', async (req, res) => {
    try {
        const files = await File.find({knowledgebase:req.params.id});
        res.json(files);

    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the file entries'});
        console.log(err);
    }
});

router.put('/:id', async (req, res) => {
    try {
        const body=req.body;
        const updatedFile = await File.findByIdAndUpdate(req.params.id, req.body, { new: true });
        console.log("put:file/:id")
        if (updatedFile) {
            res.json({ message: 'The file entry has been updated successfully', data: updatedFile });
        } else {
            res.status(404).json({ message: 'No file found with the provided name' });
        }
    } catch (err) {

        res.status(500).json({ message: 'An error occurred while updating the file entry', error: err });
    }
});


export default router;
