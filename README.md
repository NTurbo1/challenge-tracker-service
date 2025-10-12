# Challenge Tracker Service

### Requirements
- Put the storage `.csv` files inside `db/data/` directory. Otherwise the app panics.

## Data .csv file formats
* ### Users.csv:
    id,firstname,lastname,username,password
* ### Sessions.csv:
    userId,createdAt,expiresAt
* ### <useId>_<year>.csv:
    * columns: numDay,marked
    * values:
        1. numDay: day number
        2. marked:
            * 1 -> marked/done
            * 0 -> unmarked/not done
            * -1 -> blank/not counted/before the challenge start date
