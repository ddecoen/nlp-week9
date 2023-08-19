# nlp-week9
Create app for natural language processing with wails and sql database

## Use Wails to Develop and Build an App
1. Follow the installation instructions for [Wails](https://wails.io/docs/gettingstarted/installation)
2. Pick a method for app development for frontend  
    a. Svelte (what I selected)  
    b. Type Script  
3. Leverage the logic provided by Dr. Miller to create a SQL database (https://github.com/ThomasWMiller/jump-start-sqlite/tree/main)  
4. Once the database is created, it needs to be embedded in the Wails app  
    a. Hat tip to Wes West for the suggestion. See his example code below:  
    ![Example](/Example.png)  
    b. As I was still having some problems, I just created the database with the coded provided by Dr. Miller in number 3 above. I then copied that DB into the backend folder so it could be queried.  
    c. Made sure instructions were added to the front-end so the user could accurately access the data in the database.  

## Shortcomings of current process  
While the DB created was only 8KB, I am not sure that embedded is the best way to go. On the positive side, this allows for the app to work in the Wails build. In the previous assignment, I could not get my app to build, so I had to stay in Wails Dev.  

## Future Steps - What comes after this MVP?  
As this little app is good for an introduction, it would need some work to continue to scale. One possible suggestion would be to leverage a cloud database (e.g., Amazon Redshift or Google BigQuery) or a hosted database. It would also be nice if the app had a component that allowed it to learn so that it could respond to close enough queries. One possiblity would be to use ElasticSearch to return responses for fuzzy queries. I think the best path forward would be to leverage microservices. I do believe that most of the final project could be completed in Go, but it would be the choice of the firm to determine which tools to ulitmately use. 
