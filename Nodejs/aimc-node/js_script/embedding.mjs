import { TextLoader } from "langchain/document_loaders/fs/text";
import {RecursiveCharacterTextSplitter} from "langchain/text_splitter";
import { ChatOpenAI, OpenAIEmbeddings } from "@langchain/openai";
import { FaissStore } from "@langchain/community/vectorstores/faiss";

export async function EmbeddingFile(filepath) {
    try {
        const embeddings = new OpenAIEmbeddings({
            openAIApiKey:'sk-AYP9eKTcCHOqFrneC4E358Af7f12489dA601A5F012099fD1',
            batchSize:512,
            modelName: "text-embedding-3-small",
            configuration:{baseURL:"https://api.kwwai.top/v1/"}
            });
        let docs = [];
        const fileType = filepath.split(".").pop();

        let loader;
        let documents;
        let splitDocs;
        let db;
        const textSplitter = new RecursiveCharacterTextSplitter({
            chunkSize: 500,
            chunkOverlap: 200,
        });

        switch (fileType) {
            case 'txt':
                loader = new TextLoader(filepath);
                documents = await loader.load();
                docs.push(...documents);
                splitDocs = await textSplitter.splitDocuments(docs);
                db = await FaissStore.fromDocuments(splitDocs, embeddings);
                break;

            default:
                throw new Error(`Unsupported file type: ${fileType}`);
        }

        const path = `${filepath}_faiss_index`;
        await db.save(path);
        return true;
    } catch (e) {
        console.error(`Error: ${e.message}`);
        return false;
    }
}

