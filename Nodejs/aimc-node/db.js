// 导入mongoose库
import mongoose from 'mongoose';
const Schema = mongoose.Schema;

mongoose.connect('mongodb://127.0.0.1:27017/local')
// mongoose.connect('mongodb://mymongo/local')
    .then(() => {
        console.log("数据库连接成功~")
    })
    .catch((err) => {
        console.log("数据库连接失败~", err)
    })
// 定义user表
const UserSchema = new mongoose.Schema({
    name: {
        type: String,
    }

});
const ModelSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,

    }
});

const LoseSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,

    }
});

const KnowledgeBaseSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,

    },
    details: {
        type: String,
        required: true
    },
    state: {
        type: String,
    },
    time: {
        type: Date,
        default: Date.now
    },
    file: [{
        type: Schema.Types.ObjectId,
        ref: 'file'
    }],
    app: [{
            type: Schema.Types.ObjectId,
            ref: 'app'
        }],
});


const AppSchema = new mongoose.Schema({
    name: {
        type: String,
        required: true,
    },
    details: {
        type: String,
    },
    model: {
        type: Schema.Types.ObjectId,
        ref: 'model'
    },
    time: {
        type: Date,
        default: Date.now
    },
    num: {
        type: Number,
    },
    state: {
        type: String,
    },
    link: {
        type: String,
    },
    pid: {
        type: String,
    },
    port: {
        type: String,
    },
    prompt:{
        type: String,
    },
    bind:{
        type: Number,
    },
    knowledgebase: [
        {
        type: Schema.Types.ObjectId,
        ref: 'knowledgebase'
    }],
    file: [{
        type: Schema.Types.ObjectId,
        ref: 'file'
    }],
});

// 定义file表
const FileSchema = new mongoose.Schema({
    name: {
        type: String,
    },
    filename: {
        type: String,
        required: true
    },
    size: {
        type: Number,
        required: true
    },
    path: {
        type: String,
        required: true
    },
    vecpath: {
        type: String,
    },
    state: {
        type: String,
    },
    addTitlePrefix: {
        type: String,
    },
    maxSegmentLength: {
        type: String,
    },
    meth: {
        type: String,
    },
    time: {
        type: Date,
        default: Date.now
    },
    knowledgebase: {
        type: Schema.Types.ObjectId,
        ref: 'knowledgebase',
        required: true
    }
});


export const User = mongoose.model('User', UserSchema);
export const Model = mongoose.model('Model', ModelSchema);
export const KnowledgeBase = mongoose.model('KnowledgeBase', KnowledgeBaseSchema);
export const File = mongoose.model('File', FileSchema);
export const App = mongoose.model('App', AppSchema);
