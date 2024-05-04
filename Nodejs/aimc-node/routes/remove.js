import express from 'express';
import fs from 'fs';
import multer from 'multer';
import path from 'path';
import { App, Model,  User, File, KnowledgeBase } from '../db.js';

const app = express();
const router = express.Router();

app.use(express.json());

var uploadFolder = '../upload/';

router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello", (req, res) => {
    console.log(req.body);
    res.send('ok');

});

router.delete('/:id', async (req, res) => {

    try {
        const file = await File.findById(req.params.id);
        const knowledgeBase = await KnowledgeBase.findById(file.knowledgebase);
        knowledgeBase.file.pull(file._id);
        await knowledgeBase.save();
        const deletedFile = await File.findByIdAndDelete(req.params.id);
        fs.unlink(file.path, (err) => {
            if (err) {
                console.error(err);
                return;
            }
            console.log('File deleted successfully');
        });
            // 获取文件夹的路径，假设为文件名加上"_faiss_index"后缀
            const folder = file.path + "_faiss_index";
            // 使用fs模块的rmSync方法，删除文件夹及其内容
            fs.rmSync(folder, { recursive: true, force: true });
            console.log('Folder deleted successfully');
            if (deletedFile) {
                res.json({ message: 'The file entry and the folder have been deleted successfully', code: 1 });
            } else {
                res.status(404).json({ message: 'No file found ' });
            }
        } catch (err) {
            res.status(500).json({ message: 'An error occurred while deleting the file entry and the folder', error: err });
        }


    });


export default router;

