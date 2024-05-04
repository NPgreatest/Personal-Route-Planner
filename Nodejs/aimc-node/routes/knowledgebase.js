import express from 'express';
import fs from 'fs'; // 使用ESM导入文件系统模块
import path from 'path'; // 使用ESM导入路径处理模块
import { App, Model,  User, File, KnowledgeBase } from '../db.js'; // 改为使用相对路径和.js扩展名

const router = express.Router();
const app = express();

app.use(express.json());







router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello",  (req, res) => {
    console.log(req.body);
    res.send('ok');

});

router.get('/:id', async (req, res) => {
    try {

        const knowledgebase = await KnowledgeBase.findById( req.params.id );
        const file = await File.find({knowledgebase:req.params.id});
        if (knowledgebase) {
            res.json(knowledgebase,file);
        } else {
            res.status(404).json({ message: 'No app found ' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the app entry', error: err });
    }
});


router.post('/', async (req, res) => {

    const knowledgebase = await KnowledgeBase.findOne({ name: req.body.name });
    if (knowledgebase) {
        res.status(400).json({ message: 'KnowledgeBase name already exists' });
    } else {
        const newKnowledgeBase = new KnowledgeBase(req.body);
        try {
            const savedKnowledgeBase = await newKnowledgeBase.save();
            res.json({ message: 'The knowledgebase entry has been created successfully', data: savedKnowledgeBase });
        } catch (err) {
            res.status(500).json({ message: 'An error occurred while creating the knowledgebase entry', error: err });
        }
    }
    console.log("get knowledge success")
});
// create knowledgebase

router.get('/', async (req, res) => {
       try {
        const knowledgebases = await KnowledgeBase.find();
        console.log("get:knowledgebases");
        res.json(knowledgebases);
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the app entries', error: err });
    }

});

router.put('/:id', async (req, res) => {
    try {
        const updatedKnowledgeBase = await KnowledgeBase.findByIdAndUpdate(req.params.id, req.body, { new: true });
        if (updatedKnowledgeBase) {
            res.json({ message: 'The knowledgebase entry has been updated successfully', data: updatedKnowledgeBase });
        } else {
            res.status(404).json({ message: 'No knowledgebase found with the provided id' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while updating the knowledgebase entry', error: err });
    }
});

router.delete('/:id', async (req, res) => {
    try {
        const deletedKnowledgebase = await KnowledgeBase.findByIdAndDelete(req.params.id);
        const deletedApp = await App.updateMany({ $pull: { knowledgebase: req.params.id } });
        const files = await File.find({ knowledgebase: req.params.id });
        const deletedFiles = await File.deleteMany({ knowledgebase: req.params.id });
        files.forEach(file => {
            fs.unlinkSync(file.path);
            fs.rmdirSync(`${file.path}_faiss_index`, { recursive: true });
            //const folder = file.path + "_faiss_index";
            // 使用fs模块的rmSync方法，删除文件夹及其内容
            //fs.rmSync(folder, { recursive: true, force: true });
        });
        res.status(200).send({ deletedKnowledgebase, deletedApp, deletedFiles });
    } catch (err) {
        console.log(err);
        res.status(500).send(err);

    }


});

export default router;
