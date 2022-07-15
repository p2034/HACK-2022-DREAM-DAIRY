--------------------------------------------
***PUBLICATION:***
--------------------------------------------

**/publication** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "publication": {
    "goalid": 4
  }
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/publication** (GET)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",

  "creator": 67930, // this is just user id
  "publicationid": 1
}
```

Answer:
```json
{
  "error": [
    "error name" 
  ],
  "publication": {
    "title": "4 dream",
    "description": "My goal is ...",
    "photos": [
      "BASE64 encrypted photo"
    ],
    "date": "18.05.2099",
    "tasks": [
      {
        "taskid": 0,
        "title": "task 1",
        "state": true // true if it's finished
      } // ...
    ]
  }
}
```

**/publications** (GET)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",

  "creator": 67930, // this is just user id
  "firstid: 1,
  "lastid": 3,
}
```

Answer:
```json
{
  "error": [
    "error name" 
  ],
  "publications": [
    {
      "title": "4 dream",
      "description": "My goal is ...",
      "photos": [
        "BASE64 encrypted photo"
      ],
      "date": "18.05.2099",
      "tasks": [
        {
          "taskid": 0,
          "title": "task 1",
          "state": true // true if it's finished
        } // ...
      ]
    } // ...
  ]
}
```

**/publication** (DELETE)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "publicationid": 0
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/reaction** (PUT)
For reactions on publications

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",

  "creator": 67930, // this is just user id
  "publicationid": 1,
  "reactiontype": 3 // there must be emoji id
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

--------------------------------------------
***GOAL***
--------------------------------------------

**/goal ** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "goal": {
    "title": "Some Dream",
    "description": "This dream is ...",
    "photos": [
      "BASE64 encrypted photo",
      "BASE64 encrypted photo"
    ],
    "date": "18.05.2099"
  }
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/goal ** (GET)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "first": 3, // id of the goal in db from 0
  "last": 4
}
```

Answer:
```json
{
  "error": [
    "error name"
  ],
  "goals": [
    {
      "goalid": 7,
      "title": "4 dream",
      "description": "My goal is ...",
      "photos": [
        "BASE64 encrypted photo"
      ],
      "date": "18.05.2099",
      "tasks": [
        {
          "taskid": 0,
          "title": "task 1",
          "state": true // true if it's finished
        } // ...
      ]
    } // ...
  ]
}
```

**/goal/task ** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "goalid": 7,
  "task": {
    "title": "task 1"
  }
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/goal/task ** (DELETE)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "goalid": 7,
  "taskid": 0
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/goal/task ** (PUT)
Used for changing state of the task. (finished or not)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "goalid": 7,
  "taskid": 0
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

--------------------------------------------
***DREAM:***
--------------------------------------------

**/dream** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "dream": {
    "title": "Some Dream",
    "description": "This dream is ...",
    "photos": [
      "BASE64 encrypted photo",
      "BASE64 encrypted photo"
    ]
  }
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/dream** (GET)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "first": 3, // id of the dream in db from 0
  "last": 4
}
```

Answer:
```json
{
  "error": [
    "error name"
  ],
  "dreams": [
    {
      "dreamid": 2,
      "title": "3 dream",
      "description": "This dream is ...",
      "photos": [
        "BASE64 encrypted photo",
        "BASE64 encrypted photo"
      ]
    },
    {
      "dreamid": 3,
      "title": "4 dream",
      "description": "This dream is ...",
      "photos": [
        "BASE64 encrypted photo"
      ]
    }
  ]
}
```

**/dream//makegoal** (POST)
Move dream in goals

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "dreamid": 3,
  "date": "18.05.2099"
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

--------------------------------------------
***POST FEED***
--------------------------------------------

***/tips*** (GET)

Answer:
```json
{
  "tips": [
    "url on youtube videos"
  ]
}
```

--------------------------------------------
***USER:***
--------------------------------------------

**/user/photo** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78...",
  "photo": "BASE64 encypted photo"
}
```

Answer:
```json
{
  "error": [
    "error name"
  ]
}
```

**/user?username=name** (GET)
to get user profile

Answer:
```json
{
  "error": [
    "error name"
  ],
  "photo": "url",
  "name": "Steve",
  "quiz": [
    {
      "question": "",
      "answer": ""
    },
    {
      "question": "",
      "answer": ""
    }
  ],
  "publicationCount": 4
}
```