
# Boost Your RAG Chatbot’s Accuracy with Synthetic Test Data Generation!

**Published:** 2024-08-15
**Author:** AI Anytime
**Description:** In this video, I delve into the importance of synthetic test data generation, especially for evaluating RAG (Retrieval-Augmented Generation) augmented pipelines. Manually creating QA samples from documents is time-consuming and often lacks the complexity needed for a thorough evaluation. By leveraging synthetic data, developer time in the data aggregation process can be reduced by 90%.

I also introduce Ragas, a novel approach to generating evaluation data. Unlike standard methods, Ragas uses an evolutionary generation paradigm inspired by works like Evol-Instruct. This approach ensures that your evaluation dataset includes a wide range of question types and difficulty levels, providing a more comprehensive assessment of your pipeline’s performance.

If you find this video helpful, don't forget to like, comment, and subscribe for more content like this!

GitHub: https://github.com/AIAnytime/Synthetic-Test-Data-Generation-for-RAG

Join this channel to get access to perks:
https://www.youtube.com/channel/UC-zVytOQB62OwMhKRi0TDvg/join

To further support the channel, you can contribute via the following methods:

Bitcoin Address: 32zhmo5T9jvu8gJDGW3LTuKBM1KPMHoCsW
UPI: sonu1000raw@ybl

#ai #rag #ragas

---



## Section starting at 0.00


- Hello everyone, welcome to AI anytime channel.

- In this video, we are going to look at synthetic test data generation.

- So most of you would have been working with RAG, the retrieval augmented generation.

- And you would have been dealing with, you know, like large text document or large data

- sets where you have, you know, multi context data.

- Now, how can you evaluate your retrieval systems or the RAG system that you have built?

- Now to manually baseline these RAG systems are a mundane task, you know, it requires

- a lot of efforts.

- So in this video, we are going to use RAGAS to create or generate synthetic test data based

- on the context that you already have based on the data that you already have.

- And that will help you in evaluating in your RAG system.


---


## Section starting at 62.52


- Now, RAGAS is evolved to evaluation metrics like you would have heard of different evaluation

- metrics, you know, for RAG, you know, for retrievals, for synthesizing the responses.

- When you talk about retrievals, you have distance based evaluation metrics like cosine,

- jacquard, levinstein, Euclidean, blah, blah, blah.

- And you have, because these all falls under approximate nearest neighbor A and algorithms.

- And you know, synthesizing when the LLM synthesizes the context, once it synthesizes the response

- from the retrievals, then you will use different evas, both qualitative and quantitative.

- So that's where we have used RAGAS that helps you evaluate RAG systems.

- Now they have, you can use RAGAS to generate good quality synthetic test data.

- So let's jump into it here, guys, and see how we can do that.

- Now, RAGAS is an open source library, so you can use it.

- Now, to do that, first of all, we have to install a few things.

- I'm going to do this in Google Collapse.

- You can see I have my collab open here, so I'm just going to do pip install.


---


## Section starting at 125.72


- And we are, we are going to use the RAGAS.

- And I'm going to use Langshane OpenAI.

- You can use GROC, you can use Anthropic, you can use any other LLM providers if you want.

- You can do it locally as well if you have the compute.

- And I'm going to do Stentence Transformers.

- I'm going to use an open source embedding model, because I'm embedding, I don't want

- to exhaust my credit for OpenAI.

- So that's why I'm going to use BGE embeddings, and I'm probably going to show that.

- We've installed RAGAS, Langshane OpenAI, Stentence Transformers, XML dict, because you're going

- to just convert that.

- Now, once you are okay with that, what we're going to do is we're going to import a few

- things.

- So let's just import it.

- So what I'm going to do is import OS, and from Google Collab, import user data.

- Now Google Collab, import user data, why I'm doing it?

- The reason being, because all of my secret skis are here, you can see Anthropic, Gemini,


---


## Section starting at 189.48


- Groke, OpenAI, Run, Ports, Per, TabiLi.

- So I have all of my key over here.

- So I need the access to that.

- So I'm going to use the data.

- Basically, I'm going to set that key in ENV, and that's why I'm using it.

- Now I need Pandas here.

- So import Pandas as PD, because I'm going to show that in a data frame.

- Now all the Langshane thingy.

- So from Langshane community, every open source thingy can be found in Community Word module

- of Langshane, document loaders.

- So I just want to use document loaders.

- And let's take a medical use case here.

- So for example, if you are doing medical biomedical literature or something, you are building

- a rag on top of it, then how you can generate test data sets.

- So let me just do PubMed loader over here.

- If you can do the same with Lama index as well.

- So I'm going to do from Langshane.

- We need embedding.

- So I just said if everything is open source, completely open source.

- Groke is not open source by the way, don't fall into all those dilemmas.


---


## Section starting at 252.36


- Like they they they made you available, right?

- They make things available through their data centers and servers.

- They make Lama 3.1 available.

- Same happens with Azure, you know, GCP and AWS, all these open source dilemmas are available

- in their model repository.

- The only thing which is completely open source, which is available directly on Hugging

- face.

- Okay, when it comes to a large language model from Langshane community, okay, dot embedding.

- So I'm going to just do embedding here and let's use Hugging face embeddings.

- Okay, you can see I'm using Hugging face embeddings.

- And you can also use Hugging face BGE embeddings guys, it also has directly, so let me just

- do BGE, you can see it over here, okay, it's by Beijing AI Academy BAAI is one of my favorite

- embeddings model out there.

- Now I'm going to just do from Langshane.

- The old Lama and all are very new thing guys, you know, I mean, when I say very new, it

- compared that with sentence, Transformers library and all like they came up with no


---


## Section starting at 315.54


- economic embeddings, they make it available and all Lama CPP is a bedrock of cobalt of,

- you know, all Lama and you know, see Transformers, I mean, I've done a lot of others, okay.

- Langshane community and not the camps community, now let's have open AI.

- So from Langshane open AI, I'm going to import chat open AI.

- You can also use open AI, you can also use Azure chat open AI, you can use Azure open AI

- with both, all instruction, instruction based and also the base model is available.

- Now the ragas thingy, so from ragas dot test, you have to use test set generator, so this

- is the module that we're going to use test set generator and I'm going to use import.

- That's right.

- And then from ragas dot test set and then you need evaluation.

- I don't know why it's not showing evaluations, it's going to be evaluator from test set


---


## Section starting at 379.26


- or evaluations, whatever, okay.

- And then let's just get import from test set dot, I hope I spelled it right, so it

- should be, let me, you can do a directory as I think it's evolutions, yes, excuse me,

- I was typing it right, evolutions and then you imports different evolutions, it provides

- you reasoning based multi-context, simple and what not, so what I'm going to do is simple

- and what next multi-context, you can see it over here, multi-context and then a

- reasoning, yes, reasoning, perfect guys, we have our everything, now let's get the open

- a I thinky here, so oh.invahn, so I'm going to set that in environment, generator.get and

- that's it.

- When I give the grant access here, so let's just give a grant access, now I have to grant


---


## Section starting at 441.34


- the access guys and that's why I'm granting here, now once you grant the access, now let's

- go into first, declaring the LLM for generating, so I'm going to call the data generation model

- or something, and you can use ntropic, whatever you want, okay, if you want to use, let's

- don't use turbo, I want to use 4Omini, the latest by LL, OpenAI, so I'm using GPT-4Omini,

- here you can see that and then I'm going to use critic model, so critic model and in critic

- model I'm going to pass chat, openAI, and not, I'm not going to use GPT-4, same model,

- like probably let's use the more capable for complex task because we are going to use

- this to act as a critic model, so I'm going to use critic model chat, openAI model, GPT-4O

- and then I'm going to use the embedding model, which is BGE, so for that, I'm going to

- copy paste code guys from lang chain, okay, so I'm going to use BGE embeddings, BGE on


---


## Section starting at 505.10


- hugging phase, just come here and copy this, sit, so just copy this, come here and just

- paste, model and just let's make this embeddings, now what we are doing, we are saying, okay,

- model name, we are using BGE, small embedding model, you can also use the large one, but

- I think just for the demo purpose, small make sense, I'm having CPU, I'm not using

- CUDA device, so device CPU, I'm asking to normalize the embeddings, you know, just normalization

- true and then just calling that function embeddings, and this will download, even if you are

- running for the first time, we'll take a bit of time, let's do that, now let me define

- my loader thing, so for the loader, I'm going to use PubMed, PubMed is a biomedical literature

- data repository, PubMed central, where you can use to do a lot of things, like reading

- research papers, creating data to train something and whatnot, so I'm going to do is, you

- can see the model has been downloaded from hugging phase, now PubMed COVID, let's use the


---


## Section starting at 568.02


- word, for example, cancer, if you want to find out, you want to build a rack system on

- top of cancer, on top of mental health, top of ADHD, top of liver, whatever, you just

- give that word, and then you can specify a params called a load max or something, you can

- see it over here, it's showing you load max docs in teacher, so I'm just going to do load

- max docs, and in the load max docs, let's keep it five, but you can increase it depending

- on, if you have what kind of credits that you have or what kind of model you are using,

- now when you do this loader, you can just do documents, pretty much straight forward

- length and code, loader, load, so let me just do loader, load, and here, when you print

- the documents thingy, when you print the documents thing, it's just too many requests, waiting

- for 0.2 seconds, let it wait, meanwhile, if it's take a lot of time, I will just walk you

- through my GitHub code, because now let me just try it out, document right now, it will

- probably nothing is there, okay, you can see here, right, we got it, so probably it took


---


## Section starting at 630.70


- a bit of time to get that, because it's free, right, so when you hit this again and again,

- it might take a bit of time for the API, they would have some rate limits, you know, from

- users, you can see all the metadata over here, so title, externalization of scan protocols

- for RTC, T-stimulator from different vendors, and then it talk about cancers, they're all

- cancer related biomedical literature, and I want to build a rag on top of it, now if

- I'm building a rag on top of it, how can I evaluate, how can I baseline, if you are

- working on an enterprise, the first thing would be asking your client, that give me the

- question answers to baseline, it's like an baseline, you know, on top of your data, and

- if you are like, have, you know, humongous data like in GB, there is very difficult to

- go and, you know, look at that context, look at the paragraphs and create questions for

- the testing, you can just give it to Raga and it will automatically create for you, and

- that's the agenda of this video, now let me just move to the next, which is like creating

- a generator in stance here, so let me just do that, so for the generator, I'm going


---


## Section starting at 695.78


- to use test set generator from Lang chain, and you can use the same for a Lama index, now

- you can see test set generator, and then I'm going to use test from Lang chain, and in

- from Lang chain, I'm going to pass a few things, first thing is yes, data generation model,

- and then I'm going to pass my critic model, it missed the critic model, did I define critic

- model or didn't I, okay, let me just run this again, I don't know why it missed it, but

- and I'm going to pass my critic model here, so you can see critic model, and let's just

- run it, now once you run this, we have a generator in stance, if you, if you like, it will

- be an object or something, if you do that, you can find it out here, max retry train time

- out like 180, max wait 60, max worker, it gives you an inference params over here, and

- you're going to look at in the generator that instant that we've created, now you can have

- a distributions, like if you distributions can be of three different types, simple, simple

- questions, then multi context, where you have multiple contexts for a given question


---


## Section starting at 759.26


- all together, and then you can have the reasoning as well, and you can divide the distribution,

- I think it's the most important part of creating test, or any kind of data guy that you would

- use, data should be diverse, data should be really diverse in nature, some of those should

- be complex, some of those should be simple, some of those should be multi context, and so

- on and so forth, and that's what we're going to use here, so let me just do distributors,

- and you can see it's suing me, but it's it's it's a it's a, okay, now you can see simple,

- but we have to give a number guys, you know, this is all wrong, okay, let me just, this

- is not how you divide, it's simple, now in simple, we have basically give a value, so

- we have to divide, divide the this in from g of out of one, so you have to define it,

- so if you have three classes, like simple multi context reasoning, you have to keep the

- value that the sums would be one, so how I'm going to do this is, let's keep simple


---


## Section starting at 821.90


- as like more 0.6, and then I'm going to keep multi context, you can see a reasoning, how

- much is 0.6, 0.2, 0.2 or something, I can, I'm going to keep simple more, and multi context

- 0.2 a reasoning one, because I don't want, because it will take time guys, it has to do

- a lot of iterations, so 0.7, 0.2, 0.1 when the summation becomes one, so let me just

- run the distributors, and now you can just create, create the test set instance, let me

- just create test set, have a variable, and then use generator dot, generate with lang

- chain, you can see it over here, it has lama index, and it has lang chain, lang chain

- docs, and you can just pass the documents, you can pass the documents, and the distributors,

- so how many you want, for example, how many data that you want, example, so let's keep

- five, you know, five iterations, no documents, five distributors, and then you can just


---


## Section starting at 882.86


- run it, now once you run it, it will take a bit of time, so what I'm going to do, probably

- let me open my GitHub, so I can just show you right there, just to save some time, because

- we are done with our code guys, so I'll just, I'll just show you here, now this is what

- happens, I'll just make this a bit big, so you can see it, I don't know what happened

- to GitHub, but anyway, let's do first run it, if it, like if you throw me an error or something

- because I just, you can see the embedding nodes, it's creating the embedding nodes for

- you, and then it starts generating, so if you look at the generating part of it, let

- me make it a bit bigger, so you can see it, and you can see as I said, simple equal to

- more like 0.7, it's a bit fast, now if you make multi context and reasoning, then it

- has to go through, it will create multiple context, and based out of it, it will create

- a lot of questions, you can see this, I don't know, let me just replace this here, something


---


## Section starting at 947.94


- is wrong with, you know, GitHub, yeah, it's like, I don't know, but anyway, let it happen,

- I'll show you, I think it's probably due to size, should I reduce it, yeah, 150, 150

- is good, now if you look at this, this is what I kept earlier, simple 0.5 and multi context

- 0.4 as a distributor, distributors, now if you look at this, probably, it's not here,

- so you, but this is how it creates, this is how it creates questions, context, ground

- truth, and evolution type, the metadata and everything, so you also have a ground truth,

- you can use the same in the ragas eval, because ragas as both qualitative, this ragas is

- more qualitative eval's guys, where it looks at the faithfulness, relevancy, you know,

- and those kind of things, if you want more quantitative, you have to look at the, uh,

- board score, mover scores, and you know, those kind of, uh, scores to calculate it basically,


---


## Section starting at 1008.22


- you can say it's 100%, so it's, it's done, now I want to convert this in data frame,

- so let me just do test DF, and in the test DF, I'm going to create test set dot 2 pandas,

- yes, that's right, 2 pandas, and now it will create it, so let me just do test DF, when

- you do test DF, you can see the question context, and that's going to make this a bit, convert

- this into an interactive table, and I'll just remove this, now you can read it, question.

- What role does BRD for recruitment play in the formation of PML, RAR, alpha-meditated

- micro-specals in acute, promiliotic leukemia, and it has the context, it has the ground truth.

- Now, if you really want to evaluate your rag in a better manner, you should keep my

- multi-context and reasoning more, and keep simple as like less, like 0.30, 0.4, and multi-context

- a bit higher, and reasoning a bit higher, so it will give you better logical questions

- and evalves, and evaluate on that, now you can find out all the questions, if you can


---


## Section starting at 1070.76


- look at this, you can look at this multi-context, it is, it is multi-context, it says, how do

- racial disparities in cancer, predictor, docs affect immunotherapy strategies, and it

- is multi-context, and it also has your metadata, it has your ground truth, and what not, right?

- This is fantastic guys, I mean I liked it, now you can also like generate some graphs

- and all, and if you want, it will generate some graphs, you can look at the, I'm using

- like colapro, that's why I have these features, if you are probably doing it, you will not

- have these features, and if you look at categorical distribution, it will add a new cell, and

- you can just run it, excuse me, once you run it, it will probably just show you the categorical

- distributions, frequency, and everything, a lot of good graphs guys, you can just have

- a look at that, now this is fantastic, because you save a lot of time, and you can not

- only that, right?

- You can create question answers as well guys, high quality questions from this technique,


---


## Section starting at 1130.86


- you know, because a lot of you ask, how to generate questions, you know, from a set

- of document or a set of data, this is how you can do it, and not only, you state forward

- questions, but multi context based questions, or reasoning, high quality reasoning, that

- involves high quality reasoning, you know, behind a question, so you can also do that, and

- I think you should try it out, and let me know your feedbacks on this as well, because

- I think I'll be doing it for a while, now this is what I wanted to do in this video guys,

- you know, I wanted to show you how you can generate high quality synthetic data for your

- RAC systems for evaluating your RAC, for baseline in your RAC, and what not, I also have

- other videos on synthetic data generation, using grittle, grittle is a fantastic, it's

- the start up by the way, but you know, it's a fantastic framework to create synthetic

- data, I have through RGLA district label, and a lot of others, like LLM through LLM as

- well, I'm going to create some videos on the motor on soon, the notebook will be available

- on my GitHub repository, I was showing you, find the link in description, and if you like


---


## Section starting at 1195.02


- the content, please hit the like icon, if you haven't subscribed to the channel yet, please

- do subscribe to the channel guys, that motivates me to create more such videos in your future.

- Thank you so much for watching, see you in the next one.


---

