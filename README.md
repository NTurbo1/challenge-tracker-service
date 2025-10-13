# Challenge Tracker Service

### Requirements
- Put the storage `.csv` files inside `db/data/` directory. Otherwise the app panics.

## Data .csv file formats
* ### users.csv: 
    `The number of users are predetermined and fixed for now, so, there's no need a specific format.`
    id,firstname,lastname,username,password
* ### sessions.csv:
    * **row length**: 48 + 1 + 8 + 1 + 19 + 1 + 19 + 1 + 1 + 1 = 100 bytes `// last byte is '\n'`
    * **columns**: id,userId,createdAt,expiresAt,valid
    * **values**:
        1. **id**: session id; length = 48 bytes
        2. **userId**: user id number; length = (int64) 8 bytes
        3. **createdAt**: time (time.DateTime format) when the session was created; length = 19 bytes
        4. **expiresAt**: time (time.DateTime format) when the session expires; length = 19 bytes
        5. **valid**: indicates if the session is still valid; length = 1 byte
* ### challenges.csv:
    * **row length**: 8 + 1 + 8 + 1 + 256 + 1 + 10 + 1 + 10 + 1 + 1 + 1 + 8 + 1 + 512 + 1 = 821 bytes `// last byte is '\n'`
    * **columns**: offset,id,name,startDate,endDate,active,userid,dataFilePath
    * **values**: 
        1. **offset**: offset of the row in the .csv file; length = (int64) 8 bytes
        2. **id**: challenge id; length = (int64) 8 bytes
        3. **name**: challenge name; length = 256 bytes
        4. **startDate**: challenge start date in time.DateOnly format; length = 10 bytes
        5. **endDate**: challenge end date in time.DateOnly format; length = 10 bytes
        6. **active**: length = 1 byte
            * 1 -> challenge is active/going
            * 0 -> challenge is inactive/completed/stopped
        7. **userId**: user id number; length = (int64) 8 bytes
        8. **dataFilePath**: filepath to a file named in a format \<userId>\_\<challengeId>_\<year>.csv;
                             length = 512 bytes
* ### \<userId>\_\<challengeId>_\<year>.csv:
    * **row length**: 8 + 1 + 1 + 1 + 8 + 1 = 20 bytes `// last byte is '\n'`
    * **columns**: numDay,marked,offset
    * **values**:
        1. **numDay**: day number; length = (int64) 8 bytes
        2. **marked**: length = 1 byte
            * 1 -> marked/done
            * 0 -> unmarked/not done
            * -1 -> blank/not counted/before the challenge start date
        3. **offset**: offset of the row in the file; length = (int64) 8 bytes
