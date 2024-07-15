
# Alumni Management Project

## Running the Application

- To run the application, ensure you have Go installed and set up.
- Also install all the required dependency and set up MYSQL in your machine.
    ```
        go get github.com/gorilla/mux v1.8.1
	    go get gorm.io/driver/mysql v1.5.7
	    go get gorm.io/gorm v1.25.10
        go get github.com/joho/godotenv v1.5.1
        go get github.com/golang-jwt/jwt v3.2.2+incompatible
    ```
- Use the following command to start the server:

    ```
        go run main.go
    ```

## .env File Structure
Create a `.env` file in the root directory of the project with the following structure:

```env
# Microsoft OAuth2 Configuration
CLIENT_ID=your-client-id
CLIENT_SECRET=your-client-secret
TENANT_ID=your-tenant-id
REDIRECT_URL="http://localhost:8000/callback"
JWT_KEY=your-secret-key

# Database Configuration
DB_NAME="alumni_db"
DB_USER=your-database-username
DB_PASS=your-database-password
DB_HOST="127.0.0.1"
DB_PORT="3306"
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
	    "ProfilePicture": "",
        "InstagramProfile": "",
	    "TwitterProfile": ""
        // if any field is not mention then it value is null
    }
    ```
- **Success Response**: `201`
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
        "InstagramProfile": "",
	    "TwitterProfile": "",
        "CreatedAt":"2024-07-04T14:06:08.561+05:30",
        "UpdatedAt":"2024-07-04T14:06:08.561+05:30",
    }
    ```

#### Get All Alumni Profiles
- **URL**: `/alumni`
- **Method**: `GET`
- **Function**: `to get details of all alumni`
- **Parameter**: `No Parameter`
- **Success Response**: `200`
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
            "InstagramProfile": "",
	        "TwitterProfile": "",
            "CreatedAt":"2024-07-04T14:06:08.561+05:30",
            "UpdatedAt":"2024-07-04T14:06:08.561+05:30"
        }
    ]
    ```

#### Get an Alumni Profile by ID
- **URL**: `/alumni/{id}`
- **Method**: `GET`
- **Function**: `to get details of an particular alumni by its ID`
- **Parameter**: `here id represent AlumniID`
- **Success Response**: `200`
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
        "InstagramProfile": "",
	    "TwitterProfile": "",
        "CreatedAt":"2024-07-02T16:02:34.902+05:30",
        "UpdatedAt":"2024-07-02T16:02:34.902+05:30"
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
- **Success Response**: `200`
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
        "InstagramProfile": null,
	    "TwitterProfile": null,
        "CreatedAt":"2024-07-02T16:02:34.902+05:30",
        "UpdatedAt":"2024-07-04T18:38:02.965+05:30"
    }
    ```

#### Delete an Alumni Profile
- **URL**: `/alumni/{id}`
- **Method**: `DELETE`
- **Function**: `to delete a alumni`
- **Parameter**: `Here id represent AlumniID`
- **Response**: `204 No Content`


---

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
- **Success Response**: `201`
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
        "UpdatedAt":"2024-07-04T19:03:17.431+05:30"
    }
    ```

#### Get All Events
- **URL**: `/event`
- **Method**: `GET`
- **Function**: `return the list of all event`
- **Parameter**: `No Parameter`
- **Success Response**: `200`
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
            "UpdatedAt":"2024-07-02T17:04:58.344+05:30"
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
            "UpdatedAt":"2024-07-04T19:03:17.431+05:30"
        }
    ]
    ```

#### Get an Event by ID
- **URL**: `/event/{id}`
- **Method**: `GET`
- **Function**: `return specific event`
- **Parameter**: `Here id represent EventID`
- **Success Response**: `200`
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
        "UpdatedAt":"2024-07-04T19:03:17.431+05:30"
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
- **Success Response**: `200`
    ```json
    {
        "EventID":5,
        "Title":"Annual Alumni Meetup",
        "Description":"Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
        "EventType":"Networking",
        "ModeOfEvent":"Virtual","Location":"Google Meet",
        "EventDateTime":"2024-10-28T10:00:00Z",
        "CreatedAt":"2024-07-04T19:03:17.431+05:30",
        "UpdatedAt":"2024-07-04T19:09:26.716+05:30"
    }
    ```

#### Delete an Event
- **URL**: `/event/{id}`
- **Method**: `DELETE`
- **Function**: `used to delete an event`
- **Parameter**: `Here id represent EventID`
- **Success Response**: `204 No Content`

---
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
- **Success Response**: `201`
    ```json
    {
        "ProfID":5,
        "AlumniID":2,
        "CompanyName":"Tech Solutions Inc.",
        "Position":"Software Engineer",
        "Duration":"2020 - Present",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-04T19:14:09.058+05:30",
        "UpdatedAt":"2024-07-04T19:14:09.058+05:30"
    }
    ```

#### Get All Professional Informations
- **URL**: `/professionalInfo`
- **Method**: `GET`
- **Function**: `return the list of all professional Informations`
- **Parameter**: `No Parameter`
- **Success Response**: `200`
    ```json
    [
        {
            "ProfID":1,
            "AlumniID":0,
            "CompanyName":"Tech Solutions Inc.",
            "Position":"Software Engineer",
            "Duration":"2020 - Present",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:10:43.344+05:30",
            "UpdatedAt":"2024-07-02T16:10:43.344+05:30"
        },
        {
            "ProfID":2,
            "AlumniID":2,
            "CompanyName":"Microsoft",
            "Position":"Software Engineer Intern",
            "Duration":"2020 - 2022",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:11:25.505+05:30",
            "UpdatedAt":"2024-07-02T16:11:25.505+05:30"
        }
    ]
    ```

#### Get All Professional Information by Alumni ID
- **URL**: `/professionalInfo/{id}`
- **Method**: `GET`
- **Function**: `return the list of all professionalInfo of an alumni`
- **Parameter**: `url accept the params id here id represent AlumniID`
- **Success Response**: `200`
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
            "UpdatedAt":"2024-07-03T16:17:13.509+05:30"
        },
        {
            "ProfID":5,
            "AlumniID":2,
            "CompanyName":"Tech Solutions Inc.",
            "Position":"Software Engineer",
            "Duration":"2020 - Present",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
            "CreatedAt":"2024-07-04T19:14:09.058+05:30",
            "UpdatedAt":"2024-07-04T19:14:09.058+05:30"
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
- **Success Response**: `200`
    ```json
    {
        "ProfID":4,
        "AlumniID":2,
        "CompanyName":"Tech Solutions Inc.",
        "Position":"Software Engineer",
        "Duration":"2020 - 2022",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-02T16:43:50.428+05:30",
        "UpdatedAt":"2024-07-04T19:22:37.061+05:30"
    }
    ```

#### Delete Professional Information
- **URL**: `/professionalInfo/{id}`
- **Method**: `DELETE`
- **Function**: `used to delete the professionalInfo of an alumni`
- **Parameter**: `url accept the params id here id represent profID`
- **Success Response**: `204 No Content`

---
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
- **Success Response**: `201`
    ```json
    {
        "AchievementID":7,
        "AlumniID":1,
        "Title":"First Prize in Coding Competition",
        "Description":"Won first prize in the national coding competition organized by ACM.",
        "DateAchieved":"2023-05-20T15:04:05Z",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null},
        "CreatedAt":"2024-07-04T19:24:59.596+05:30",
        "UpdatedAt":"2024-07-04T19:24:59.596+05:30"
    }
    ```

#### Get All Achievements
- **URL**: `/achievement`
- **Method**: `GET`
- **Function**: `return the list of all achievement`
- **Parameter**: `No Parameter`
- **Success Response**: `200`
    ```json
    [
        {
            "AchievementID":1,
            "AlumniID":1,
            "Title":"Winner of Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:39:00.826+05:30",
            "UpdatedAt":"2024-07-04T19:29:20.524+05:30"
        },
        {
            "AchievementID":2,
            "AlumniID":2,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:39:46.213+05:30",
            "UpdatedAt":"2024-07-02T16:39:46.213+05:30"
        },
        {
            "AchievementID":3,
            "AlumniID":1,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.","DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:41:20.823+05:30",
            "UpdatedAt":"2024-07-02T16:41:20.823+05:30"
        }
    ]
    ```

#### Get All Achievements by Alumni ID
- **URL**: `/achievement/{id}`
- **Method**: `GET`
- **Function**: `used to list all achivements of an alumni`
- **Parameter**: `url accept the params id here id represent AlumniID`
- **Success Response**: `200`
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
        },
        {
            "AchievementID":3,
            "AlumniID":1,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:41:20.823+05:30",
            "UpdatedAt":"2024-07-02T16:41:20.823+05:30"
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
- **Success Response**: `200`
    ```json
    {
        "AchievementID":1,
        "AlumniID":1,
        "Title":"Winner of Coding Competition",
        "Description":"Won first prize in the national coding competition organized by ACM.",
        "DateAchieved":"2023-05-20T20:34:05+05:30",
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
        "CreatedAt":"2024-07-02T16:39:00.826+05:30",
        "UpdatedAt":"2024-07-04T19:29:20.524+05:30",
    }
    ```

#### Delete Achievement

* **URL**: `/achievement/{id}`
* **Method**: `DELETE`
* **Description**: Deletes an Achievement record by its ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the Achievement to delete.
* **Success Response**:
    * **Code**: `204 No Content`
    * **Content**: Empty response body.

---

### Interest Hobbies

#### Add Interest Hobby

* **URL**: `/interesthobbies`
* **Method**: `POST`
* **Description**: Creates a new interest or hobby for an alumni.
* **Request Body**:
    ```json
    {
        "AlumniID": 1,
        "InterestHobby": "Listening rap"
    }
    ```
* **Success Response**:
    * **Code**: `201 Created`
    * **Content**:
      ```json
      {
          "InterestID": 1,
          "AlumniID": 1,
          "InterestHobby": "Listening rap",
          "CreatedAt": "2024-07-06T12:00:00Z",
          "UpdatedAt": "2024-07-06T12:00:00Z"
      }
      ```


#### Update Interest Hobby

* **URL**: `/interesthobbies/{id}`
* **Method**: `PUT`
* **Description**: Updates an existing interest or hobby by its ID.
* **Request Body**:
    ```json
    {
        "InterestHobby": "Playing Videogames"
    }
    ```
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the interest hobby to update.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      {
          "InterestID": 1,
          "AlumniID": 1,
          "InterestHobby": "Playing Videogames",
          "CreatedAt": "2024-07-06T12:00:00Z",
          "UpdatedAt": "2024-07-06T12:00:00Z"
      }
      ```


#### Delete Interest Hobby

* **URL**: `/interesthobbies/{id}`
* **Method**: `DELETE`
* **Description**: Deletes an interest or hobby by its ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the interest hobby to delete.
* **Success Response**:
    * **Code**: `204 No Content`
    * **Content**: Empty response body.



#### Get All Interest Hobbies by Alumni ID

* **URL**: `/interesthobbies/alumni/{id}`
* **Method**: `GET`
* **Description**: Retrieves all interest hobbies for a specific alumni by their ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the alumni.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      [
          {
              "InterestID": 1,
              "AlumniID": 1,
              "InterestHobby": "Listening rap",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          },
          {
              "InterestID": 2,
              "AlumniID": 1,
              "InterestHobby": "Playing football",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          }
      ]
      ```

---

### Interview Experiences

#### Add Interview Experience

* **URL**: `/interviewexperiences`
* **Method**: `POST`
* **Description**: Creates a new interview experience record for an alumni.
* **Request Body**:
    ```json
    {
        "AlumniID": 1,
        "CompanyName": "Google",
        "JobTitle": "Backend Golang Developer."
        "Description": "first round is coding Round and second is ...",
        "InterviewDate": "2024-07-01T09:00:00Z"
    }
    ```
* **Success Response**:
    * **Code**: `201 Created`
    * **Content**:
      ```json
      {
          "ExperienceID": 1,
          "AlumniID": 1,
          "CompanyName": "Google",
          "JobTitle": "Backend Golang Developer."
          "Description": "first round is coding Round and second is ...",
          "InterviewDate": "2024-07-01T09:00:00Z",
          "CreatedAt": "2024-07-06T12:00:00Z",
          "UpdatedAt": "2024-07-06T12:00:00Z"
      }
      ```

#### Get All Interview Experience
- **URL**: `/interviewexperiences`
- **Method**: `GET`
- **Description**: `return the list of all Interview Experience`
- **Success Response**: `200 Ok`
    ```json
    [
        {
              "ExperienceID": 1,
              "AlumniID": 1,
              "CompanyName": "Google",
              "JobTitle": "Backend Golang Developer.",
              "Description": "first round is coding Round and second is ...",
              "InterviewDate": "2024-07-01T09:00:00Z",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          },
          {
              "ExperienceID": 2,
              "AlumniID": 1,
              "CompanyName": "Amazon",
              "JobTitle": "MERN stack developer",
              "Description": "First round is DSA round 2nd round is project showcase 3rd is interview...",
              "InterviewDate": "2024-06-15T11:00:00Z",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          }
    ]
    ```

#### Update Interview Experience

* **URL**: `/interviewexperiences/{id}`
* **Method**: `PUT`
* **Description**: Updates an existing interview experience record by its ID.
* **Request Body**:
    ```json
    {
        "CompanyName": "Amazon",
        "Description": "First round is DSA round 2nd round is project showcase 3rd is interview...",
        "InterviewDate": "2024-08-01T10:00:00Z"
    }
    ```
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the interview experience to update.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      {
          "ExperienceID": 1,
          "AlumniID": 1,
          "CompanyName": "Amazon",
          "JobTitle": "MERN stack developer",
          "Description": "First round is DSA round 2nd round is project showcase 3rd is interview...",
          "InterviewDate": "2024-08-01T10:00:00Z",
          "CreatedAt": "2024-07-06T12:00:00Z",
          "UpdatedAt": "2024-07-06T12:00:00Z"
      }
      ```



#### Delete Interview Experience

* **URL**: `/interviewexperiences/{id}`
* **Method**: `DELETE`
* **Description**: Deletes an interview experience record by its ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the interview experience to delete.
* **Success Response**:
    * **Code**: `204 No Content`
    * **Content**: Empty response body.



#### Get All Interview Experiences by Alumni ID

* **URL**: `/interviewexperiences/alumni/{id}`
* **Method**: `GET`
* **Description**: Retrieves all interview experiences for a specific alumni by their ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the alumni.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      [
          {
              "ExperienceID": 1,
              "AlumniID": 1,
              "CompanyName": "Google",
              "JobTitle": "Backend Golang Developer.",
              "Description": "first round is coding Round and second is ...",
              "InterviewDate": "2024-07-01T09:00:00Z",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          },
          {
              "ExperienceID": 2,
              "AlumniID": 1,
              "CompanyName": "Amazon",
              "JobTitle": "MERN stack developer",
              "Description": "First round is DSA round 2nd round is project showcase 3rd is interview...",
              "InterviewDate": "2024-06-15T11:00:00Z",
              "CreatedAt": "2024-07-06T12:00:00Z",
              "UpdatedAt": "2024-07-06T12:00:00Z"
          }
      ]
      ```

---
### AlumniAttending

#### Create an Alumni Attending Record
* **URL**: `/alumniattending`
* **Method**: `POST`
* **Description**: Creates a new alumni attending record.
* **Request Body**:
    ```json
    {
        "EventID" : 3,
	    "AlumniID" : 8
    }
    ```
* **Success Response**:
    * **Code**: `201 Created`
    * **Content**:
    ```json
    {
        "AttendID":4,
        "EventID":3,
        "AlumniID":8,
        "Event":{"EventID":0,"Title":"","Description":"","EventType":"","ModeOfEvent":"","Location":"","EventDateTime":"0001-01-01T00:00:00Z","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
        "CreatedAt":"2024-07-09T17:41:31.794+05:30",
        "UpdatedAt":"2024-07-09T17:41:31.794+05:30"
    }
    ```

#### Get All Alumni by EventID

* **URL**: `/alumniattending/event/{id}`
* **Method**: `GET`
* **Description**: Retrieves all Alumni Details who are attending specific event by help of EventID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the Event.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
    ```json
    [
        {
            "AlumniID": 8,
            "FirstName": "Mohit",
            "LastName": "Gusain",
            "Branch": "CSE",
            "BatchYear": 2022,
            "MobileNo": "96697nn696",
            "Email": "mohitgn871@gmail.com",
            "EnrollmentNo": "08529702722",
            "Tenth": "89.5",
            "Xllth": "",
            "Degree": "B.tech(CSE)",
            "GithubProfile": "github.com/mohitgusain8671",
    "LeetCodeProfile": "leetcode.com/mohitgusain8671",
            "LinkedInProfile": "linkedin.com",
            "CodeforceProfile": "link.com",
            "CodeChefProfile": "",
            "InstagramProfile": null,
            "TwitterProfile": null,
            "ProfilePicture": "",
            "CreatedAt": "2024-07-04T14:06:08.561+05:30",
            "UpdatedAt": "2024-07-04T14:06:08.561+05:30"
        },
        {
            "AlumniID": 2,
            "FirstName": "Vikas",
            "LastName": "",
            "Branch": "CSE",
            "BatchYear": 2022,
            "MobileNo": "96684436456",
            "Email": "vikas60@gmail.com",
            "EnrollmentNo": "08220802722",
            "Tenth": "54",
            "Xllth": "84",
            "Degree": "B.tech(CSE)",
            "GithubProfile": "github.com/sahil0603",
            "LeetCodeProfile": "leetcode.com/sahil0603",
            "LinkedInProfile": "linkedin.com/sahilChauhan",
            "CodeforceProfile": null,
            "CodeChefProfile": null,
            "InstagramProfile": null,
            "TwitterProfile": null,
            "ProfilePicture": "",
            "CreatedAt": "2024-07-02T16:02:34.902+05:30",
            "UpdatedAt": "2024-07-04T18:38:02.965+05:30"
        }
    ]

    ```


#### Get All Event by AlumniID

* **URL**: `/alumniattending/alumni/{id}`
* **Method**: `GET`
* **Description**: Retrieves all Event Details who are attending specific event by help of EventID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the Alumni.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
    ```json
    [
        {
            "EventID": 1,
            "Title": "Annual Alumni Meetup",
            "Description": "Join us for our annual alumni meetup where we   reconnect and reminisce about our college days.",
            "EventType": "Networking",
            "ModeOfEvent": "Virtual",
            "Location": "Google Meet",
            "EventDateTime": "2024-10-28T15:30:00+05:30",
            "CreatedAt": "2024-07-02T16:03:19.83+05:30",
            "UpdatedAt": "2024-07-02T16:50:09.953+05:30"
        },
        {
            "EventID": 3,
            "Title": "Annual Alumni Meetup",
            "Description": "Join us for our annual alumni meetup where we reconnect and reminisce about our college days.",
            "EventType": "Networking",
            "ModeOfEvent": "Virtual",
            "Location": "Online",
            "EventDateTime": "2024-08-15T15:30:00+05:30",
            "CreatedAt": "2024-07-02T17:04:58.344+05:30",
            "UpdatedAt": "2024-07-02T17:04:58.344+05:30"
        }
    ]

    ```




#### Update AlumniAttending

* **URL**: `/alumniattending/{id}`
* **Method**: `PUT`
* **Description**: Updates an existing AlumniAttending record by its ID.
* **Request Body**:
    ```json
    {
        "EventID" : 3,
	    "AlumniID" : 8
    }
    ```
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the AlumniAttending record to update.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
    ```json
    {
        "AttendID":1,
        "EventID":3,
        "AlumniID":8,
        "Event":{"EventID":0,"Title":"","Description":"","EventType":"","ModeOfEvent":"","Location":"","EventDateTime":"0001-01-01T00:00:00Z","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
        "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
        "CreatedAt":"2024-07-09T17:20:55.294+05:30",
        "UpdatedAt":"2024-07-09T17:35:59.294+05:30"
    }
      ```



#### Delete AlumniAttending

* **URL**: `/alumniattending/{id}`
* **Method**: `DELETE`
* **Description**: Deletes an AlumniAttending record by its ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the AlumniAttending Record to delete.
* **Success Response**:
    * **Code**: `204 No Content`
    * **Content**: Empty response body.


---

## Database Migrations

The API will automatically create the necessary database tables if they do not exist when an endpoint is accessed.

## Error Handling

The API will return appropriate HTTP status codes and error messages for invalid requests or server errors.

- `400 Bad Request`: The request was invalid.
- `404 Not Found`: The requested record was not found.
- `500 Internal Server Error`: An error occurred on the server.



