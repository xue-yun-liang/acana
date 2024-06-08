Nowadays, more and more people are learning on social Q&A platforms, such as Zhihu in the Chinese community and StackOverflow in the English community. However, with the passage of time, there have been more and more posts on such websites, and it has become difficult to quickly obtain the desired information through simple keyword searches in the past. Therefore, this project combines the RAG system to build a knowledge based social platform, aiming to improve the platform's search and knowledge generation capabilities in conjunction with LLM, and ultimately achieve the goal of improving user learning efficiency.

Usually, this is also a project used for learning with Go.
Currently in the development stage, the technologies planned to be used are as follows: web part: `Golang`, `MySQL` (user&connect store), `Redis`, `Zap` (logger record), `Viper` (config manager), `gin` (web frame); RAG section: `langchaingo`, `weaviate` (vector database).