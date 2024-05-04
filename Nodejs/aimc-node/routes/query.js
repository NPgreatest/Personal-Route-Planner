import express from 'express';
import net from 'net';
import fs from 'fs';
import multer from 'multer';
import { spawn, exec } from 'child_process';
import { type as osType } from 'os';
import { App, Model, User, File, KnowledgeBase } from '../db.js';
import {GetKnowledgeBaseContent, GetKnowledgeBasedAnswer, WebSearchFunction} from '../js_script/app_prompt.mjs';
import { QueryOpenAI }  from '../js_script/normal_openai.mjs';
const app = express();
const router = express.Router();

app.use(express.json());

// 沟通python导入的文件
function generateResponse(code, data = null, message = '', error = null) {
    return {
        code,
        data,
        message,
        error,
    };
}

let chatHistory = [];
router.post('/', async (req,res) => {
    console.log("post:query");
    const query = req.body.query;
    const historyEnabled = req.body.history;

    // 根据history字段更新聊天历史
    if (historyEnabled === "on") {
        chatHistory.push(query); // 添加当前查询到历史记录
    } else {
        chatHistory = []; // 清除所有历史记录
    }

    const issueApp = await App.findOne({bind: req.body.bind});
    if (issueApp ===null){
        res.json(generateResponse(404,"找不到APP","失败"));
        return;
    }
    const issuefile = await File.find({ knowledgebase: issueApp.knowledgebase });
    var pathArray = issuefile.map(file => file.path);
    let params = [
        JSON.stringify(pathArray),
        JSON.stringify(issueApp),
    ];
    if (issueApp.state!=='已发布') {
        res.json(generateResponse(404,"APP未发布","失败"));
        return;
    }

    let knowledgebaseContent = await GetKnowledgeBaseContent(query,pathArray);
    let webContent = null;


    let res_str = await GetKnowledgeBasedAnswer(query,pathArray,knowledgebaseContent,3,webContent,chatHistory);
    res.json(generateResponse(200, res_str, '成功'));

});



router.post('/normal', async (req,res) => {
    console.log("query openAI");
    const query = req.body.query;

    let res_str;
    res_str = await QueryOpenAI(query);
    res.json(generateResponse(200, res_str, '成功'));

});




export default router;
