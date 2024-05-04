import express from 'express';
import fs from 'fs';
import multer from 'multer';
import { App, Model, User, File, KnowledgeBase } from '../db.js'; // Ensure this path is correct
import path from 'path';

const app = express();
const router = express.Router();
app.use(express.json());

const uploadFolder = '../upload/';


router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello", (req, res) => {
    console.log(req.body);
    res.send('ok');

});

router.get('/:id', async (req, res) => {
  console.log("get:download/:id");
// 定义一个路由，用于处理get请求，假设请求的url为/file/download?id=123
    // 获取请求中的id参数，假设为123
    var id = req.params.id;

    // 使用File模型的findById方法，查找id为123的文件数据
    try {
        // 使用await等待File.findById(id)返回结果
        let file = await File.findById(id);
        // 如果找到文件，执行以下操作
        // 使用path模块的extname方法，从文件的路径中获取文件的后缀，假设为".txt"
        var ext = path.extname(file.path);
        // 使用res.download方法，下载文件，并设置文件的名字和类型
        res.download(file.path, file.filename, {headers: {'Content-Type': 'text/plain' + ext}});
      } catch (err) {
        // 如果发生错误，返回错误信息
        res.status(500).send(err);
      }
  });
export default router;
