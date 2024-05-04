import { OpenAI } from "@langchain/openai";


export async function QueryOpenAI(query) {
    const openAI = new OpenAI({
        openAIApiKey:process.env.API_Key ,
        modelName: "gpt-3.5-turbo",
        configuration:{baseURL:process.env.Base_URL}
    });

    const completion = await openAI.call(  query );
    return completion;
}
