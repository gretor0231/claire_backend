# claire_backend
Claire_English_Backend_Assignment

Backend Email

POST https://yourwebsite.com/emails
with JSON body:
{
    mail: {
        to: "you@there.com",
        subject: "Test email",
        message: "Hello world!",
    }
}
"yourwebsite.com" will be replaced with the domain name that you set for your public website. It doesn’t need to be a top-level domain.



How to start:
go to project folder

1. sudo make up_build
2. sudo make start

test page: localhost:80
Mail Hog: localhost:8025