Topic: API for RealWorld application

The problem of coming up with a good educational application is actually quite relevant :)

Almost always it is CRUD in one form or another. But just CRUD over one entity is not so interesting. Therefore, it is necessary to have more entities with all sorts of connections between them. This will allow you to gain experience in the application architecture and the approach with the repository pattern.

There is a repository on the Internet https://github.com/gothinkster/realworld . It is written in a clone of the Medium project in various languages, both frontend and backend. It is called Conduit. You can see it in reality on the website https://demo.realworld.io/#/ .

As you can guess, we will write our own backend for this application.

Or rather, some subset of it - I made a little less entities in the test than there are in realworld itself. +2 entities will not teach anything radically new, but will waste time.

Realworld itself has a set of tests via Postman ( https://github.com/gothinkster/realworld/tree/master/api ) - but I was not satisfied with those tests, so I implemented some of them on goshka so that you would not have to install any node. But it would be cool if your implementation passes some of these tests.

As part of this task, you need to try:
1. Splitting the project into separate components (repository handlers)
2. Sessions (transmitted via the Authorization header) must be stateful. If you decide to use JWT, the session must still be stateful. And it must be stored normally, and not so that the entire token is remembered.

The following entities will have to be implemented:
* User
* Session
* Articles with various filters

Of course, you can implement all the other entities with separate tests if you wish.

You can follow the proposed structure and not hit the packages, or you can slightly correct and divide as if it were in reality - see crudapp as a basis. The code should be written in the realworld.go file if you decide to follow the 1st option.

As usual in all assignments of this and previous courses - this description is a basic statement of the problem, everything else will have to be obtained from tests. There are not many tests, but they are done quite universally in a tabular form and with some magic, which will also be useful for you to understand :)

The realworld project also has a swagger scheme. It is attached to the assignment in the swagger folder. You can run the documentation server (located in the rwa/swagger folder):

* go get -u github.com/go-swagger/go-swagger/cmd/swagger
* swagger serve swagger.json -p 8085

Good luck!

P.S. There may be bugs in this homework, be careful. But I have a working solution, so first make sure that this is a bug, and not an implementation issue.

How tests work:

* In this task, you have a large set of integration tests. Integration means that they test the entire chain taking into account the changing state of the system - if I added something, then I should be able to read it later.

* Due to this, you can move literally one step at a time, finishing the code to execute the next test case - as it was in the 2nd homework with the game.

* All tests are located in app_test.go - there you get an HTTP handler (see the work plan below) and push it into the test server.

* the test has a name - it usually tells you what we are testing. The rest of the +- fields speak for themselves. Note - in some places there are triggers "before" and "after" - in them, for example, the installation and replacement of an authorization token occurs.

About the token:

A token is basically a session. But it is transmitted not via cookies, but via HTTP headers. You can either use the token as a session key, or make a JWT token with additional ones (you can read about it in redditclone). I recommend starting with a session key (which will simply be a key in your session map).

About POST&ko requests:

The request body is sent to the server in JSON format, not form-URL-encoded. So you need to read the body and unpack the JSON. Be sure to check for errors! The previous lecture covered how to read the request body.

Work plan:

* In terms of code, the entire lecture is dedicated to how to do this task. But do not copy-paste the code from there!!! write everything yourself
* As a result, you need to split the solution into packages in accordance with https://github.com/golang-standards/project-layout and what I told in the lecture with crudapp
* There are no parameters in the URL, but there is a division by GET / POST / etc methods - you can either go to the desired function through switch-case, or screw in another router (for example, gorilla / mux or something faster) and immediately hook the route to the desired method. There was an example in the lecture. but pay attention to how to use middleware in gorilla / mux - read there in the doc
* Remember that we need to split the logic into layers: handler -> repository -> db (currently slices and maps).
* If you want a clean architecture, then you will have more layers: delivery / http -> usercase -> repository -> db (currently slices and maps)
