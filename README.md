#### Objectives

Visit N(number of websites) if the website is working or sending some response on determined amount of time mark it as working website. 

If Website failed to respond then mark it as non-working website. 

At the receiving end of the channel  
Display the statistics of the result like 

Successful count => 45  
UnSuccessful count => 15

#### Description 

##### Pervious version was blocking so it was slow. It lacks concurrency. 

In this version we don't wait for one operations to finish we don't defer our worker thread. 

ALL error or success response is passed to the done channel and we loops the channel to get count. 

