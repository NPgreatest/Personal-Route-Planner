import express from 'express';
import { App, Model,  User, File, KnowledgeBase } from '../db.js';

const app = express();
const router = express.Router();

// 由于Express已经内置了body-parser的功能，这里直接使用express.json()来解析JSON格式的请求体
app.use(express.json());




router.get("/hello", (req, res) => {
    res.send("Hello World!!!");
});


router.post('/', async (req, res) => {
    const app = await App.findOne({ name: req.body.name });
    if (app) {
        res.status(400).json({ message: 'App name already exists' });
    } else {
        const newApp = new App(req.body);
        try {
            req.body.bind=parseInt(req.body.bind,10);
            const savedApp = await newApp.save();
            // 如果请求体中包含knowledgebase数组，则更新每个knowledgebase的app字段
            if (Array.isArray(req.body.knowledgebase)) {
                for (let id of req.body.knowledgebase) {
                    await KnowledgeBase.findByIdAndUpdate(id, { $set: { app: savedApp._id } });
                }
            }

            res.json({ message: 'The app entry has been created successfully', data: savedApp });
        } catch (err) {
            res.status(500).json({ message: 'An error occurred while creating the app entry', error: err });
        }
    }
});

router.get('/', async (req, res) => {
    try {
        const apps = await App.find();
        console.log("get:app/");
        res.json(apps);
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the app entries', error: err });
    }
});

router.get('/:id', async (req, res) => {
    console.log("get:app/:id");
    try {

        const app = await App.findById(req.params.id);
        const knowledgebase = await KnowledgeBase.find({ _id: app.knowledgebase });

        if (app) {
            res.json({ app, knowledgebase });
            // res.json(app+knowledgebase);
            // res.json(app,knowledgebase);
        } else {
            res.status(404).json({ message: 'No app found ' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while retrieving the app entry', error: err });
    }
});

router.put('/:id', async (req, res) => {

    try {
        const oldApp = await App.findById(req.params.id);
        const oldKnowledgeBaseIds = oldApp.knowledgebase;
        const newKnowledgeBaseIds = req.body.knowledgebase;

        // 找出需要删除的knowledgebase
        const knowledgeBaseIdsToDelete = oldKnowledgeBaseIds.filter(id => !newKnowledgeBaseIds.includes(id));
        // 找出需要添加的knowledgebase
        const knowledgeBaseIdsToAdd = newKnowledgeBaseIds.filter(id => !oldKnowledgeBaseIds.includes(id));

        // 删除knowledgebase的app字段
        await KnowledgeBase.updateMany({ _id: { $in: knowledgeBaseIdsToDelete } }, { $unset: { app: 1 } });
        // 添加knowledgebase的app字段
        await KnowledgeBase.updateMany({ _id: { $in: knowledgeBaseIdsToAdd } }, { $set: { app: req.params.id } });

        const updatedApp = await App.findByIdAndUpdate(req.params.id, req.body, { new: true });
        res.status(200).send({ message: 'The app entry has been updated successfully', data: updatedApp });
    } catch (err) {
        console.log(err);
        res.status(500).send({ message: 'An error occurred while updating the app entry', error: err });
    }
    // try {
    //     console.log(req.body);
    //     const updatedApp = await App.findByIdAndUpdate(req.params.id, req.body, { new: true });
    //     const updatedKnowledgebase = await KnowledgeBase.updateMany({ _id: { $in: updatedApp.knowledgebase } }, { $set: { app: updatedApp._id } });
    //     res.status(200).send({ updatedApp, updatedKnowledgebase });
    // } catch (err) {
    //     console.log(err);
    //     res.status(500).send(err);
    // }
});


//{"knowledgebase": ["655acbca7f3dc94ba84a9bf0","655acbfa7f3dc94ba84a9bf6"]},直接替换，不保留原本的
router.put('/prompt/:id', async (req, res) => {
    console.log('/prompt/:id');
    console.log(req.body);
    try {
        const updatedApp = await App.findByIdAndUpdate(req.params.id, req.body, { new: true });
        // const updatedKnowledgebase = await KnowledgeBase.updateMany({ _id: { $in: updatedApp.knowledgebase } }, { $set: { app: updatedApp._id } });
        res.status(200).send( updatedApp);
        console.log(updatedApp);
    } catch (err) {
        console.log(err);
        res.status(500).send(err);
    }
});
    // try {
    //     const putApp = await App.findByIdAndUpdate(req.params.id, req.body, { new: true });
    //     if (putApp) {
    //         res.json({ message: 'The App entry has been updated successfully', data: putApp });
    //     } else {
    //         res.status(404).json({ message: 'No file found with the provided name' });
    //     }
    // } catch (err) {
    //     res.status(500).json({ message: 'An error occurred while updating the file entry', error: err });
    // }


router.put('/file/:id', async (req, res) => {
    try {
        const updatedApp = await App.findById(req.params.id);
        updatedApp.file = req.body.file;
        updatedApp.knowledgebase = req.body.knowledgebase;
        console.log("updatedApp");
        await updatedApp.save();
        if (updatedApp) {
            res.json({ message: 'The app entry has been updated successfully', data: updatedApp });
        } else {
            res.status(404).json({ message: 'No app found with the provided name' });
        }
    } catch (err) {
        res.status(500).json({ message: 'An error occurred while updating the app entry', error: err });
    }
});

router.delete('/:id', async (req, res) => {
    try {
        const deletedApp = await App.findByIdAndDelete(req.params.id);
        const deletedKnowledgebase = await KnowledgeBase.updateMany({ $pull: { app: req.params.id } });
        res.status(200).send({ deletedApp, deletedKnowledgebase });
    } catch (err) {
        console.log(err);
        res.status(500).send(err);
    }



    // try {
    //     const deletedApp = await App.findByIdAndDelete(req.params.id);
    //     console.log("deletedApp");
    //     res.json({ message: 'The app entry has been deleted successfully' });
    //     // 获取被删除的app数据中的knowledgebase属性，假设为一个数组
    //     var knowledgebase = deletedApp.knowledgebase;

    //     // 遍历knowledgebase数组中的每个元素，假设为一个ObjectId
    //     for (var kbId of knowledgebase) {
    //         // 使用KnowledgeBase模型的updateOne方法，更新id为kbId的knowledgebase数据，将app属性中的id删除
    //         KnowledgeBase.updateOne({ _id: kbId }, { $pull: { app: id } }, function (err, result) {
    //         });
    //     }

    // } catch (err) {
    //     res.status(500).json({ message: 'An error occurred while deleting the app entry', error: err });
    // }
});

export default router;
