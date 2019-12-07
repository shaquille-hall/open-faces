# Open Faces Blockchain Simulator

This repository contains my solution for a blockchain simulator for my SE575 Software Design class.


## Background

The more we use technology, the more it becomes interwoven into our daily lives.
With security breaches seeming a daily occurence, and proper authentication
becoming even more enforced, a great playground arose to explore blockchain's
fundamental security characteristics. As the adage goes, nothing is more secure
than what is hidden in plain sight.

This simulator can be found live at https://open-faces-8877b.firebaseapp.com/ 
An associated youtube demonstration can be found at: https://youtu.be/ygvUv7MiUJo


## Tool choice and setup

This solution was built using the Angular 8 web framework, powered by a Go API.
The frontend is hosted on Firebase, while the API makes use of Google Cloud
Platforms' Cloud Functions.


![High Level Diagram](https://drive.google.com/uc?export=view&id=1kSEG9uzA7p0w_nsUyNffCKNxbJYWsBJf)


**Angular:** *[link](https://angular.io/start)* Angular has long been known for
its prominence among web frameworks. Particularly, quick prototyping was an
absolute for this project, which Angular seamlessly allows.

**Go:** *[link](https://golang.org/)* Blockchain involves quite a bit of
hashing, which can often come increased execution time. Being that Go was
built with performance in mind, I thought it would be an appropriate choice
to handle this use case.

**Firebase:** *[link](https://firebase.google.com/)* My original goal was to
have a single host for both my client and API. However, as it turns out
Angular + Firebase is a bit simpler to setup than Angular + GCP App Engine.
I admit this could have been due to my lacking familiarity with web hosting.
Further realization alluded that there may not actually be that much of a
difference between the two deployment steps. But I decided to stick with the
setup for two reasons:
*	It provided a great introduction to working with CORS.
*	Firebase Cloud Functions + Go is not an option

**GCP Cloud Functions:** *[link](https://cloud.google.com/functions/)* I
particularly like the idea of having my API managed for me on a cloud server,
as is common with most Cloud Function providers. However, I stuck with GCP here
because it offers Go as an option. Using Go as the language of my API was
perhaps the strictest requirement I had, for performance. As it turns out, it is
not an option for Cloud Functions on Firebase, at least as far as I could tell
at the time.

Another reason I chose each of these is the fact that I have never used any of
them before, much less together. This project offered an opportunity to dive
into, and learn new tools and technologies, and these were the bricks I used
to build it with.


## Detailed Design

I attempted to structure this predominantly in a parent child hierarchy, with
the blockchain being the parent, and the nodes its children. This just seemed
most natural to me. Only the blockchain has visibility into all of the nodes,
interpreting and assigning values where appropriate.


![Client Side Architecture](https://drive.google.com/uc?export=view&id=1uMoojAByt_76EiXTxnRVe5u4OGrJaM-m)


![Blockchain-Nodes](https://drive.google.com/uc?export=view&id=1Jmx5G_uN9kkdYYUMaKNm6_pgwP25aKR6)


This pattern in mimicked on the server side, illustrated via the following diagram: 


![Server Side Architecture](https://drive.google.com/uc?export=view&id=1PIE-Y5YSiPVfjv6R0sDRK19o2lBrWOFW)


Although, note the additional separation hashing of our hashing module.


## Limitations / Challenges

This entire project presented a new space for me to operate in. While I have
worked on many applications in the past, both large and small, I had never
built one of my own, much less both the client and server side of it. I also
typically prefer not to use frameworks, however Angular didn’t seem like a
terrible exception, especially given the limited time I had.

Another challenge was understanding blockchain, how it works and how it relates
to digital identity. I admit, before this assignment, the term ‘blockchain’ was
more hearsay to me than anything else. I hope I did it justice with this
implementation :).



