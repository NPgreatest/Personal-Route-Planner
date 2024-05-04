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


router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});
router.post("/hello", (req, res) => {
    console.log(req.body);
    res.send('ok');
});




router.post('/:id', async (req, res) => {
    console.log("post:testapp/:id");

    const testapp = await App.findById(req.params.id);
    const issuefile = await File.find({ knowledgebase: testapp.knowledgebase });

    var pathArray = issuefile.map(file => file.path);

    let params = [
        JSON.stringify(pathArray),
    ];
    console.log(params[0]);
    let process = spawn('python', ['./python_script/app_prompt.py', ...params]);

    process.stdout.on('data', (data) => {
        let pythonOutput = data.toString();
        console.log(pythonOutput);
        let [python1, python2, python3,python4] = pythonOutput.split('\n');
        const pid = python1;
        testapp.pid = pid;
        newlink = python3.replace(/"/g, '')
        // testapp.prompt = req.body.prompt;
        testapp.link = newlink;
        testapp.state="调试中";
        testapp.port= python4;
        testapp.save();
        res.send(testapp);
    });

    process.stderr.on('data', (data) => {
        console.error(`执行出错: ${data}`);
    });

    process.on('close', (code) => {
        console.log(`Python script finished with code ${code}`);
    });
});

















// router.post('/:id', async (req, res) => {
//     console.log("post:issue/:id");

//     const issueApp = await App.findById(req.params.id);
//     const issuefile = await File.find({ knowledgebase: issueApp.knowledgebase });
//     var pathArray = {};
//     var newprompt = {};
//     prompt = req.body.prompt;
//     for (var i = 0; i < issuefile.length; i++) {
//         pathArray[i] = issuefile[i].path;
//     }


//     const { spawn } = require('child_process');

//     let params = [
//         JSON.stringify(pathArray),
//         JSON.stringify(newprompt),
//     ];


//     let process = spawn('python', ['./python_script/app_final.py', ...params]);


//     process.stdout.on('data', (data) => {

//         // console.log(`执行: ${data}`);
//         let pythonOutput = data.toString();
//         let [python1, python2, python3] = pythonOutput.split('\n');
//         const pid = python1;
//         issueApp.pid = pid;
//         newlink = python3.replace(/"/g, '')
//         issueApp.prompt = newprompt;
//         issueApp.link = newlink;
//         issueApp.state = "调试中";
//         issueApp.save();
//         res.send(issueApp);
//     });
//     process.stderr.on('data', (data) => {
//         console.error(`执行出错: ${data}`);
//     });

//     // 当Python脚本结束时，这个函数会被调用
//     process.on('close', (code) => {
//         console.log(`Python script finished with code ${code}`);
//     });



// });


export default router;
