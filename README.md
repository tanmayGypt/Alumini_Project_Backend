
# Alumni Management Project

## Running the Application

- To run the application, ensure you have Go installed and set up.
- Also install all the required dependency and set up MYSQL in your machine.
    ```
        go get github.com/gorilla/mux v1.8.1
	    go get gorm.io/driver/mysql v1.5.7
	    go get gorm.io/gorm v1.25.10
    ```
- Use the following command to start the server:

```
go run main.go
```

## Alumni Management API

Working of API's for managing alumni profiles, events, professional information, and achievements etc are given below.

## Endpoints

### Alumni Profiles

#### Create an Alumni Profile
- **URL**: `/alumni`
- **Method**: `POST`
- **Function**: `to create an alumni`
- **Parameter**: `No Parameter`
- **Request Body**:
    ```json
    {
        "FirstName": "Mohit", 
	    "LastName": "Gusain", 
	    "Branch": "CSE", 
        "BatchYear": 2022,
	    "MobileNo": "966979696",//unique
        "Email": "mohitgn8671@gmail.com",//unique
	    "EnrollmentNo": "08520702722",//unique
	    "Tenth": "89.5",
	    "Xllth": "90",
	    "Degree": "B.tech(CSE)",
	    "GithubProfile": "github.com/mohitgusain8671",
	    "LeetCodeProfile": "leetcode.com/mohitgusain8671",
	    "LinkedInProfile": "linkedin.com",
        "CodeforceProfile": "link.com",
	    "CodeChefProfile": "",
	    "ProfilePicture": ""
        // if any field is not mention then it value is null
    }
    ```
- **Response**: `201`
    ```json
    {
        "AlumniID":8,
        "FirstName":"Mohit",
        "LastName":"Gusain",
        "Branch":"CSE",
        "BatchYear":2022,
        "MobileNo":"96697nn696",
        "Email":"mohitgn871@gmail.com",
        "EnrollmentNo":"08529702722",
        "Tenth":"89.5",
        "Xllth":"",
        "Degree":"B.tech(CSE)",
        "GithubProfile":"github.com/mohitgusain8671",
        "LeetCodeProfile":"leetcode.com/mohitgusain8671",
        "LinkedInProfile":"linkedin.com",   
        "CodeforceProfile":"link.com",
        "CodeChefProfile":"",
        "ProfilePicture":"",
        "CreatedAt":"2024-07-04T14:06:08.561+05:30",
        "UpdatedAt":"2024-07-04T14:06:08.561+05:30",
        "DeletedAt":null
    }
    ```

#### Get All Alumni Profiles
- **URL**: `/alumni`
- **Method**: `GET`
- **Function**: `to get details of all alumni`
- **Parameter**: `No Parameter`
- **Response**: `200`
    ```json
    [
        {
            "AlumniID":8,
            "FirstName":"Mohit",
            "LastName":"Gusain",
            "Branch":"CSE",
            "BatchYear":2022,
            "MobileNo":"96697nn696",
            "Email":"mohitgn871@gmail.com",
            "EnrollmentNo":"08529702722",
            "Tenth":"89.5",
            "Xllth":"",
            "Degree":"B.tech(CSE)",
            "GithubProfile":"github.com/mohitgusain8671",
            "LeetCodeProfile":"leetcode.com/mohitgusain8671",
            "LinkedInProfile":"linkedin.com",   
            "CodeforceProfile":"link.com",
            "CodeChefProfile":"",
            "ProfilePicture":"",
            "CreatedAt":"2024-07-04T14:06:08.561+05:30",
            "UpdatedAt":"2024-07-04T14:06:08.561+05:30",
            "DeletedAt":null
        }
    ]
    ```

#### Get an Alumni Profile by ID
- **URL**: `/alumni/{id}`
- **Method**: `GET`
- **Function**: `to get details of an particular alumni by its ID`
- **Parameter**: `here id represent AlumniID`
- **Response**: `200`
    ```json
    {
        "AlumniID":2,
        "FirstName":"Sahil",
        "LastName":"Chauham",
        "Branch":"CSE",
        "BatchYear":2022,
        "MobileNo":"9839493490",
        "Email":"sahil0603@gmail.com",
        "EnrollmentNo":"08220802722",
        "Tenth":"95.5",
        "Xllth":"94",
        "Degree":"B.tech(CSE)",
        "GithubProfile":"github.com/sahil0603",
        "LeetCodeProfile":"leetcode.com/sahil0603",
        "LinkedInProfile":"linkedin.com/sahilChauhan",
        "CodeforceProfile":null,
        "CodeChefProfile":null,
        "ProfilePicture":null,
        "CreatedAt":"2024-07-02T16:02:34.902+05:30",
        "UpdatedAt":"2024-07-02T16:02:34.902+05:30",
        "DeletedAt":null
    }
    ```

#### Update an Alumni Profile
- **URL**: `/alumni/{id}`
- **Method**: `PUT`
- **Function**: `to update the details of an alumni`
- **Parameter**: `Here id represent AlumniID`
- **Request Body**:
    ```json
    {
        "AlumniID":2,
        "FirstName":"Vikas",
        "LastName":"",
        "MobileNo":"96684436456",
        "Email":"vikas60@gmail.com",
        "Tenth":"54",
        "Xllth":"84",
        "CodeChefProfile":null
        // only put the item you want to update
    }
    ```
- **Response**: `200`
    ```json
    {
        "AlumniID":2,
        "FirstName":"Vikas",
        "LastName":"",
        "Branch":"CSE",
        "BatchYear":2022,
        "MobileNo":"96684436456",
        "Email":"vikas60@gmail.com",
        "EnrollmentNo":"08220802722",
        "Tenth":"54",
        "Xllth":"84",
        "Degree":"B.tech(CSE)",
        "GithubProfile":"github.com/sahil0603",
        "LeetCodeProfile":"leetcode.com/sahil0603",
        "LinkedInProfile":"linkedin.com/sahilChauhan",
        "CodeforceProfile":null,
        "CodeChefProfile":null,
        "ProfilePicture":null,
        "CreatedAt":"2024-07-02T16:02:34.902+05:30",
        "UpdatedAt":"2024-07-04T18:38:02.965+05:30",
        "DeletedAt":null
    }
    ```

#### Delete an Alumni Profile
- **URL**: `/alumni/{id}`
- **Method**: `DELETE`
- **Function**: `to delete a alumni`
- **Parameter**: `Here id represent AlumniID`
- **Response**: `204 No Content`

### Events

#### Create a New Event
- **URL**: `/event`
- **Method**: `POST`
- **Function**: `to create a new event`
- **Parameter**: `No Parameter`
- **Request Body**:
    ```json
    {
        "title": "Annual Alumni Meetup",
        "description": "Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
        "eventType": "Networking",
        "modeOfEvent": "Virtual",
        "location": "Online",
        "eventDateTime": "2024-08-15T10:00:00Z"
    }
    ```
- **Response**: `201`
    ```json
    {
        "EventID":5,
        "Title":"Annual Alumni Meetup",
        "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
        "EventType":"Networking",
        "ModeOfEvent":"Virtual",
        "Location":"Online",
        "EventDateTime":"2024-08-15T10:00:00Z",
        "CreatedAt":"2024-07-04T19:03:17.431+05:30",
        "UpdatedAt":"2024-07-04T19:03:17.431+05:30",
        "DeletedAt":null
    }
    ```

#### Get All Events
- **URL**: `/event`
- **Method**: `GET`
- **Function**: `return the list of all event`
- **Parameter**: `No Parameter`
- **Response**: `200`
    ```json
    [
        {
            "EventID":3,
            "Title":"Annual Alumni Meetup",
            "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
            "EventType":"Networking",
            "ModeOfEvent":"Virtual",
            "Location":"Online",
            "EventDateTime":"2024-08-15T15:30:00+05:30",
            "CreatedAt":"2024-07-02T17:04:58.344+05:30",
            "UpdatedAt":"2024-07-02T17:04:58.344+05:30",
            "DeletedAt":null
        },
        {
            "EventID":5,
            "Title":"Annual Alumni Meetup",
            "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
            "EventType":"Networking",
            "ModeOfEvent":"Virtual",
            "Location":"Online",
            "EventDateTime":"2024-08-15T15:30:00+05:30",
            "CreatedAt":"2024-07-04T19:03:17.431+05:30",
            "UpdatedAt":"2024-07-04T19:03:17.431+05:30",
            "DeletedAt":null
        }
    ]
    ```

#### Get an Event by ID
- **URL**: `/event/{id}`
- **Method**: `GET`
- **Function**: `return specific event`
- **Parameter**: `Here id represent EventID`
- **Response**: `200`
    ```json
    {
        "EventID":5,
        "Title":"Annual Alumni Meetup",
        "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
        "EventType":"Networking",
        "ModeOfEvent":"Virtual",
        "Location":"Online",
        "EventDateTime":"2024-08-15T15:30:00+05:30",
        "CreatedAt":"2024-07-04T19:03:17.431+05:30",
        "UpdatedAt":"2024-07-04T19:03:17.431+05:30",
        "DeletedAt":null
    }
    ```

#### Update an Event
- **URL**: `/event/{id}`
- **Method**: `PUT`
- **Function**: `to update the details of an event`
- **Parameter**: `Here id represent EventID`
- **Request Body**:
    ```json
    {
        "location": "Google Meet",
        "eventDateTime": "2024-10-28T10:00:00Z"
        // only put the item you want to update
    }
    ```
- **Response**: `200`
    ```json
    {
        "EventID":5,
        "Title":"Annual Alumni Meetup",
        "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
        "EventType":"Networking",
        "ModeOfEvent":"Virtual","Location":"Google Meet",
        "EventDateTime":"2024-10-28T10:00:00Z",
        "CreatedAt":"2024-07-04T19:03:17.431+05:30",
        "UpdatedAt":"2024-07-04T19:09:26.716+05:30",
        "DeletedAt":null
    }
    ```

#### Delete an Event
- **URL**: `/event/{id}`
- **Method**: `DELETE`
- **Function**: `used to delete an event`
- **Parameter**: `Here id represent EventID`
- **Response**: `204 No Content`


### Professional Information

#### Add Professional Information
- **URL**: `/professionalInfo`
- **Method**: `POST`
- **Function**: `used to add professionalInfo of an alumni`
- **Parameter**: `No Parameter`
- **Request Body**:
    ```json
    {
        "AlumniID": 2,
        "companyName": "Tech Solutions Inc.",
        "position": "Software Engineer",
        "duration": "2020 - Present"
    }
    ```
- **Response**: `201`
    ```json
    {
        "ProfID":5,
        "AlumniID":2,
        "CompanyName":"Tech Solutions Inc.",
        "Position":"Software Engineer",
        "Duration":"2020 - Present",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-04T19:14:09.058+05:30",
        "UpdatedAt":"2024-07-04T19:14:09.058+05:30",
        "DeletedAt":null
    }
    ```

#### Get All Professional Information by Alumni ID
- **URL**: `/professionalInfo/{id}`
- **Method**: `GET`
- **Function**: `return the list of all professionalInfo of an alumni`
- **Parameter**: `url accept the params id here id represent AlumniID`
- **Response**: `200`
    ```json
    [
        {
            "ProfID":4,
            "AlumniID":2,
            "CompanyName":"Tech Solutions Inc.",
            "Position":"Software Engineer intern",
            "Duration":"2020 - 2023",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}, 
            "CreatedAt":"2024-07-02T16:43:50.428+05:30",
            "UpdatedAt":"2024-07-03T16:17:13.509+05:30",
            "DeletedAt":null
        },
        {
            "ProfID":5,
            "AlumniID":2,
            "CompanyName":"Tech Solutions Inc.",
            "Position":"Software Engineer",
            "Duration":"2020 - Present",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
            "CreatedAt":"2024-07-04T19:14:09.058+05:30",
            "UpdatedAt":"2024-07-04T19:14:09.058+05:30",
            "DeletedAt":null
        }
    ]
    ```

#### Update Professional Information
- **URL**: `/professionalInfo/{id}`
- **Method**: `PUT`
- **Function**: `used to update a professionalInfo of an alumni`
- **Parameter**: `url accept the params id here id represent profID`
- **Request Body**:
    ```json
    {
        "companyName": "Tech Solutions Inc.",
        "position": "Software Engineer",
        "duration": "2020 - 2022"
        // only put the item you want to update
    }
    ```
- **Response**: `200`
    ```json
    {
        "ProfID":4,
        "AlumniID":2,
        "CompanyName":"Tech Solutions Inc.",
        "Position":"Software Engineer",
        "Duration":"2020 - 2022",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-02T16:43:50.428+05:30",
        "UpdatedAt":"2024-07-04T19:22:37.061+05:30",
        "DeletedAt":null
    }
    ```

#### Delete Professional Information
- **URL**: `/professionalInfo/{id}`
- **Method**: `DELETE`
- **Function**: `used to delete the professionalInfo of an alumni`
- **Parameter**: `url accept the params id here id represent profID`
- **Response**: `204 No Content`

### Achievements

#### Add Achievements
- **URL**: `/achievement`
- **Method**: `POST`
- **Function**: `used to add achievement of an alumni`
- **Parameter**: `No Parameter`
- **Request Body**:
    ```json
    {
        "AlumniID": 1,
        "title": "First Prize in Coding Competition",
        "description": "Won first prize in the national coding competition organized by ACM.",
        "dateAchieved": "2023-05-20T15:04:05Z"
    }
    ```
- **Response**: `201`
    ```json
    {
        "AchievementID":7,
        "AlumniID":1,
        "Title":"First Prize in Coding Competition",
        "Description":"Won first prize in the national coding competition organized by ACM.",
        "DateAchieved":"2023-05-20T15:04:05Z",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-04T19:24:59.596+05:30",
        "UpdatedAt":"2024-07-04T19:24:59.596+05:30",
        "DeletedAt":null
    }
    ```

#### Get All Achievements by Alumni ID
- **URL**: `/achievement/{id}`
- **Method**: `GET`
- **Function**: `used to list all achivements of an alumni`
- **Parameter**: `url accept the params id here id represent AlumniID`
- **Response**: `200`
    ```json
    [
        {
            "AchievementID":1,
            "AlumniID":1,
            "Title":"First position in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
            "CreatedAt":"2024-07-02T16:39:00.826+05:30",
            "UpdatedAt":"2024-07-03T16:55:15.801+05:30",
            "DeletedAt":null
        },
        {
            "AchievementID":3,
            "AlumniID":1,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
            "CreatedAt":"2024-07-02T16:41:20.823+05:30",
            "UpdatedAt":"2024-07-02T16:41:20.823+05:30",
            "DeletedAt":null
        }
    ]
    ```

#### Update Achievement Information
- **URL**: `/achievement/{id}`
- **Method**: `PUT`
- **Function**: `used to update achievements of an alumni`
- **Parameter**: `url accept the params id here id represent AchievementID`
- **Request Body**:
    ```json
    {
        "AlumniID":1,
        "Title":"Winner of Coding Competition"
        // only put the item you want to update
    }
    ```
- **Response**: `200`
    ```json
    {
        "AchievementID":1,
        "AlumniID":1,
        "Title":"Winner of Coding Competition",
        "Description":"Won first prize in the national coding competition organized by ACM.",
        "DateAchieved":"2023-05-20T20:34:05+05:30",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-02T16:39:00.826+05:30",
        "UpdatedAt":"2024-07-04T19:29:20.524+05:30",
        "DeletedAt":null
    }
    ```

## Database Migrations

The API will automatically create the necessary database tables if they do not exist when an endpoint is accessed.

## Error Handling

The API will return appropriate HTTP status codes and error messages for invalid requests or server errors.

- `400 Bad Request`: The request was invalid.
- `404 Not Found`: The requested record was not found.
- `500 Internal Server Error`: An error occurred on the server.



