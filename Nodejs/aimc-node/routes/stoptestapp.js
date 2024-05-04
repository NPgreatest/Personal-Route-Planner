import express from 'express';
import net from 'net';
import fs from 'fs';
import multer from 'multer';
import { App, Model,  User, File, KnowledgeBase } from '../db.js'; // Make sure the path is correct

const app = express();
const router = express.Router();

app.use(express.json());




router.post('/:id',async (req, res) => {
    console.log("post:stoptestapp/:id")
    const stopApp = await App.findById(req.params.id);
    const pid=stopApp.pid;
    if (!pid) {
        stopApp.pid="";
        stopApp.link="";
        stopApp.state="未发布";
        stopApp.port="";
        stopApp.save();
        res.send({msg:"success",data:stopApp});
    }else{
        try {
            process.kill(pid, 'SIGTERM');

        } catch (error) {
            console.log(error);
        }
        stopApp.pid="";
        stopApp.link="";
        stopApp.state="未发布";
        stopApp.port="";
        stopApp.save();
        res.send({msg:"success",data:stopApp});

    }
});

export default router;
