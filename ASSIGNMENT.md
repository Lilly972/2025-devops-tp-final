# TP Final Devops M1 Cyber

Hello All!
For this assignment you will have to apply everything we have seen during the course.
It is not hard, you almost know everything you need to do it.
But you will have to think a bit and read carefully the instructions.

## Let's introduce the team

Hello! I'm Jho, I have the best business idea ever and I want to pitch it to get investors.
I don't know anything about development, in fact, I vibe-coded the entire prototype.
The thing is, investors do not care about the code, if I want to explain them the idea, I want to show them the website.
I sent them the website on `http://localhost:3000` but it does not seem to be working on their computer.
Odd, it works on my machine. Anyway, I asked my best friend ChatGPT and he told me about DevOps.
I think I need to setup a DevOps strategy for this project and that's where I need you.
Because I saw a penguin, a big whale and an octocat and now, I am totally lost.

In the next few months, we'll be at least 3 on the project, me on vibe-coding, Bob on the project management and Yann to clean the vibe-coded code.
It is important that we can see the new features of the app in a shared environment without impacting the users.

The project is a Christmas gift organizer. You can add your family and friends to the app and add gifts ideas so you can manage this easily.

## Your task

You need to define a complete DevOps strategy for this team, from the usage of Git to the deployed application.
You will have to document all of that in the project.

### Minimum requirements

- Automated CI/CD pipelines
- Documentation (with storybook), frontend, backend, database deployed (any provider, you can use a single one or multiple)
- Dockerfile for frontend and backend and deployed images on Docker Hub
- docker-compose file to run the full application (frontend, backend, database)
- Source code on GitHub
- Automatic creation of GitHub release for the project
- documentation on what you did in the `docs/` directory of the project
  - if you chose another platform than the ones we saw in class, document why and add clear steps on how you deploed on those platforms

### Your choices

You are free to choose the tools you want to use for this project:

- CI/CD runners
- Deployment platforms
- git strategy
- environment management
- And so on...

All changes you make should be documented in the project.

Be careful, the app's code should not be modified.

### Bonus points

You absolutly don't have to do that and if you do that instead of a reqauired part, you will still have less points.
Depending on how cleaned you made

- on premise deployment
- kubernetes deployment
- monitoring setup
- anything I think is more complicated than what I asked but brings you more knowledge or is more relevant for the project.

## Important notice

You might want to setup multiple environments to solve this project, you can add the required elements in your CI/CD to do that and always deploy on the same platform. For example, if you want a preprod and a prod environment on render, you can use the same one and erase the production deployment when you deploy on preprod. This is only to make it easier for the assignment, it should require minimal changes to add the preprod environment (minimal = less than 5 lines in a file)

If you wanted to implement something and you didn't have time or you didn't manage to do it, please document what you wanted to do and why you didn't do it.
