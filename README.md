#### Objectives

Visit N(number of websites) if the website is working or sending some response on determined amount of time mark it as working website. 

If Website failed to respond then mark it as non-working website. 

At the receiving end of the channel  
Display the statistics of the result like 

Successful count => 45  
UnSuccessful count => 15

#### Description 

* `checkStatus` function returns a channel that can be read from to retrieve results of an iteration of our loop

* We make results channel with make keyword. 

* We create another anyomous fuction which loops over the supplied Website url's.

* Then the result of out put is assigned to `result` type.

* We then write the `res` result to channel.

* We then loops to `checkStatus` function and count the success and failure's. And print result to console.