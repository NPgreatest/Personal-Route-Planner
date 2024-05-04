const express = require('express');
const app = express();
app.use(express.urlencoded({ extended: true }));
app.use(express.json());
const userRouter = require('./routes/api/user')
app.use('/api/user', userRouter)
const routes = {
    '/api/model': require('./routes/model'),
    '/api/knowledgebase': require('./routes/knowledgebase'),
    '/api/app': require('./routes/app'),
    '/api/file': require('./routes/file'),
    '/api/upload': require('./routes/upload'),
    '/api/remove': require('./routes/remove'),
    '/api/issue': require('./routes/issue'),
    '/api/stopapp': require('./routes/stopapp'),
    '/api/testapp': require('./routes/testapp'),
    '/api/stoptestapp': require('./routes/stoptestapp'),
};

for (let path in routes) {
    app.use(path, routes[path]);
}



// 导入数据库
const { Lose, User, File, KnowledgeBase, App } = require('./db'); // 引入我们之前定义的模型

app.post("/", (req, res) => {
    res.send("Hello World!!!");
})
app.post('/publish', async (req, res) => {
    // 1. 获取前端传过来的数据
    // 2. 存入Lose 数据表里
    try {
        const { name } = req.body;
        await Lose.create({
            name
        });
        res.send("success")
    } catch (error) {
        res.send("error")
    }

})




app.listen(4000, () => {
    console.log("server4000 running!")
})