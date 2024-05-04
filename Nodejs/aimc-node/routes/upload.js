import express from 'express';
import net from 'net';
import fs from 'fs';
import multer from 'multer';
import { App, Model, User, File, KnowledgeBase } from '../db.js'; // 确保db模块使用ESM方式导出

const app = express();
const router = express.Router();

app.use(express.json());


import { EmbeddingFile }  from '../js_script/embedding.mjs';

// const spawn = require('child_process').spawn
// const { exec } = require('child_process');
// const { type } = require('os');
// 沟通python导入的文件


router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello", (req, res) => {
    console.log(req.body);
    res.send('ok');

});

var createFolder = function (folder) {
    try {
        fs.accessSync(folder);
    } catch (e) {
        fs.mkdirSync(folder);
    }
};

var uploadFolder = '../upload/';

createFolder(uploadFolder);


const storage = multer.diskStorage({
    destination(req, file, cb) {
        cb(null, uploadFolder)
    },
    filename: function (req, file, cb) {
        var fileFormat = (file.originalname).split(".");
        cb(null, Date.now() + "." + fileFormat[fileFormat.length - 1]);
    }

})
const upload = multer({
    storage: storage, fileFilter(req, file, callback) {
        // 解决中文名乱码的问题
        file.originalname = Buffer.from(file.originalname, "latin1").toString("utf8");
        callback(null, true);
    },
});

router.post('/:id', upload.single('file'), async (req, res, next) => {
    console.log(req.file);
    const kb = await KnowledgeBase.findById( req.params.id );
    try {
        const uploadfile = new File({
            name: req.file.filename,
            filename: req.file.originalname,
            size: req.file.size,
            path: req.file.path,
            knowledgebase: req.params.id
        });
        const savedfile = await uploadfile.save();
        const knowledgeBase = await KnowledgeBase.findById(req.params.id);
        knowledgeBase.file.push(savedfile._id);
        await knowledgeBase.save();
        const filepath = req.file.path;
        await EmbeddingFile(filepath);
    } catch (err) {
        fs.unlinkSync(req.file.path);
        res.status(500).json({ message: 'An savedfile error occurred ' });
    }


});





router.get('/:id', async (req, res) => {
    // try {
        const file = await File.findById(req.params.id);
        const filepath = file.path;

        let ok;
        EmbeddingFile(filepath);
        // const pythonProcess = spawn('python', ['embedding_txt.py']);
        // // const jsonStr = JSON.stringify(filepath);
        // // console.log(jsonStr);
        // // 向子进程发送参数
        // pythonProcess.stdin.write(filepath);
        // // console.log(jsonStr);
        // pythonProcess.stdin.end();
        // pythonProcess.stdout.on('data', (data) => {
        //     console.log(`执行: ${data}`);
        // });
        //
        // pythonProcess.stderr.on('data', (data) => {
        //     console.error(`执行出错: ${data}`);
        // });
        //
        // pythonProcess.on('close', (code) => {
        //     console.log(`子进程退出，退出码 ${code}`);
        // });
        if( ok===true){
            res.json("完成");
        }else{
            res.json("失败");
        }

    // } catch (err) {
    //     res.status(500).json({ message: 'An error occurred', error: err });
    // }

});

// router.get("/", (req, res) => {
//   res.send("Hello World!!!");
// });



export default router;




