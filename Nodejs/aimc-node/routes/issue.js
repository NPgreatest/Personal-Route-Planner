import express from 'express';
import net from 'net';
import fs from 'fs';
import multer from 'multer';
import { spawn, exec } from 'child_process';
import { type as osType } from 'os';
import { App, Model, User, File, KnowledgeBase } from '../db.js';
const app = express();
const router = express.Router();
app.use(express.json());


router.post('/:id', async (req, res) => {

    console.log("post:issue/:id");

    const issueApp = await App.findById(req.params.id);
    const issuefile = await File.find({ knowledgebase: issueApp.knowledgebase });

    var pathArray = issuefile.map(file => file.path);

    issueApp.state = "已发布";
    issueApp.save();
    res.send(issueApp);

    // let params = [
    //     JSON.stringify(pathArray),
    //     JSON.stringify(issueApp),
    // ];
    // console.log(params);
    // console.log(pathArray);
    // console.log(issueApp);
    // let process = spawn('python', ['./python_script/app_final.py', ...params]);
    //
    // process.stdout.on('data', (data) => {
    //     let pythonOutput = data.toString();
    //     console.log(pythonOutput);
    //     let [python1, python2, python3] = pythonOutput.split('\n');
    //     const pid = python1;
    //     issueApp.pid = pid;
    //     newlink = python3.replace(/"/g, '')
    //     issueApp.link = newlink;
    //     issueApp.state = "已发布";
    //     issueApp.save();
    //     res.send(issueApp);
    // });
    // process.stderr.on('data', (data) => {
    //     console.error(`执行出错: ${data}`);
    // });
    //
    // // 当Python脚本结束时，这个函数会被调用
    // process.on('close', (code) => {
    //     console.log(`Python script finished with code ${code}`);
    // });



});



export default router;
