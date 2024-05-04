import express from 'express';
const app = express();
app.use(express.urlencoded({ extended: true }));
app.use(express.json());
import dotenv from 'dotenv';
dotenv.config();
import('./routes/api/user.js').then(userRouter => {
    app.use('/admin_api/user', userRouter.default)
});

const routes = {
    '/admin_api/model': 'model',
    '/admin_api/knowledgebase': 'knowledgebase',
    '/admin_api/app': 'app',
    '/admin_api/file': 'file',
    '/admin_api/upload': 'upload',
    '/admin_api/remove': 'remove',
    '/admin_api/issue': 'issue',
    '/admin_api/stopapp': 'stopapp',
    '/admin_api/testapp': 'testapp',
    '/admin_api/stoptestapp': 'stoptestapp',
    '/admin_api/download': 'download',
    '/admin_api/query':'query',
};

Object.entries(routes).forEach(([path, route]) => {
    import(`./routes/${route}.js`).then(module => {
        app.use(path, module.default);
    });
});

for (let path in routes) {
    import(`./routes/${routes[path]}.js`).then(module => {
        app.use(path, module.default);
    });
}

// 使用ESM导入数据库模型
import { User, File, KnowledgeBase, App } from './db.js';

app.post("/", (req, res) => {
    res.send("Hello World!!!");
});




app.listen(3000, () => {
    console.log("server on port 3000 running!");
});
