# Glossarize backend

## About Glossarize

Currently, there are lots of various applications in the market which help you improve your vocabulary in different areas. The most known of them are _Lingualeo_, _Duolingo_, Skyeng's _Words_. They all allow you to add words to your list and learn them by trainings like cards which's quite useful and comfy. But you can only add words to your list from their words base which definitely can't have all possible words definitions and in many cases will make you sad in scenarios when you've heard or seen a nice collocation or idiom and want to keep it in the list. 

As an alternative, you could've used your own copybook or even _Google Sheets_ to store words in your own custom way which gives you a comprehensive opportunity to manage the vocabulary. But it's for sure distressful to cope with the learning part: you should hide the translations by your hand to not let yourself cheat, you can't mix the words etc. It's obviously a drawback of this approach.

The glossarize idea is to combine the learning applications convenience and the manual vocabularies flexibility by:
- **first:** allowing a user to create its own distinct vocabularies (the French one, the Zoomers one, the Internet Technologies one with corresponding terms) each with its own words, descriptions, examples and photos to increase the memorize efficiently by incorporation as many memorize approaches as possible. 
- and **second:** providing a nice way to learn and repeat the added words with different exercises and manage the exercice type and its content (the list of words to be learnt in the exercise) by yourself out of a words folder.

## Personal aims concerning the project

By the date of describing the project (14-09-2020), I'd been working as a software developer for almost two and a half years. Almost 100% of the experience I had was about managing and developing services that had been schemed by someone else and all their components had already been coupled and launched by someone else. In these cases, I integrated into the workflow and accepted the rules of the game which had been written by someone else. Of course, I further participated when it came to the architecture deviations (refactoring) and improvements, but it was always about complete-ish applications, architectures and compositions. As far as I'd been told, a good software engineer should be capable of making such things by himself (or herself). The following is the list of what I was about to achieve by building this application:
1. get a reliable perfect-for-me application to improve vocabulary;
2. try me in things I'd never tried:
2.1. microservices containerization and coupling (architecture experience);
2.2. services throughput improvement (scaling experience);
2.3. thorough usage of gizmos invented for all these (k8s, rabbitMQ, redis or similar applications);
2.4. work with hired designers and other developers including frontenders and mobilers (manager experience);