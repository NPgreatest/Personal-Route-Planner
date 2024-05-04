import { TextLoader } from "langchain/document_loaders/fs/text";
import {RecursiveCharacterTextSplitter} from "langchain/text_splitter";
import { OpenAI, OpenAIEmbeddings } from "@langchain/openai";
import { createStructuredOutputRunnable } from "langchain/chains/openai_functions";
import { FaissStore } from "@langchain/community/vectorstores/faiss";

import cheerio from 'cheerio';
import axios from "axios";

const openAI = new OpenAI({
    openAIApiKey:'sk-AYP9eKTcCHOqFrneC4E358Af7f12489dA601A5F012099fD1',
    modelName: "gpt-3.5-turbo",
    temperature: 0.01,
    topP:0.9,
    configuration:{baseURL:"https://api.kwwai.top/v1/"}
});
const embeddings = new OpenAIEmbeddings({
    openAIApiKey:'sk-AYP9eKTcCHOqFrneC4E358Af7f12489dA601A5F012099fD1',
    batchSize:512,
    modelName: "text-embedding-3-small",
    configuration:{baseURL:"https://api.kwwai.top/v1/"}
});



export async function WebSearchFunction(query) {
    try {
        const response = await axios.get('https://duckduckgo.com/html/', {
            params: { q: query }
        });
        const data = response.data;
        const $ = cheerio.load(data);
        let webContent = '';

        $('div.result').each(function () {
            const title = $(this).find('h2.result__title').text().trim();
            const snippet = $(this).find('a.result__snippet').text().trim();
            webContent += title + "\n" + snippet + "\n\n";
        });

        return webContent;
    } catch (error) {
        console.error('Error fetching web search results:', error);
        return '搜索时遇到错误，请稍后再试。';
    }
}


export async function GetKnowledgeBaseContent(query,filepaths){
    let newDb = await FaissStore.load(filepaths[0] +'_faiss_index', embeddings);
    for (let i = 1; i < filepaths.length; i++) {
        const otherDb = await FaissStore.load(filepaths[i]+'_faiss_index', embeddings);
        await newDb.mergeFrom(otherDb);
    }

    const docs = await newDb.similaritySearch(query, 3);

    var mergedContent = docs.reduce(function(merged, obj) {
        return merged + obj.pageContent;
    }, '');

    return mergedContent
}

export async function GetKnowledgeBasedAnswer(query,  filepaths, knowledgebaseContent,VECTOR_SEARCH_TOP_K=3, webContent, chatHistory = [], historyLen = 3) {
    let historyText = "历史问题有：";
    let count = 0;
    for (let i = chatHistory.length - 1; i >= 0; i--) {
        historyText += chatHistory[i][0] + "。";
        count++;
        if (count === historyLen) break;
    }
    historyText += "\n";

    let promptTemplate = `基于知识库片段内容回答当前问题，请简洁并专业地回答用户的问题,回答问题时字数尽可能少，尽可能控制在200字。\n如果知识库片段内容不匹配，无法回答该问题，则直接回答问题，并且说明这是通过网络搜索回答的。回答前先说明信息来源是网络搜索还是知识库。另外，答案请使用中文。历史问题:\n${historyText}知识库片段:\n${knowledgebaseContent}\n当前问题:\n${query}`;

    const completion = await openAI.call(  promptTemplate );
    return completion;
}
