
# Alumni Management Project

## Running the Application

- To run the application, ensure you have Go installed and set up.
- Also install all the required dependency and set up PostgresSQL in your machine.
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
JWT_KEY=your-secret-key

# Database Configuration
DB_NAME="alumni_db"
DB_USER=your-database-username
DB_PASS=your-database-password
DB_HOST="localhost"
DB_PORT="5432"
SSL_MODE=your-ssl-mode

#smtp credentials
SMTP_SERVER = smtp.example.com
SMTP_PORT = port-no.
EMAIL_USER = senders-email
EMAIL_PASSWORD = sender-email-password
```

## Alumni Management API

Working of API's for managing alumni profiles, events, professional information, and achievements etc are given below.

## Endpoints

## Authentication


### Login

- **URL**: `/login`
- **Method**: `POST`
- **Description**: Authenticate and log in an alumni user, returning a JWT token and storing it in cookies.
- **Request Body**:
    ```json
    {
        "Email": "sahil@go.dev",
        "Password": "password"
    }
    ```
- **Success Response**: `200 OK`
    ```json
    {
        "message": "Login successful"
    }
    ```

### Signup

#### Create an Alumni Profile (with validation)
- **URL**: `/signup`
- **Method**: `POST`
- **Function**: `To create an alumni profile with validation`
- **Parameters**: `No parameters`
- **Request Body**:
    ```json
    {
        "FirstName": "John",
        "LastName": "Doe",
        "Fathername": "Michael Doe",
        "Password": "password123",
        "Status": "student", // "student" or "alumni"
        "Branch": "CSE",
        "BatchYear": 2022,
        "MobileNo": "1234567890", // unique
        "Email": "john.doe@example.com", // unique
        "EnrollmentNo": "ENR123456", // unique
        "Degree": "B.Tech"
    }
    ```
- **Success Response**: `200 OK`
    ```json
    {
        "message": "OTP sent successfully"
    }
    ```
- **Error Responses**:
    - `400 Bad Request`:
        ```json
        {
            "error": "Password cannot be empty"
        }
        ```
    - `409 Conflict` (if email already exists):
        ```json
        {
            "error": "Email already exists"
        }
        ```
    - `500 Internal Server Error`:
        ```json
        {
            "error": "Internal server error"
        }
        ```
### Verify OTP

- **URL**: `/signup`
- **Method**: `POST`
- **Function**: To check if the input OTP is correct and to register the alumni
- **Parameters**: None
- **Request Body**:
  ```json
  {
    "Email": "mohitgusain8671@gmail.com", // unique
    "OTP": "247077"
  }
  ```
- **Success Response**: `200 OK`
    ```json
    {
        "message":"User Verified Successfully",
    }
    ```
- **Error Responses**:
    - `400 Bad Request`:
        ```json
        {
            "message": "Invalid OTP"
        }
        ```
    - `409 Conflict` (if email already exists):
        ```json
        {
            "error": "Email already exists"
        }
        ```
    - `400 Bad Reques`:
        ```json
        {
            "message": "OTP expired"
        }
        ```
    - `500 Internal Server Error`:
        ```json
        {
            "error": "Internal server error"
        }


### Forgot Password

- **URL**: `/forgotPassword`
- **Method**: `POST`
- **Function**: To send a password reset email to the user
- **Parameters**: None
- **Headers**: None
- **Request Body**:
  ```json
  {
    "email": "user@example.com"
  }
- **Success Response**: `200 OK`
    ```json
    {
        "message": "Email received for reset Password",
        "token": "generated_token",
        "email": "user@example.com"
    }
    ```
- **Error Responses**:
    - `400 Bad Request`:
        ```json
        {
            "message": "Invalid request payload"
        }
        ```
    - `404 Not Found`:
        ```json
        {
            "message": "Email not found in AlumniProfile"
        }
        ```
    - `500 Internal Server Error`:
        ```json
        {
            "error": "Internal server error"
        }

### Reset Password


- **URL**: `/resetPassword`
- **Method**: `POST`
- **Function**: To reset the user's password using the token sent to their email
- **Parameters**: None
- **Headers**: None
- **Request Body**:
  ```json
  {
    "NewPassWord": "newpassword123",
    "token": "generated_token",
    "email": "user@example.com",
    "ConfirmNewPassword": "newpassword123"
  }
- **Success Response**:
    - `200 OK`
        - **Description**: Password has been reset successfully.
        - **Body**:
            ```json
            {
                "Password has been reset successfully"
            }
            ```
- **Error Responses**:
    - `400 Bad Request`

        - **Description**: Invalid input or token.

        - **Body**:

            - If the request body is invalid
                ```json
                {
                    "Invalid input"
                }
                ```
            - If the token is empty:
                ```json
                {
                    "Invalid token"
                }
                ```
            - If the token is invalid or expired:
                ```json
                {
                    "Invalid or expired token"
                }
                ```
            - If the token has expired:
                ```json
                {
                    "Token has expired"
                }
                ```
            - If the passwords do not match:
                ```json
                {
                    "Passwords do not match"
                }
                ```
    - `404 Not Found`

        - **Description**: Alumni not found.

        - **Body**:

            ```json
            {
                "Alumni not found"
            }
            ```
    - `500 Internal Server Error`

        - **Description**: Internal server error.

        - **Body**:

            - If there is an error hashing the password:
                ```json
                {
                    "Failed to hash password"
                }
                ```
            - If there is an error finding the alumni by email:
                ```json
                {
                    "Internal server error"
                }
                ```
            - If there is an error updating the password:
                ```json
                {
                    "Failed to update password"
                }
                ```
---
### ContactUS Form Submission
- **URL**: `/contactUS`
- **Method**: `POST`
- **Function**: `to send the details of ContactUS form to the designated email address/admin.`
- **Parameter**: `No Parameter`
- **Request Body**:
  ```json
  {
    "Name": "John Doe",
    "Email": "johndoe@example.com",
    "Contact": "1234567890",
    "Subject": "Inquiry about Alumni Event",
    "Message": "I would like to know more about the upcoming alumni event."
  }
- **Success Response**:
    - `200 OK`
        - **Description**: A JSON object indicating the success of the email sending operation.
        - **Body**:
            ```json
            {
                "message": "mail sent successfully"

            }
            ```
- **Error Responses**:
    - `500 Internal Server Error`
        - **Description**: Returned if there is an error sending the email.
        ```json
        {
            "error": "Failed to send email"
        }
        ```

    - `400 Bad Request`:
       - **Description**: Returned if the request body is not properly formatted JSON or missing required fields.

---
---
### Feedback Form Submission
- **URL**: `/feedback`
- **Method**: `POST`
- **Function**: `to send the Feedback to the designated email address/admin.`
- **Parameter**: `No Parameter`
- **Request Body**:
  ```json
  {
    "Name": "John Doe",
    "Email": "john.doe@example.com",
    "Feedback": "I really enjoyed using the platform. It is very user-friendly and efficient."
  }
  ```
- **Success Response**:
    - `200 OK`
        - **Description**: A JSON object indicating the success of the email sending operation.
        - **Body**:
            ```json
            {
                "message": "mail sent successfully"
            }
            ```
- **Error Responses**:
    - `500 Internal Server Error`
        - **Description**: Returned if there is an error sending the email.
        ```json
        {
            "error": "Failed to send email"
        }
        ```

    - `400 Bad Request`:
       - **Description**: Returned if the request body is not properly formatted JSON or missing required fields.

---
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
            "Status": "Student",
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
            "GeeksForGeeksProfile": "",
            "CodingNinjasProfile": "",
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
        "Status": "Student",
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
        "GeeksForGeeksProfile": "",
        "CodingNinjasProfile": "",
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
        "startDate": "2022-07-04T19:14:09.058+05:30",
        "endDate": "2024-02-04T19:14:09.058+05:30" // null for current company
    }
    ```
- **Success Response**: `201`
    ```json
    {
        "ProfID":5,
        "AlumniID":2,
        "CompanyName":"Tech Solutions Inc.",
        "Position":"Software Engineer",
        "StartDate": "2022-07-04T19:14:09.058+05:30",
        "EndDate": "2024-02-04T19:14:09.058+05:30",
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
            "StartDate": "2022-07-04T19:14:09.058+05:30",
            "EndDate": "2024-02-04T19:14:09.058+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"InstagramProfile":null,"TwitterProfile":null,"ProfilePicture":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"},
            "CreatedAt":"2024-07-02T16:10:43.344+05:30",
            "UpdatedAt":"2024-07-02T16:10:43.344+05:30"
        },
        {
            "ProfID":2,
            "AlumniID":2,
            "CompanyName":"Microsoft",
            "Position":"Software Engineer Intern",
            "StartDate": "2022-07-04T19:14:09.058+05:30",
            "EndDate": "2024-02-04T19:14:09.058+05:30",
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
            "StartDate": "2022-07-04T19:14:09.058+05:30",
            "EndDate": "2024-02-04T19:14:09.058+05:30",
            "Alumni":{"AlumniID":0,"FirstName":"","LastName":"","Branch":"","BatchYear":0,"MobileNo":"","Email":"","EnrollmentNo":"","Tenth":"","Xllth":"","Degree":"","GithubProfile":null,"LeetCodeProfile":null,"LinkedInProfile":null,"CodeforceProfile":null,"CodeChefProfile":null,"ProfilePicture":null,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}, 
            "CreatedAt":"2024-07-02T16:43:50.428+05:30",
            "UpdatedAt":"2024-07-03T16:17:13.509+05:30"
        },
        {
            "ProfID":5,
            "AlumniID":2,
            "CompanyName":"Tech Solutions Inc.",
            "Position":"Software Engineer",
            "startDate": "2022-07-04T19:14:09.058+05:30",
            "endDate": "2024-02-04T19:14:09.058+05:30",
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
        "startDate": "2022-07-04T19:14:09.058+05:30",
        "endDate": "2024-02-04T19:14:09.058+05:30"
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
        "startDate": "2022-07-04T19:14:09.058+05:30",
        "endDate": "2024-02-04T19:14:09.058+05:30",
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
            "CreatedAt":"2024-07-02T16:39:00.826+05:30",
            "UpdatedAt":"2024-07-04T19:29:20.524+05:30"
        },
        {
            "AchievementID":2,
            "AlumniID":2,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
            "CreatedAt":"2024-07-02T16:39:46.213+05:30",
            "UpdatedAt":"2024-07-02T16:39:46.213+05:30"
        },
        {
            "AchievementID":3,
            "AlumniID":1,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.","DateAchieved":"2023-05-20T20:34:05+05:30",
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
            "CreatedAt":"2024-07-02T16:39:00.826+05:30",
            "UpdatedAt":"2024-07-03T16:55:15.801+05:30",
        },
        {
            "AchievementID":3,
            "AlumniID":1,
            "Title":"First Prize in Coding Competition",
            "Description":"Won first prize in the national coding competition organized by ACM.",
            "DateAchieved":"2023-05-20T20:34:05+05:30",
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
## Gallery

### Add Image

* **URL**: `/gallery`
* **Method**: `POST`
* **Description**: Add a new image in gallery
* **Request Body**:
    ```json
    {
        "ImageLink": "image123.jpeg",   // image LINK
	    "ImageTitle": "Alumni Reunion 23"
    }
    ```
* **Success Response**:
    * **Code**: `201 Created`
    * **Content**:
      ```json
      {
          "ImageID": 1,
          "ImageLink": "image123.jpeg",   // image LINK
	      "ImageTitle": "Alumni Reunion 23",
          "ImageDescription": ""
      }
      ```


### Update Image

* **URL**: `/gallery/{id}`
* **Method**: `PUT`
* **Description**: Updates an information about existing image in gallery
* **Request Body**:
    ```json
    {
        "ImageDescription": "Reunion event held in 2023 for batch 2022"
    }
    ```
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the image to update.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      {
          "ImageID": 1,
          "ImageLink": "image123.jpeg",   // image LINK
	      "ImageTitle": "Alumni Reunion 23",
          "ImageDescription": "Reunion event held in 2023 for batch 2022"
      }
      ```


### Delete Image

* **URL**: `/gallery/{id}`
* **Method**: `DELETE`
* **Description**: Deletes an image from gallery by its ID.
* **URL Params**:
    * **Required**: `id=[integer]` - The ID of the image to delete.
* **Success Response**:
    * **Code**: `204 No Content`
    * **Content**: Empty response body.



### Get All Images

* **URL**: `/gallery`
* **Method**: `GET`
* **Description**: Retrieves all images from gallery.
* **URL Params**:
    No Parameters
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      [
          {
              "ImageID": 1,
              "ImageLink": "image123.jpeg",   // image LINK
	          "ImageTitle": "Alumni Reunion 23",
              "ImageDescription": "Reunion event held in 2023 for batch 2022"
          }
      ]
      ```
### Get All Albums

* **URL**: `/gallery/albums`
* **Method**: `GET`
* **Description**: Retrieves a list of album titles with a random image and the count of images for each title.
* **URL Params**:
    No Parameters
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      [
        {
            "image_title": "Alumni Reunion 21",
            "image_link": "https://img.com/images/alumni2021.jpg",
            "image_count": 2
        },
        {
            "image_title": "Alumni Reunion 22",
            "image_link": "https://img.com/images/alumni_reunion_2022.jpg",
            "image_count": 4
        },
        {
            "image_title": "Guest Lecture Series 2024",
            "image_link": "https://example.com/images/guest_lecture_2024.jpg",
            "image_count": 1
        }
      ]
      ```

### Get Image By Title (Read the Description)

* **URL**: `/gallery/{title}`
* **Method**: `GET`
* **Description**: Retrieves all images that match the specified title.<u><i>The title should be URL-encoded if it contains special characters (e.g., spaces)</i></u>.
* **URL Params**:
    `title` (path parameter): The title of the images to retrieve.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
      [
          {
            "ImageID": 990504068507566081,
            "ImageLink": "https://img.com/images/alumni2021.jpg",
            "ImageTitle": "Alumni Reunion 21",
            "ImageDescription": "Group photo of the Class of 2012 during their 10-year reunion."
          },
          {
            "ImageID": 990504088507875329,
            "ImageLink": "https://img.com/alumni2021.jpg",
            "ImageTitle": "Alumni Reunion 21",
            "ImageDescription": "Group photo of the Class of 2012 during their 10-year reunion."
          }
      ]
      ```
---
## Admin Routes

---
### Networking

#### Get All Data

* **URL**: `/admin/alumniattending`
* **Method**: `GET`
* **Description**: Retrieves all alumnis that are attending any events.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
        [
            {
                "AttendID": 5674477,
	            "EventID": 34566,
	            "AlumniID": 6745364767,
                "FirstName": "John",
                "LastName": "Doe",
                "Position": "Panelist",
                "Title": "Annual Alumni Meetup",
                "EventDateTime": "2024-08-15T15:30:00+05:30",
                "Location": "Online"
            }
        ]
      ```

#### Create New Networking

* **URL**: `/admin/alumniattending`
* **Method**: `POST`
* **Description:** Creates a new record of an alumni attending an event.
* **Request Body:**
    ```json
    {
        "AlumniID": 45789356,
        "Position": "Speaker",
        "Title": "TECH TALKS",
        "Description": "Speaker will give info about webD",
        "EventType": "Webinar",
        "ModeOfEvent": "Online",     
        "Location": "Google Meet",
        "EventDateTime": "2024-08-15T15:30:00+05:30"  
    }
    ```

* **Responses:**

    201 Created
    ```json
        {
            "message": "Event and AlumniAttending created successfully"
        }
    ```

    400 Bad Request: Invalid request payload.

    500 Internal Server Error: If there’s a server error. and failed to create any table/data

#### Update Networking
* **URL**: `/admin/alumniattending/{event_id}`
* **Method**: `PUT`
* **Description**: Update the Details about Event.
* **Path Parameters:**
    - `event_id` (integer): The ID of the event.
* **Request Body**:
    ```json
    {
        "AlumniID": 45789356,
        "Position": "Panelist",
        "Title": "TECH TALKS",
        "Description": "Speaker will give info about webD",
        "EventType": "Webinar",
        "ModeOfEvent": "Offline",     
        "Location": "On Campus",
        "EventDateTime": "2024-08-15T15:30:00+05:30" 
        // Only Put field that you want to update 
    }
    ```
* **Responses:**
    - 200 OK
    ```json
    {
        "message": "Event and AlumniAttending updated successfully"
    }
    ```
    - 400 Bad Request: Invalid request payload.
    - 404 Not Found: Alumni attending or Event record not found.
    - 500 Internal Server Error: If there’s a server error or failed to update.

#### Get Networking Details by Alumni ID

* **URL**: `/admin/alumniattending/{alumni_id}`
* **Method**: `GET`
* **Description:** Retrieves all events an alumni is attending based on their `alumni_id`.
* **Path Parameters:**
    - `alumni_id` (integer): The ID of the alumni.

* **Responses:**

    - 200 OK
    ```json
        [
            {
                "AttendID": 5674477,
	            "EventID": 34566,
	            "AlumniID": 6745364767,
                "FirstName": "John",
                "LastName": "Doe",
                "Position": "Panelist",
                "Title": "Annual Alumni Meetup",
                "EventDateTime": "2024-08-15T15:30:00+05:30",
                "Location": "Online"
            },
            ...
        ]
    ```

    - 404 Not Found: No attending records found for the given alumni ID.

    - 500 Internal Server Error: If there’s a server error.

#### Delete Networking 
- [**Delete API of EVENT**](#delete-an-event) 

---
### Achievements

#### Add Achievements
- [**Refer Create API of Achievements**](#add-achievements)

#### Update Achievements
- [**Refer Update API of Achievements**](#update-achievement-information)

#### Delete the Achievements
- [**Delete API of Achievements**](#delete-achievement)

#### Get All Achievements And Alumni Details
* **URL**: `/admin/achievements`
* **Method**: `GET`
- **Description**:
Fetches a list of achievements for alumni along with their associated profile information.

- **Response**:

    - Status Codes:
        - **200 OK**: Successfully retrieved the list of achievements.
        - **500 Internal Server Error**: An error occurred while retrieving the data.

    - **Response Body**:

    ```json
    [
        {
            "AchievementID": int64,
            "AlumniID": int64,
            "FirstName": "string",
            "LastName": "string",
            "Branch": "string",
            "BatchYear": int64,
            "Title": "string",
            "Description": "string",
            "DateAchieved": "YYYY-MM-DDTHH:MM:SSZ"
        },
        ...
    ]
    ```
---
### Student Details
#### Add Student Details
- [**Refer Create API of Student Details**](#create-an-alumni-profile): 
    ```
    NOTE:-  in Request body provide the ("Status" = "student")
    ```
#### Update Student Details
- [**Refer Update API of Student Details**](#update-an-alumni-profile)

#### Delete Student Details
- [**Delete API of Student Details**](#delete-an-alumni-profile)

#### Get All Student Details
- **URL**: `/alumni/achievements?status=student`
- **Method**: `GET`
- **Description**:
Fetches a list of students.
- **Response**: it is same as of [Fetch Alumni API](#get-all-alumni-profiles)
---
### Alumni Details

#### Add Alumni Details
- **NOTE**: We need two forms one is to add alumni and other is to add professional info of alumni after alumni creation
- **Add Alumni Details**
    - [**Refer Create API of Alumni Details**](#create-an-alumni-profile):
        ```
        NOTE:-  in Request body provide the ("Status" = "alumni")
        ```
- **Add Professional Info**
    - if you are currently working then in Enddate insert today's date
    - [**Refer Create API of Professional Info**](#add-professional-information)
    - use AlumniID for creation

#### Delete Alumni
- [**Delete API of Alumni Details**](#delete-an-alumni-profile)

#### Update 
- [**Update Api of Professional info**](#update-professional-information)
- [**Update Api of Alumni**](#update-an-alumni-profile)
#### Get All Alumni
- **URL**: `/admin/alumniprofiles`
- **Method**: `GET`
- **Description**: Fetches a list of alumni and its current company or the last company in which he/she works.
- **Response**:
    - Status Codes:
        - **200 OK**: Successfully retrieved the list of achievements.
        - **500 Internal Server Error**: An error occurred while retrieving the data.

    - **Response Body**:

    ```json
    [
        {
            "AlumniID": 991079807932465153,
            "FullName": "Mohit Gusain",
            "BatchYear": 2022,
            "Branch": "CSE",
            "Email": "mohitgn8671@gmail.com",
            "MobileNo": "9667897066",
            "CurrentCompany": {
                "ProfID": 991083253114732545,
                "AlumniID": 991079807932465153,
                "CompanyName": "Code Solutions",
                "Position": "Senior Developer",
                "StartDate": "2024-07-01T05:30:00+05:30",
                "EndDate": "2024-12-31T05:30:00+05:30",
                "CreatedAt": "2024-08-01T20:40:04.615969+05:30",
                "UpdatedAt": "2024-08-01T20:40:04.615969+05:30"
            }
        },
        {
            "AlumniID": 991080155239055361,
            "FullName": "Mohit Gusain",
            "BatchYear": 2022,
            "Branch": "CSE",
            "Email": "mohitgnk8671@gmail.com",
            "MobileNo": "96699796896",
            "CurrentCompany": {
                "ProfID": 991083122149588993,
                "AlumniID": 991080155239055361,
                "CompanyName": "Creative Solutions LLC",
                "Position": "Marketing Manager",
                "StartDate": "2022-05-01T05:30:00+05:30",
                "EndDate": "2024-08-01T05:30:00+05:30",
                "CreatedAt": "2024-08-01T20:39:24.65292+05:30",
                "UpdatedAt": "2024-08-01T20:39:24.65292+05:30"
            }
        },
        ...
    ]
    ```

#### Get an Alumni Details
- **URL**: `/admin/alumniprofiles/{id}`
- **Method**: `GET`
- **Request Body**: None
* **Path Parameters:**
    - `id` (integer): The ID of the alumni.
- **Response Body**: Alumni details in JSON format.
    ```json
    {
        "AlumniID": 991080155239055361,
        "FirstName": "Mohit",
        "LastName": "Gusain",
        "Fathername": "",
        "Password": "",
        "status": "alumni",
        "Branch": "CSE",
        "BatchYear": 2022,
        "MobileNo": "96699796896",
        "Email": "mohitgnk8671@gmail.com",
        "EnrollmentNo": "0895207022",
        "IsVerified": false,
        "IsApproved": false,
        "Code": "",
        "ExpiresAt": "0001-01-01T05:30:00+05:30",
        "Tenth": "89.5",
        "Xllth": "90",
        "Degree": "B.tech(CSE)",
        "GithubProfile": "github.com/mohitgusain8671",
        "LeetCodeProfile": "leetcode.com/mohitgusain8671",
        "LinkedInProfile": "linkedin.com",
        "CodeforceProfile": "link.com",
        "CodeChefProfile": "",
        "InstagramProfile": "",
        "TwitterProfile": "",
        "GeeksForGeeksProfile": null,
        "CodingNinjasProfile": null,
        "ProfilePicture": "",
        "ProfessionalInformation": [
            {
                "ProfID": 991083008906395649,
                "AlumniID": 991080155239055361,
                "CompanyName": "Marketing Masters",
                "Position": "Product Marketing Specialist",
                "StartDate": "2024-08-02T05:30:00+05:30",
                "EndDate": "0001-01-01T05:30:00+05:30",
                "CreatedAt": "2024-08-01T20:38:49.946307+05:30",
                "UpdatedAt": "2024-08-01T20:38:49.946307+05:30"
            },
            {
                "ProfID": 991083122149588993,
                "AlumniID": 991080155239055361,
                "CompanyName": "Creative Solutions LLC",
                "Position": "Marketing Manager",
                "StartDate": "2022-05-01T05:30:00+05:30",
                "EndDate": "2024-08-01T05:30:00+05:30",
                "CreatedAt": "2024-08-01T20:39:24.65292+05:30",
                "UpdatedAt": "2024-08-01T20:39:24.65292+05:30"
            }
        ],
        "Achievements": null,
        "InterestsHobbies": null,
        "AlumniAttending": null,
        "InterviewExperience": null,
        "CreatedAt": "2024-08-01T20:24:19.164954+05:30",
        "UpdatedAt": "2024-08-01T20:24:19.164954+05:30"
    }
    ```
---

### NEWS

- **NOTE** : The news section is generated dynamically from existing tables without creating a dedicated News table. hence we cannot add new/update/delete news directly from here 

#### Get All News
- **URL**: `/admin/news`
- **Method**: `GET`
- **Request Body**: None
- **Response**: JSON object containing all news items
    ```json
    [
        {
            "title": "Achievement",
            "description": "Mohit Gusain achieved First Prize in Flipkart Grid on May 20, 2023.",
            "date": "2024-08-01T20:34:06.988088+05:30"
        },
        {
            "title": "Achievement",
            "description": "Mohit Gusain achieved First Prize in Amazon Hackon on May 20, 2023.",
            "date": "2024-08-01T20:33:47.105+05:30"
        },
        {
            "title": "Achievement",
            "description": "Mohit Gusain achieved First Prize in Hackathon on May 20, 2023.",
            "date": "2024-08-01T20:33:16.0039+05:30"
        },
        {
            "title": "Achievement",
            "description": "Mohit Gusain achieved First Prize in Coding Competition on May 20, 2023.",
            "date": "2024-08-01T20:32:50.042142+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A TECH TALKS event is going to be held on August 15, 2024 at Google Meet.It is an Online event",
            "date": "2024-07-31T21:57:16.870142+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A TECH TALKS event is going to be held on August 15, 2024 at Google Meet.It is an Online event",
            "date": "2024-07-31T21:57:14.425599+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A TECH TALKS event is going to be held on August 15, 2024 at Google Meet.It is an Online event",
            "date": "2024-07-31T21:57:07.913411+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A TECH TALKS event is going to be held on August 15, 2024 at Google Meet.It is an Online event",
            "date": "2024-07-31T21:55:37.52872+05:30"
        },
        {
            "title": "Professional Update",
            "description": "test test got placed in Google at the position of Software Engineer in 0001-01-01 00:00:00 +0000 UTC",
            "date": "2024-07-29T21:32:40.166579+05:30"
        },
        {
            "title": "Professional Update",
            "description": "test test got placed in Tech Solutions Inc. at the position of Software Engineer in 0001-01-01 00:00:00 +0000 UTC",
            "date": "2024-07-29T21:32:11.36637+05:30"
        },
        {
            "title": "Achievement",
            "description": "test test achieved  on January 1, 0001.",
            "date": "2024-07-29T11:32:49.049891+05:30"
        },
        {
            "title": "Achievement",
            "description": "test test achieved Congratulations for qualifying GATE-2023  on May 20, 2023.",
            "date": "2024-07-29T11:02:17.194358+05:30"
        },
        {
            "title": "Achievement",
            "description": "John Doe achieved Third Prize in ICPC on May 20, 2023.",
            "date": "2024-07-29T10:58:38.622578+05:30"
        },
        {
            "title": "Achievement",
            "description": "John Doe achieved Won the Flipkart Grid  on May 20, 2023.",
            "date": "2024-07-29T10:51:45.46616+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A Know your alumni event is going to be held on August 15, 2024 at oncampus.It is an Offline event",
            "date": "2024-07-20T13:15:51.391998+05:30"
        },
        {
            "title": "Upcoming Event",
            "description": "A Annual Alumni Meetup event is going to be held on August 15, 2024 at Online.It is an Virtual event",
            "date": "2024-07-20T13:14:45.366126+05:30"
        }
    ]
    ```
<!-- ### Get All Alumni Attending Events

* **URL**: `/admin/alumniattending`
* **Method**: `GET`
* **Description**: Retrieves all alumnis that are attending any events.
* **Success Response**:
    * **Code**: `200 OK`
    * **Content**:
      ```json
        [
            {
                "FirstName": "John",
                "LastName": "Doe",
                "Position": "Panelist",
                "Title": "Annual Alumni Meetup",
                "EventDateTime": "2024-08-15T15:30:00+05:30",
                "Location": "Online"
            }
        ]
      ```

### Create Alumni Attending

* **Endpoint:**  POST `/alumniattending`
* **Description:** Creates a new record of an alumni attending an event.
* **Request Body:**
    ```json
        {
            "alumni_id": 123,
            "event_id": 456,
            "position": "Speaker"
        }
    ```
---

* **Responses:**

    201 Created
    ```json
        {
            "message": "Alumni attending record created successfully"
        }
    ```

    400 Bad Request: Invalid request payload.

    500 Internal Server Error: If there’s a server error.


### Get Alumni Attending by Alumni ID and Event ID

* **Endpoint:** `GET /alumniattending/{alumni_id}/{event_id}`
* **Description:** Retrieves a specific record of an alumni attending an event based on `alumni_id` and `event_id`.
* **Path Parameters:**

    - `alumni_id` (integer): The ID of the alumni.
    - `event_id` (integer): The ID of the event.

* **Responses:**
    - 200 OK
    ```json
        {
            "attend_id": 789,
            "alumni_id": 123,
            "event_id": 456,
            "position": "Speaker",
            "created_at": "2024-07-30T12:00:00Z",
            "updated_at": "2024-07-30T12:00:00Z"
        }
    ```
    - 404 Not Found: Alumni attending record not found.
    - 500 Internal Server Error: If there’s a server error.

### Update Alumni Attending
* **Endpoint:** `PUT /alumniattending/{alumni_id}/{event_id}`
* **Description:** Updates an existing record of an alumni attending an event.
* **Path Parameters:**

    - `alumni_id` (integer): The ID of the alumni.
    - `event_id` (integer): The ID of the event.
* **Request Body:**
    ```json
        {
            "position": "Moderator"
        }
    ```
* **Responses:**
    - 200 OK
    ```json
    {
        "message": "Alumni attending record updated successfully"
    }
    ```
    - 400 Bad Request: Invalid request payload.
    - 404 Not Found: Alumni attending record not found.
    - 500 Internal Server Error: If there’s a server error.

### Delete Alumni Attending
* **Endpoint:** `DELETE /alumniattending/{alumni_id}/{event_id}`
* **Description:** Deletes a specific record of an alumni attending an event.
* **Path Parameters:**
    - `alumni_id` (integer): The ID of the alumni.
    - `event_id` (integer): The ID of the event.
* **Responses:**
    - 200 OK
    ```json
    {
        "message": "Alumni attending record deleted successfully"
    }
    ```
    - 404 Not Found: Alumni attending record not found.
    - 500 Internal Server Error: If there’s a server error. -->

## Database Migrations

The API will automatically create the necessary database tables if they do not exist when an endpoint is accessed.

## Error Handling

The API will return appropriate HTTP status codes and error messages for invalid requests or server errors.

- `400 Bad Request`: The request was invalid.
- `404 Not Found`: The requested record was not found.
- `500 Internal Server Error`: An error occurred on the server.



