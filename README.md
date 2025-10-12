# Challenge Tracker Service

### Requirements
- Put the storage `.csv` files inside `db/data/` directory. Otherwise the app panics.

## Data .csv file formats
* ### Users.csv:
    id,firstname,lastname,username,password
* ### Sessions.csv:
    userId,createdAt,expiresAt
* ### <useId>_<year>.csv:
    numDay,marked
