# Gera-AI backend

---
## General

This is a backend for the gera-ai site, the deploy of which failed.<br>
Essence: The user enters a task (some question, to which there is a correct answer), then lists the things he is interested in, and the site, with the help of OpenAI generates a task, the condition of which is changed to suit the interests, without affecting the actual essence of it 
The site supports both an interest template (fixed interests, conditioned on one student) and an assignment template (fixed assignments, conditioned on one assignment variant for a whole group of ideas, each with different interests)
This site could be a great tool for tutors and teachers, because the main goal was to increase student interest in problem solving

Example generation:

Condition:

```
2 sulfur oxide 4 + oxygen = 2 sulfur oxide 6
how much sulfur oxide 4 should be taken if 76 grams of sulfur oxide 6 is taken
and the yield is 80%.
```

Interests:

```
Chemistry
Biology
```

Result:
```
An excited biology chemist studying the effects of various oxides on the growth of mutant tomatoes takes up his new experiment. Muttering under his breath that “a nation of tomatoes will soon lead the world,” he decides to use sulfur oxide 6 in his experiments.

Well, he begins his transformations, turning 76 grams of sulfur oxide 6. But he is a biological chemist, not a mathematical chemist, so it is difficult for him to calculate how much sulfur oxide 4 he would need to get the desired amount of sulfur oxide 6, given that the chemical reaction occurs in 80% yield. He knew that 2 molecules of sulfur oxide 4 would turn into 2 molecules of sulfur oxide 6 in a reaction with oxygen, but how much weight this would result in, he had no way of knowing.

“Eh, there, in the theory of genetic engineering, everything is simple...” - said the scientist thoughtfully, staring at his calculations. A task, oh, glory to science, he leaves to you.+
```
You can try everything yourself if you specify your token and proxy (when running from countries where openai does not work).

---
## Technology
Stack/Libs used:
- Bcrypt
- Fiber
- Gorm
- PostgreSQL
- Swagger
- Validation

Configuration via environment variables (```.env``` file for docker):
```env
POSTGRES_HOST=db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=geraai
POSTGRES_PORT=5432
JWT_SECRET=your_jwt_secret_key
OPENAI_API_KEY=your_api_key
PROXY_URL=your_proxy_url
```

# Run
To build and run use
```shell
docker compose up --build
```

To run use
```shell
docker compose up -d
```